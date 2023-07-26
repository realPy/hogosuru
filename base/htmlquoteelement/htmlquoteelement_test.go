package htmlquoteelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	q= document.createElement("blockquote")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLQuoteElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "q"); testingutils.AssertErr(t, err) {

		if meter, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLQuoteElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Cite", "resultattempt": ""},
	{"method": "SetCite", "args": []interface{}{"test"}, "gettermethod": "Cite", "type": "contains", "resultattempt": "/test"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "q"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
