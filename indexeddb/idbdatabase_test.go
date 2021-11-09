package indexeddb

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

//cant direct call but discover do it for us
func TestIDBDatabaseNewFromJSObject(t *testing.T) {

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "iddb"); testingutils.AssertErr(t, err) {

		if factory, err := IDBFactoryNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if openrequest, err := factory.Open("db", "1"); testingutils.AssertErr(t, err) {

				openrequest.OnSuccess(func(e event.Event) {

					if result, err := openrequest.Result(); err == nil {

						if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

							testingutils.AssertExpect(t, "[object IDBDatabase]", db.IDBDatabase_().ToString_())
							io <- true

						}
					}

				})

			}

			factory.DeleteDatabase("test")
		}

	}

	select {
	case <-io:
	case <-time.After(time.Duration(500) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

var methodsIDBDatabaseAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": "databasemethods"},
	{"method": "Version", "resultattempt": int64(1)},
	{"method": "CreateObjectStore", "args": []interface{}{"utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}}, "type": "constructnamechecking", "resultattempt": "IDBObjectStore"},
	{"method": "ObjectStoreNames", "type": "constructnamechecking", "resultattempt": "DOMStringList"},
	{"method": "DeleteObjectStore", "args": []interface{}{"utilisateur"}, "type": "error", "resultattempt": nil},
	{"method": "Close", "type": "error", "resultattempt": nil},
}

func TestIDBDatabaseMethods(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	dbdatabase=iddb.open("databasemethods")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "dbdatabase"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, ok := result.(IDBDatabaseFrom); testingutils.AssertExpect(t, true, ok) {

						for _, result := range methodsIDBDatabaseAttempt {
							testingutils.InvokeCheck(t, db.IDBDatabase_(), result)
						}

						io <- true

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
	iddb.deleteDatabase("dbdatabase")
	`)
}
