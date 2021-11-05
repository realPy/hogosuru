package htmlprogresselement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	l= document.createElement("label")
	p= document.createElement("progress")
	l.appendChild(p)`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLProgressElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if meter, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLProgressElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Max", "resultattempt": float64(1)},
	{"method": "SetMax", "args": []interface{}{float64(10.0)}, "gettermethod": "Max", "resultattempt": float64(10.0)},
	{"method": "Position", "resultattempt": float64(-1)},
	{"method": "Value", "resultattempt": float64(0.0)},
	{"method": "SetValue", "args": []interface{}{float64(5.3)}, "gettermethod": "Value", "resultattempt": float64(5.3)},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
