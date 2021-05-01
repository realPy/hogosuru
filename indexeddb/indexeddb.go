package indexeddb

import (
	"errors"
	"fmt"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

type IndexedDB struct {
	dbobject     js.Value
	objinterface object.ObjectInterface
}

type SuccessFailure struct {
	Success bool
	Payload []js.Value
}

func getEventTargetResult(ev js.Value) (js.Value, error) {
	if target, err := ev.GetWithErr("target"); err == nil {
		if result, err := target.GetWithErr("result"); err == nil {
			return result, nil
		} else {
			return js.Value{}, fmt.Errorf("result not found")
		}
	} else {
		return js.Value{}, fmt.Errorf("target not found")
	}
}

func getEventTargetError(ev js.Value) (js.Value, error) {
	if target, err := ev.GetWithErr("target"); err == nil {
		if errorResult, err := target.GetWithErr("error"); err == nil {
			return errorResult, nil
		} else {
			return js.Value{}, fmt.Errorf("error not found")
		}
	} else {
		return js.Value{}, fmt.Errorf("target not found")
	}
}

func stringFromTargetError(ev js.Value) (string, error) {
	var err error
	var jserr js.Value

	if jserr, err = getEventTargetError(ev); err == nil {

		return object.StringWithErr(jserr)

	}

	return "", err

}

func OnSuccessFailure(awaitable js.Value) chan SuccessFailure {
	ch := make(chan SuccessFailure)
	cbok := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- SuccessFailure{Success: true, Payload: args}
		return nil
	})
	cberror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		ch <- SuccessFailure{Success: false, Payload: args}
		return nil
	})
	awaitable.Set("onsuccess", cbok)
	awaitable.Set("onerror", cberror)
	return ch
}

func OpenIndexedDB(name string, version int, automigrate func(js.Value) error) (IndexedDB, error) {
	var indexdb IndexedDB
	var err error
	var window, indexedDBObject, waitableOpen, db js.Value

	if indexdb.objinterface, err = object.NewObjectInterface(); err != nil {
		return indexdb, err
	}

	if window, err = js.Global().GetWithErr("window"); err == nil {

		if indexedDBObject, err = window.GetWithErr("indexedDB"); err == nil {
			if version == 0 {
				waitableOpen, err = indexedDBObject.CallWithErr("open", js.ValueOf(name))
			} else {
				waitableOpen, err = indexedDBObject.CallWithErr("open", js.ValueOf(name), js.ValueOf(version))
			}

			if err == nil {
				migrateDBFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

					db := args[0].Get("target").Get("result")
					return automigrate(db)

				})
				waitableOpen.Set("onupgradeneeded", migrateDBFunc)
				ch := OnSuccessFailure(waitableOpen)
				results := <-ch
				if results.Success {
					ev := results.Payload[0]
					if db, err = getEventTargetResult(ev); err == nil {
						indexdb.dbobject = db
						return indexdb, nil
					}
				} else {
					// recuperer error https://developer.mozilla.org/fr/docs/Web/API/IDBRequest/error
					err = fmt.Errorf("Unable to open indexeddb")
					ev := results.Payload[0]
					var errorString string
					if errorString, err = stringFromTargetError(ev); err == nil {
						err = errors.New(errorString)
					}

				}

			}
		}

	}
	return indexdb, err
}

func (i IndexedDB) GetObjectStore(table string, permission string) (Store, error) {
	if transaction, err := i.dbobject.CallWithErr("transaction", js.ValueOf(table), js.ValueOf(permission)); err == nil {

		if objectstore, err := transaction.CallWithErr("objectStore", js.ValueOf(table)); err == nil {
			return Store{objstore: objectstore, objinterface: i.objinterface}, nil
		} else {
			return Store{}, err
		}
	} else {
		return Store{}, err
	}
}
