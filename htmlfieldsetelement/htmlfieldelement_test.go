package htmlfieldsetelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`f= document.createElement("fieldset")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLFieldSetElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLFieldSetElement", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Disabled", "resultattempt": false},
	{"method": "Elements", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Form", "type": "error", "resultattempt": ErrNoForm},
	{"method": "Name", "resultattempt": ""},
	{"method": "Type", "resultattempt": "fieldset"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
	{"method": "WillValidate", "resultattempt": false},
	{"method": "ReportValidity", "resultattempt": true},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, button, result)
			}

		}

	}
}

var setterAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "SetCustomValidity", "args": []interface{}{"hello"}, "type": "error", "resultattempt": nil},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, button, result)

			}

		}

	}

}
