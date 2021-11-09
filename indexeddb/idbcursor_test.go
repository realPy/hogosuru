package indexeddb

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

var methodsIDBCursorAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Direction", "resultattempt": "next"},
	{"method": "Key", "resultattempt": 1},
	{"method": "PrimaryKey", "resultattempt": 1},
	{"method": "Source", "type": "constructnamechecking", "resultattempt": "IDBObjectStore"},
	{"method": "Request", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Update", "args": []interface{}{map[string]interface{}{"id": 1, "email": "test", "data": "name"}}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Delete", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
}

func TestIDBCursorMethods(t *testing.T) {
	baseobject.Eval(`iddbcursor=window.indexedDB
	objectstorecursor=iddbcursor.open("idbcursor")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "objectstorecursor"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); testingutils.AssertErr(t, err) {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							store.Add(map[string]interface{}{"email": "yesmymaail", "data": "name"})

							if request, err := store.OpenCursor(); testingutils.AssertErr(t, err) {

								request.OnSuccess(func(e event.Event) {

									if result, err := request.Result(); testingutils.AssertErr(t, err) {

										if cursorfrom, ok := result.(IDBCursorFrom); testingutils.AssertExpect(t, true, ok) {

											for _, result := range methodsIDBCursorAttempt {
												testingutils.InvokeCheck(t, cursorfrom.IDBCursor_(), result)
											}

											io <- true

										}
									}

								})

							}

						}

					}
				}

			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	baseobject.Eval(`
	iddb.deleteDatabase("idbcursor")
	`)
}

func TestIDBContinue(t *testing.T) {
	baseobject.Eval(`iddb2cursor=window.indexedDB
	objectstorecursor2=iddb2cursor.open("idbcursorc")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "objectstorecursor2"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); testingutils.AssertErr(t, err) {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("utilisateur2", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							store.Add(map[string]interface{}{"email": "yesmymaail", "data": "name"})

						}

					}
				}

			})

			openrequest.OnSuccess(func(e event.Event) {
				if result, err := openrequest.Result(); testingutils.AssertErr(t, err) {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {
						if transaction, err := db.IDBDatabase_().Transaction("utilisateur2", "readwrite"); err == nil {
							if store, err := transaction.ObjectStore("utilisateur2"); err == nil {
								//warn opencursor will excute On Success for each IDBCursor and finalluyu execute last time with null
								//t.Errorf not work on go routine so an error occur it crash. Another way must be find

								if request, err := store.OpenCursor(); testingutils.AssertErr(t, err) {

									request.OnSuccess(func(e event.Event) {

										if result, err := request.Result(); err != baseobject.ErrUndefinedValue {

											if cursorfrom, ok := result.(IDBCursorFrom); testingutils.AssertExpect(t, true, ok) {

												cursorfrom.IDBCursor_().Continue()

												io <- true

											}
										}

									})

								}

							}
						}

					}
				}

			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	baseobject.Eval(`
	iddb.deleteDatabase("idbcursorc")
	`)
}
