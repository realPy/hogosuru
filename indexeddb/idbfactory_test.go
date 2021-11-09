package indexeddb

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestIDBFactoryNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "iddb"); testingutils.AssertErr(t, err) {

		if factory, err := IDBFactoryNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "IDBFactory", factory.ConstructName_())
		}

	}

}

var methodsFactoryAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Open", "args": []interface{}{"dbtest"}, "type": "constructnamechecking", "resultattempt": "IDBOpenDBRequest"},
	{"method": "DeleteDatabase", "args": []interface{}{"dbtest"}, "type": "constructnamechecking", "resultattempt": "IDBOpenDBRequest"},
	{"method": "Databases", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Cmp", "args": []interface{}{2, 2}, "resultattempt": 0},
}

func TestFactoryMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "iddb"); testingutils.AssertErr(t, err) {

		if factory, err := IDBFactoryNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsFactoryAttempt {
				testingutils.InvokeCheck(t, factory, result)
			}

		}

	}
}
