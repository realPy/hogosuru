package htmlscriptelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`s= document.createElement("script")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if a, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLScriptElement", a.ConstructName_())
		}

	}

}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if a, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLScriptElement", a.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Type", "resultattempt": ""},
	{"method": "SetType", "args": []interface{}{"module"}, "gettermethod": "Type", "resultattempt": "module"},
	{"method": "Src", "resultattempt": ""},
	{"method": "SetSrc", "args": []interface{}{"value"}, "gettermethod": "Src", "type": "contains", "resultattempt": "/value"},
	{"method": "Async", "resultattempt": true},
	{"method": "SetAsync", "args": []interface{}{false}, "gettermethod": "Async", "resultattempt": false},
	{"method": "Defer", "resultattempt": false},
	{"method": "SetDefer", "args": []interface{}{true}, "gettermethod": "Defer", "resultattempt": true},
	{"method": "Text", "resultattempt": ""},
	{"method": "SetText", "args": []interface{}{"test"}, "gettermethod": "Text", "resultattempt": "test"},
	{"method": "NoModule", "resultattempt": false},
	{"method": "SetNoModule", "args": []interface{}{true}, "gettermethod": "NoModule", "resultattempt": true},
	{"method": "ReferrerPolicy", "resultattempt": ""},
	{"method": "SetReferrerPolicy", "args": []interface{}{"no-referrer"}, "gettermethod": "ReferrerPolicy", "resultattempt": "no-referrer"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, area, result)
			}

		}

	}
}
