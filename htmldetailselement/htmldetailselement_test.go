package htmldetailselement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`d= document.createElement("details")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLDetailsElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "d"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLDetailsElement", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Open", "resultattempt": false},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "d"); testingutils.AssertErr(t, err) {

		if base, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, base, result)
			}

		}

	}
}

var setterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SetOpen", "args": []interface{}{true}, "gettermethod": "Open", "resultattempt": true},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "d"); testingutils.AssertErr(t, err) {

		if base, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, base, result)

			}

		}

	}

}
