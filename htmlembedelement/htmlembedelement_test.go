package htmlembedelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`e= document.createElement("embed")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLEmbedElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "e"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLEmbedElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Height", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "Type", "resultattempt": ""},
	{"method": "Width", "resultattempt": ""},
	{"method": "SetHeight", "args": []interface{}{"value"}, "gettermethod": "Height", "resultattempt": "value"},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "SetType", "args": []interface{}{"value"}, "gettermethod": "Type", "resultattempt": "value"},
	{"method": "SetWidth", "args": []interface{}{"value"}, "gettermethod": "Width", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "e"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, button, result)
			}

		}

	}
}
