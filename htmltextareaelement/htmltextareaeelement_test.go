package htmltextareaelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`f=document.createElement("form")
	ta= document.createElement("textarea")
	f.appendChild(ta)
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if ta, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTextAreaElement", ta.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ta"); testingutils.AssertErr(t, err) {

		if ta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTextAreaElement", ta.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AccessKey", "resultattempt": ""},
	{"method": "SetAccessKey", "args": []interface{}{"i"}, "gettermethod": "AccessKey", "resultattempt": "i"},
	{"method": "Autocapitalize", "resultattempt": ""},
	{"method": "SetAutocapitalize", "args": []interface{}{"words"}, "gettermethod": "Autocapitalize", "resultattempt": "words"},
	{"method": "Autocomplete", "resultattempt": ""},
	{"method": "SetAutocomplete", "args": []interface{}{"on"}, "gettermethod": "Autocomplete", "resultattempt": "on"},
	{"method": "Autofocus", "resultattempt": false},
	{"method": "SetAutofocus", "args": []interface{}{true}, "gettermethod": "Autofocus", "resultattempt": true},
	{"method": "Cols", "resultattempt": 20},
	{"method": "SetCols", "args": []interface{}{23}, "gettermethod": "Cols", "resultattempt": 23},
	{"method": "DefaultValue", "resultattempt": ""},
	{"method": "SetDefaultValue", "args": []interface{}{"testamazingworld"}, "gettermethod": "DefaultValue", "resultattempt": "testamazingworld"},
	{"method": "Disabled", "resultattempt": false},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "MaxLength", "resultattempt": -1},
	{"method": "SetMaxLength", "args": []interface{}{150}, "gettermethod": "MaxLength", "resultattempt": 150},
	{"method": "MinLength", "resultattempt": -1},
	{"method": "SetMinLength", "args": []interface{}{10}, "gettermethod": "MinLength", "resultattempt": 10},
	{"method": "Name", "resultattempt": ""},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "Placeholder", "resultattempt": ""},
	{"method": "SetPlaceholder", "args": []interface{}{"thisplaceholder"}, "gettermethod": "Placeholder", "resultattempt": "thisplaceholder"},
	{"method": "ReadOnly", "resultattempt": false},
	{"method": "SetReadOnly", "args": []interface{}{true}, "gettermethod": "ReadOnly", "resultattempt": true},
	{"method": "Required", "resultattempt": false},
	{"method": "SetRequired", "args": []interface{}{true}, "gettermethod": "Required", "resultattempt": true},
	{"method": "Rows", "resultattempt": 2},
	{"method": "SetRows", "args": []interface{}{10}, "gettermethod": "Rows", "resultattempt": 10},
	{"method": "SelectionStart", "resultattempt": 0},
	{"method": "SetSelectionStart", "args": []interface{}{1}, "gettermethod": "SelectionStart", "resultattempt": 1},
	{"method": "SelectionEnd", "resultattempt": 1},
	{"method": "SetSelectionEnd", "args": []interface{}{8}, "gettermethod": "SelectionEnd", "resultattempt": 8},
	//My chrome return none, CI return default "forward"... so ensure no errror
	{"method": "SelectionDirection", "type": "error", "resultattempt": nil},
	{"method": "SetSelectionDirection", "args": []interface{}{"forward"}, "gettermethod": "SelectionDirection", "resultattempt": "forward"},
	{"method": "TabIndex", "resultattempt": 0},
	{"method": "SetTabIndex", "args": []interface{}{10}, "gettermethod": "TabIndex", "resultattempt": 10},
	{"method": "TextLength", "resultattempt": 16},
	{"method": "Type", "resultattempt": "textarea"},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
	{"method": "Value", "resultattempt": "testamazingworld"},
	{"method": "SetValue", "args": []interface{}{"great"}, "gettermethod": "Value", "resultattempt": "great"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "WillValidate", "resultattempt": false},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
	{"method": "Wrap", "resultattempt": ""},
	{"method": "SetWrap", "args": []interface{}{"hard"}, "gettermethod": "Wrap", "resultattempt": "hard"},
	{"method": "CheckValidity", "resultattempt": true},
	{"method": "ReportValidity", "resultattempt": true},
	{"method": "SetCustomValidity", "args": []interface{}{"hello"}, "type": "error", "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ta"); testingutils.AssertErr(t, err) {

		if ta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, ta, result)
			}

		}

	}
}

func TestSetRangeText(t *testing.T) {
	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if ta, err := New(doc); testingutils.AssertErr(t, err) {
			ta.SetValue("hello tiny world")
			testingutils.AssertErr(t, ta.SetRangeText("big", 6, 10))
			v, _ := ta.Value()
			testingutils.AssertExpect(t, "hello big world", v)
		}

	}

}
