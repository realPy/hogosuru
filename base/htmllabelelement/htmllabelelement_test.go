package htmllabelelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`l= document.createElement("label")
	i= document.createElement("input")
	l.appendChild(i)
	l.appendChild(document.createElement("input"))
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLLabelElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLLabelElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Control", "type": "constructnamechecking", "resultattempt": "HTMLInputElement"},
	{"method": "Form", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "HtmlFor", "resultattempt": ""},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, button, result)
			}

		}

	}
}
