package indexeddb

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

var methodsIDBIndexeAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "KeyPath", "resultattempt": "hello"},
	{"method": "Name", "resultattempt": "hello"},
	{"method": "MultiEntry", "resultattempt": false},
	{"method": "ObjectStore", "type": "constructnamechecking", "resultattempt": "IDBObjectStore"},
	{"method": "Unique", "resultattempt": false},
	{"method": "Count", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Get", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetKey", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAll", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAllKeys", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenKeyCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
}

func TestIDBIndexMethods(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	objectstore=iddb.open("idbindex")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "objectstore"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("user", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							if index, err := store.CreateIndex("hello", "hello"); testingutils.AssertErr(t, err) {

								for _, result := range methodsIDBIndexeAttempt {
									testingutils.InvokeCheck(t, index, result)
								}

								io <- true

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
	iddb.deleteDatabase("idbindex")
	`)
}
