package htmltimeelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`t= document.createElement("time")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if ti, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTimeElement", ti.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if ti, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTimeElement", ti.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "DateTime", "resultattempt": ""},
	{"method": "SetDateTime", "args": []interface{}{"13h"}, "gettermethod": "DateTime", "resultattempt": "13h"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if selectobj, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, selectobj, result)
			}

		}

	}
}
