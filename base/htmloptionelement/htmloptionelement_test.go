package htmloptionelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	f=document.createElement("form")
	s= document.createElement("select")
	o= document.createElement("option")
	s.appendChild(o)
	f.appendChild(s)
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLOptionElement", b.ConstructName_())
		}

	}
}

func TestOption(t *testing.T) {

	if o, err := Option("test"); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "HTMLOptionElement", o.ConstructName_())
	}

}
func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "o"); testingutils.AssertErr(t, err) {

		if meter, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLOptionElement", meter.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "DefaultSelected", "resultattempt": false},
	{"method": "SetDefaultSelected", "args": []interface{}{true}, "gettermethod": "DefaultSelected", "resultattempt": true},
	{"method": "Disabled", "resultattempt": false},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "Index", "resultattempt": 0},
	{"method": "Label", "resultattempt": ""},
	{"method": "SetLabel", "args": []interface{}{"test"}, "gettermethod": "Label", "resultattempt": "test"},
	{"method": "Selected", "resultattempt": true},
	{"method": "SetSelected", "args": []interface{}{false}, "gettermethod": "Selected", "resultattempt": false},
	{"method": "Text", "resultattempt": ""},
	{"method": "SetText", "args": []interface{}{"test"}, "gettermethod": "Text", "resultattempt": "test"},
	{"method": "Value", "resultattempt": "test"},
	{"method": "SetValue", "args": []interface{}{"t3st"}, "gettermethod": "Value", "resultattempt": "t3st"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "o"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
