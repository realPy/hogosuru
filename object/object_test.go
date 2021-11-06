package object

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	objnew= new Object()
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if o, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object Object]", o.ToString_())
	}

}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`
	objnew= new Object()
	`)

	if obj, err := baseobject.Get(js.Global(), "objnew"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Object]", o.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Keys", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "Values", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "Map", "type": "constructnamechecking", "resultattempt": "Map"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "Array"); testingutils.AssertErr(t, err) {

		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, o, result)
			}

		}

	}
}
