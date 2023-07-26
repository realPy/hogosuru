package htmltablecolelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`col= document.createElement("col")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if col, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableColElement", col.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "col"); testingutils.AssertErr(t, err) {

		if col, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableColElement", col.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Span", "resultattempt": 1},
	{"method": "SetSpan", "args": []interface{}{10}, "gettermethod": "Span", "resultattempt": 10},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "col"); testingutils.AssertErr(t, err) {

		if selectobj, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, selectobj, result)
			}

		}

	}
}
