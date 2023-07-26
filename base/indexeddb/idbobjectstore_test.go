package indexeddb

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/testingutils"
)

var methodsIDBObjectStoreAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "CreateIndex", "args": []interface{}{"email", "emailkey", map[string]interface{}{"unique": true}}, "type": "constructnamechecking", "resultattempt": "IDBIndex"},
	{"method": "Add", "args": []interface{}{map[string]interface{}{"email": "yesmymaail", "data": "name"}}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Get", "args": []interface{}{"yesmymaail"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAll", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetAllKeys", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "GetKey", "args": []interface{}{"email"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Index", "args": []interface{}{"email"}, "type": "constructnamechecking", "resultattempt": "IDBIndex"},
	{"method": "Put", "args": []interface{}{map[string]interface{}{"email": "yesm2", "data": "name"}}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "OpenKeyCursor", "type": "constructnamechecking", "resultattempt": "IDBRequest"},

	{"method": "Delete", "args": []interface{}{"yesmymaail"}, "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Count", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "Clear", "type": "constructnamechecking", "resultattempt": "IDBRequest"},
	{"method": "DeleteIndex", "args": []interface{}{"email"}, "type": "error", "resultattempt": nil},
}

func TestIDBObjectStoreMethods(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	objectstore=iddb.open("objectstore")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "objectstore"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

						if store, err := db.IDBDatabase_().CreateObjectStore("utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							testingutils.AssertExpect(t, "[object IDBObjectStore]", store.ToString_())

							for _, result := range methodsIDBObjectStoreAttempt {
								testingutils.InvokeCheck(t, store, result)
							}

							io <- true

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
	iddb.deleteDatabase("objectstore")
	`)
}
