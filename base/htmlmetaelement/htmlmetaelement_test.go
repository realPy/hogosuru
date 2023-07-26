package htmlmetaelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`m= document.createElement("meta")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if meta, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLMetaElement", meta.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "m"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLMetaElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Content", "resultattempt": ""},
	{"method": "SetContent", "args": []interface{}{"test"}, "gettermethod": "Content", "resultattempt": "test"},
	{"method": "HttpEquiv", "resultattempt": ""},
	{"method": "SetHttpEquiv", "args": []interface{}{"test"}, "gettermethod": "HttpEquiv", "resultattempt": "test"},
	{"method": "Name", "resultattempt": ""},
	{"method": "SetName", "args": []interface{}{"test"}, "gettermethod": "Name", "resultattempt": "test"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "m"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
