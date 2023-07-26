package htmlbuttonelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`b= document.createElement("button")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLButtonElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "b"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLButtonElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Autofocus", "resultattempt": false},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Form", "type": "error", "resultattempt": ErrNoForm},
	{"method": "FormAction", "type": "contains", "resultattempt": "http://localhost"},
	{"method": "FormEncType", "resultattempt": ""},
	{"method": "FormMethod", "resultattempt": ""},
	{"method": "FormNoValidate", "resultattempt": false},
	{"method": "FormTarget", "resultattempt": ""},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
	{"method": "Name", "resultattempt": ""},
	{"method": "TabIndex", "resultattempt": 0},
	{"method": "Type", "resultattempt": "submit"},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "WillValidate", "resultattempt": true},
	{"method": "Value", "resultattempt": ""},
	{"method": "SetAutofocus", "args": []interface{}{true}, "gettermethod": "Autofocus", "resultattempt": true},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetFormAction", "args": []interface{}{"postdata"}, "gettermethod": "FormAction", "type": "contains", "resultattempt": "/postdata"},
	{"method": "SetFormEncType", "args": []interface{}{"application/x-www-form-urlencoded"}, "gettermethod": "FormEncType", "resultattempt": "application/x-www-form-urlencoded"},
	{"method": "SetFormMethod", "args": []interface{}{"post"}, "gettermethod": "FormMethod", "resultattempt": "post"},
	{"method": "SetFormNoValidate", "args": []interface{}{true}, "gettermethod": "FormNoValidate", "resultattempt": true},
	{"method": "SetFormTarget", "args": []interface{}{"_self"}, "gettermethod": "FormTarget", "resultattempt": "_self"},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "SetTabIndex", "args": []interface{}{777}, "gettermethod": "TabIndex", "resultattempt": 777},
	{"method": "SetType", "args": []interface{}{"reset"}, "gettermethod": "Type", "resultattempt": "reset"},
	{"method": "SetValue", "args": []interface{}{"value"}, "gettermethod": "Value", "resultattempt": "value"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "b"); testingutils.AssertErr(t, err) {

		if button, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, button, result)
			}

		}

	}
}
