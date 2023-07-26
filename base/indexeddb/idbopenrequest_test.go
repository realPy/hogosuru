package indexeddb

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestIDBOpenDBRequestNewFromJSObject(t *testing.T) {
	baseobject.Eval(`iddb=window.indexedDB
	opendbrequest=iddb.open("openrequest")
	`)

	if obj, err := baseobject.Get(js.Global(), "opendbrequest"); testingutils.AssertErr(t, err) {

		if opendbrequest, err := IDBOpenDBRequestNewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "IDBOpenDBRequest", opendbrequest.ConstructName_())
		}

	}
	baseobject.Eval(`
	opendbrequest=iddb.deleteDatabase("openrequest")
	`)

}
