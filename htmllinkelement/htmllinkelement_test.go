package htmllinkelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`l=document.createElement("link")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLLinkElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLLinkElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "As", "resultattempt": ""},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Media", "resultattempt": ""},
	{"method": "Href", "resultattempt": ""},
	{"method": "Hreflang", "resultattempt": ""},
	{"method": "ReferrerPolicy", "resultattempt": ""},
	{"method": "Rel", "resultattempt": ""},
	{"method": "RelList", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "Sizes", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "Type", "resultattempt": ""},
	{"method": "SetAs", "args": []interface{}{"font"}, "gettermethod": "As", "resultattempt": "font"},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetHref", "args": []interface{}{"myFont.woff2"}, "gettermethod": "Href", "type": "contains", "resultattempt": "/myFont.woff2"},
	{"method": "SetHreflang", "args": []interface{}{"lang"}, "gettermethod": "Hreflang", "resultattempt": "lang"},
	{"method": "SetReferrerPolicy", "args": []interface{}{"no-referrer"}, "gettermethod": "ReferrerPolicy", "resultattempt": "no-referrer"},
	{"method": "SetRel", "args": []interface{}{"stylesheet"}, "gettermethod": "Rel", "resultattempt": "stylesheet"},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if anchor, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, anchor, result)
			}

		}

	}
}
