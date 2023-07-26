package htmldataelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`d= document.createElement("data")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLDataElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "d"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLDataElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Value", "resultattempt": ""},
	{"method": "SetValue", "args": []interface{}{"value"}, "gettermethod": "Value", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "d"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, button, result)
			}

		}

	}
}
