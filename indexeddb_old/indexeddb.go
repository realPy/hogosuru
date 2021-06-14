package indexeddb_old

import (
	"errors"
	"fmt"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/indexeddb/idbdatabase"
)

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

		return baseobject.StringWithErr(jserr)

	}

	return "", err

}

var singleton sync.Once

var indexeddbinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var indexeddbinstance JSInterface
		var window js.Value
		var err error

		if window, err = js.Global().GetWithErr("window"); err == nil {
			if indexeddbinstance.objectInterface, err = window.GetWithErr("indexedDB"); err == nil {
				indexeddbinterface = &indexeddbinstance
			}

		}
	})

	return indexeddbinterface
}

type IDBOpenDBRequest struct {
	baseobject.BaseObject
}

func Open(name string, version int,
	automigratehandler func(idbdatabase.IDBDatabase) error,
	onsuccesshandler func(idbdatabase.IDBDatabase) error,
	onerrorhandler func(error)) error {
	var waitableOpen js.Value
	var err error

	if dbi := GetJSInterface(); dbi != nil {
		if version == 0 {
			waitableOpen, err = dbi.objectInterface.CallWithErr("open", js.ValueOf(name))
		} else {
			waitableOpen, err = dbi.objectInterface.CallWithErr("open", js.ValueOf(name), js.ValueOf(version))
		}

		if err == nil {
			migrateDBFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

				if idbObject, err := waitableOpen.GetWithErr("result"); err == nil {

					idb, _ := idbdatabase.NewFromJSObject(idbObject)

					return automigratehandler(idb)
				} else {
					fmt.Printf("result not found..\n")
				}

				return nil
			})

			waitableOpen.Set("onupgradeneeded", migrateDBFunc)

			onsuccess := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

				if idbObject, err := waitableOpen.GetWithErr("result"); err == nil {

					idb, _ := idbdatabase.NewFromJSObject(idbObject)

					onsuccesshandler(idb)
				} else {
					fmt.Printf("result not found..\n")
				}

				return nil
			})

			onerror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				err = fmt.Errorf("Unable to open indexeddb")
				var errorString string

				if len(args) > 0 {
					event := args[0]
					if errorString, err = stringFromTargetError(event); err == nil {
						err = errors.New(errorString)
					}
				}

				onerrorhandler(err)
				return nil
			})

			waitableOpen.Set("onsuccess", onsuccess)
			waitableOpen.Set("onerror", onerror)

		}

	}

	return err
}
