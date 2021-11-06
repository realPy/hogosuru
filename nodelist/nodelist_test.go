package nodelist

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	i1=document.createElement("input")
	i1.name="toto"
	i2=document.createElement("input")
	i2.name="toto"
	document.body.appendChild(i1)
	document.body.appendChild(i2)
	list=document.getElementsByName("toto")
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {
		if rect, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object NodeList]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "HTMLInputElement"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {

		if image, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, image, result)
			}

		}

	}
}
