package htmlformelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`f= document.createElement("form")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLFormElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLFormElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": ""},
	{"method": "Method", "resultattempt": "get"},
	{"method": "Target", "resultattempt": ""},
	{"method": "Action", "type": "contains", "resultattempt": "http://localhost"},
	{"method": "Encoding", "resultattempt": "application/x-www-form-urlencoded"},
	{"method": "Enctype", "resultattempt": "application/x-www-form-urlencoded"},
	{"method": "AcceptCharset", "resultattempt": ""},
	{"method": "Autocomplete", "resultattempt": "on"},
	{"method": "NoValidate", "resultattempt": false},
	{"method": "CheckValidity", "resultattempt": true},
	{"method": "ReportValidity", "resultattempt": true},
	{"method": "RequestSubmit", "type": "error", "resultattempt": nil},
	{"method": "Reset", "type": "error", "resultattempt": nil},
	{"method": "Submit", "type": "error", "resultattempt": nil},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "SetMethod", "args": []interface{}{"post"}, "gettermethod": "Method", "resultattempt": "post"},
	{"method": "SetTarget", "args": []interface{}{"thistarget"}, "gettermethod": "Target", "resultattempt": "thistarget"},
	{"method": "SetAction", "args": []interface{}{"/test"}, "gettermethod": "Action", "type": "contains", "resultattempt": "/test"},
	{"method": "SetEncoding", "args": []interface{}{"multipart/form-data"}, "gettermethod": "Encoding", "resultattempt": "multipart/form-data"},
	{"method": "SetEnctype", "args": []interface{}{"multipart/form-data"}, "gettermethod": "Enctype", "resultattempt": "multipart/form-data"},
	{"method": "SetAcceptCharset", "args": []interface{}{"UTF-8"}, "gettermethod": "AcceptCharset", "resultattempt": "UTF-8"},
	{"method": "SetAutocomplete", "args": []interface{}{"off"}, "gettermethod": "Autocomplete", "resultattempt": "off"},
	{"method": "SetNoValidate", "args": []interface{}{true}, "gettermethod": "NoValidate", "resultattempt": true},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if form, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, form, result)
			}

		}

	}
}
