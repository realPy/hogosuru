package htmlbaseelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`b= document.createElement("base")
	b.href="https://myuser:mypass@www.test.com:444?q=123#tag"
	b.target="thistarget"
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLBaseElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "b"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLBaseElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Href", "resultattempt": "https://myuser:mypass@www.test.com:444/?q=123#tag"},
	{"method": "Target", "resultattempt": "thistarget"},
	{"method": "SetHref", "args": []interface{}{"http://pp:ss@www.noone.com:444?q=456#nosecure"}, "gettermethod": "Href", "resultattempt": "http://pp:ss@www.noone.com:444/?q=456#nosecure"},
	{"method": "SetTarget", "args": []interface{}{"yes"}, "gettermethod": "Target", "resultattempt": "yes"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "b"); testingutils.AssertErr(t, err) {

		if base, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, base, result)
			}

		}

	}
}
