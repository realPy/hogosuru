package indexdb

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type IndexDB struct {
	dbObject js.Value
}

type SuccessFailure struct {
	Success bool
	Payload []js.Value
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

func OpenIndexDB(name string, automigrate func(js.Value) error) (IndexDB, error) {
	var indexdb IndexDB

	if window, err := js.Global().GetWithErr("window"); err == nil {

		if indexedDB, err := window.GetWithErr("indexedDB"); err == nil {

			if waitableOpen, err := indexedDB.CallWithErr("open", js.ValueOf(name)); err == nil {
				test := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

					db := args[0].Get("target").Get("result")
					return automigrate(db)

				})
				waitableOpen.Set("onupgradeneeded", test)
				ch := OnSuccessFailure(waitableOpen)
				results := <-ch
				if results.Success {
					ev := results.Payload[0]
					if target, err := ev.GetWithErr("target"); err == nil {
						if db, err := target.GetWithErr("result"); err == nil {
							indexdb.dbObject = db
						} else {
							return IndexDB{}, fmt.Errorf("result not found")
						}
					} else {
						return IndexDB{}, fmt.Errorf("target not found")
					}

					return indexdb, nil
				} else {
					return IndexDB{}, fmt.Errorf("Open IndexDB failed")
				}

			} else {
				return IndexDB{}, err
			}
		} else {
			return IndexDB{}, err
		}

	} else {
		return IndexDB{}, err
	}

}

func (i IndexDB) Store(table string, value map[string]interface{}) error {

	if transaction, err := i.dbObject.CallWithErr("transaction", js.ValueOf(table), js.ValueOf("readwrite")); err == nil {

		if objectstore, err := transaction.CallWithErr("objectStore", js.ValueOf(table)); err == nil {
			objectstore.CallWithErr("add", js.ValueOf(value))
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
