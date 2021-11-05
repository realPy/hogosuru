package htmltitleelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`t= document.createElement("title")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if ti, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTitleElement", ti.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if ti, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTitleElement", ti.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Text", "resultattempt": ""},
	{"method": "SetText", "args": []interface{}{"mytitle"}, "gettermethod": "Text", "resultattempt": "mytitle"},
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
