package htmltemplateelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`t= document.createElement("template")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if template, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTemplateElement", template.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if template, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTemplateElement", template.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Content", "type": "constructnamechecking", "resultattempt": "DocumentFragment"},
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
