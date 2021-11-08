package htmlselectelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmloptionelement"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`f=document.createElement("form")
	s= document.createElement("select")
	f.appendChild(s)
	o= document.createElement("option")
	o.value="t3st"
	s.appendChild(o)
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if selectobj, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLSelectElement", selectobj.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if selectobj, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLSelectElement", selectobj.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Autofocus", "resultattempt": false},
	{"method": "SetAutofocus", "args": []interface{}{true}, "gettermethod": "Autofocus", "resultattempt": true},
	{"method": "Disabled", "resultattempt": false},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "Length", "resultattempt": 1},
	{"method": "Name", "resultattempt": ""},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "Options", "type": "constructnamechecking", "resultattempt": "HTMLOptionsCollection"},
	{"method": "Multiple", "resultattempt": false},
	{"method": "SetMultiple", "args": []interface{}{true}, "gettermethod": "Multiple", "resultattempt": true},
	{"method": "Required", "resultattempt": false},
	{"method": "SetRequired", "args": []interface{}{true}, "gettermethod": "Required", "resultattempt": true},
	{"method": "SelectedIndex", "resultattempt": 0},
	{"method": "SetSelectedIndex", "args": []interface{}{0}, "gettermethod": "SelectedIndex", "resultattempt": 0},
	{"method": "SelectedOptions", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Size", "resultattempt": 0},
	{"method": "SetSize", "args": []interface{}{3}, "gettermethod": "Size", "resultattempt": 3},
	{"method": "Type", "resultattempt": "select-multiple"},
	{"method": "Value", "resultattempt": "t3st"},
	{"method": "SetValue", "args": []interface{}{"t3st"}, "gettermethod": "Value", "resultattempt": "t3st"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "WillValidate", "resultattempt": false},
	{"method": "ReportValidity", "resultattempt": true},
	{"method": "SetCustomValidity", "args": []interface{}{"hello"}, "type": "error", "resultattempt": nil},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "s"); testingutils.AssertErr(t, err) {

		if selectobj, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, selectobj, result)
			}

		}

	}
}

func TestAdd(t *testing.T) {
	baseobject.Eval(`
	sadd= document.createElement("select")
	`)

	if obj, err := baseobject.Get(js.Global(), "sadd"); testingutils.AssertErr(t, err) {

		if sadd, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if l, err := sadd.Length(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, 0, l)

				if option, err := htmloptionelement.Option("test"); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, sadd.Add(option))

					if l2, err := sadd.Length(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, 1, l2)

					}

				}
			}

		}

	}

}

func TestItem(t *testing.T) {
	baseobject.Eval(`
	sitem= document.createElement("select")
	oitem= document.createElement("option")
	oitem.name="t3stitem"
	sitem.appendChild(oitem)

	`)

	if obj, err := baseobject.Get(js.Global(), "sitem"); testingutils.AssertErr(t, err) {

		if sitem, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if item, err := sitem.Item(0); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLOptionElement", item.ConstructName_())
			}

		}

	}

}

func TestNamedItem(t *testing.T) {
	baseobject.Eval(`
	sitemn= document.createElement("select")
	oitemn= document.createElement("option")
	oitemn.id="t3stitem"
	sitemn.appendChild(oitemn)

	`)

	if obj, err := baseobject.Get(js.Global(), "sitemn"); testingutils.AssertErr(t, err) {

		if sitem, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if item, err := sitem.NamedItem("t3stitem"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLOptionElement", item.ConstructName_())
			}

		}

	}

}

func TestRemove(t *testing.T) {
	baseobject.Eval(`
	sitemr= document.createElement("select")
	oitemr= document.createElement("option")
	sitemr.appendChild(oitemr)
	`)

	if obj, err := baseobject.Get(js.Global(), "sitemr"); testingutils.AssertErr(t, err) {

		if sitemr, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if l, err := sitemr.Length(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, 1, l)

				testingutils.AssertErr(t, sitemr.Remove(0))
				if l2, err := sitemr.Length(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, 0, l2)

				}

			}

		}

	}

}
