package htmlsourceelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`s=document.createElement("source")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if source, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLSourceElement", source.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if source, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLSourceElement", source.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Media", "resultattempt": ""},
	{"method": "Sizes", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "SrcSet", "resultattempt": ""},
	{"method": "Type", "resultattempt": ""},

	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetSizes", "args": []interface{}{"no"}, "gettermethod": "Sizes", "resultattempt": "no"},
	{"method": "SetSrcSet", "args": []interface{}{"small.jpg 1x, large.jpg 2x"}, "gettermethod": "SrcSet", "resultattempt": "small.jpg 1x, large.jpg 2x"},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if anchor, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, anchor, result)
			}

		}

	}
}
