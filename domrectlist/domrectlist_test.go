package domrectlist

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`d=document.createElement("div")
	document.body.appendChild(d)
	rectlist=d.getClientRects()
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "rectlist"); testingutils.AssertErr(t, err) {
		if rect, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DOMRectList]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "DOMRect"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "rectlist"); testingutils.AssertErr(t, err) {

		if image, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, image, result)
			}

		}

	}
}
