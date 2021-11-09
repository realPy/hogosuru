package indexeddb

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

func TestIDBRequestNewFromJSObject(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	dbrequest=iddb.open("openrequest2")
	`)

	if obj, err := baseobject.Get(js.Global(), "dbrequest"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "IDBOpenDBRequest", openrequest.ConstructName_())
		}

	}
	baseobject.Eval(`
	dbrequest=iddb.deleteDatabase("openrequest2")
	`)

}

var methodsIDBRequestAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Error", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "ReadyState", "resultattempt": "done"},
	{"method": "Source", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "Transaction", "type": "constructnamechecking", "resultattempt": "IDBTransaction"},
}

func TestIDBRequestMethods(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	dbrequest=iddb.open("openrequestmethods")
	`)

	var io chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "dbrequest"); testingutils.AssertErr(t, err) {

		if openrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				for _, result := range methodsIDBRequestAttempt {
					testingutils.InvokeCheck(t, openrequest, result)
				}

				io <- true
			})

			select {
			case <-io:
			case <-time.After(time.Duration(5000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}

		}

	}

	baseobject.Eval(`
	dbrequest=iddb.deleteDatabase("openrequestmethods")
	`)
}
