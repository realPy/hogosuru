package htmlstyleelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`style=document.createElement("style")
	style.type="text/css"
	style.textContent="p { color: #26b72b; }"
	document.head.appendChild(style)
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if source, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLStyleElement", source.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {

		if source, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLStyleElement", source.ConstructName_())
		}
	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Sheet", "type": "constructnamechecking", "resultattempt": "CSSStyleSheet"},
	{"method": "Media", "resultattempt": ""},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Type", "resultattempt": "text/css"},
	{"method": "SetMedia", "args": []interface{}{"print"}, "gettermethod": "Media", "resultattempt": "print"},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetType", "args": []interface{}{"mytype"}, "gettermethod": "Type", "resultattempt": "mytype"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {

		if anchor, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, anchor, result)
			}

		}

	}
}
