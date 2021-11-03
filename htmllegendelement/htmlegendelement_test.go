package htmllegendelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`l=document.createElement("legend")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLLegendElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLLegendElement", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AccessKey", "resultattempt": ""},
	{"method": "Form", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, area, result)
			}

		}

	}
}

var setterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SetAccessKey", "args": []interface{}{"i"}, "gettermethod": "AccessKey", "resultattempt": "i"},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, area, result)

			}

		}

	}

}