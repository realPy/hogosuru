package htmliframeelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`i= document.createElement("iframe")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLIFrameElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLIFrameElement", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AllowPaymentRequest", "resultattempt": false},
	{"method": "ContentDocument", "type": "error", "resultattempt": ErrNoContentDocument},
	{"method": "Height", "resultattempt": ""},
	{"method": "Src", "resultattempt": ""},
	{"method": "Name", "resultattempt": ""},
	{"method": "Width", "resultattempt": ""},
	{"method": "Srcdoc", "resultattempt": ""},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if iframe, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, iframe, result)
			}

		}

	}
}

//

var setterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SetAllowPaymentRequest", "args": []interface{}{true}, "gettermethod": "AllowPaymentRequest", "resultattempt": true},
	{"method": "SetHeight", "args": []interface{}{"value"}, "gettermethod": "Height", "resultattempt": "value"},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "SetName", "args": []interface{}{"value"}, "gettermethod": "Name", "resultattempt": "value"},
	{"method": "SetWidth", "args": []interface{}{"value"}, "gettermethod": "Width", "resultattempt": "value"},
	{"method": "SetSrcdoc", "args": []interface{}{"value"}, "gettermethod": "Srcdoc", "resultattempt": "value"},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if iframe, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, iframe, result)

			}

		}

	}

}
