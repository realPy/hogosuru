package namednodemap

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/attr"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`b=document.createElement("button")
	b.setAttribute("hello","world")
	b.setAttributeNS("name","high","low")
	listattr=b.attributes
	`)

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "listattr"); testingutils.AssertErr(t, err) {
		if namednodemap, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object NamedNodeMap]", namednodemap.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Item", "args": []interface{}{0}, "type": "constructnamechecking", "resultattempt": "Attr"},
	{"method": "GetNamedItem", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "Attr"},
	{"method": "GetNamedItemNS", "args": []interface{}{"name", "high"}, "type": "constructnamechecking", "resultattempt": "Attr"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "listattr"); testingutils.AssertErr(t, err) {

		if namednodemap, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, namednodemap, result)
			}

		}

	}
}

func TestSetNamedItem(t *testing.T) {
	baseobject.Eval(`b1=document.createElement("button")
	listattr1=b1.attributes
	attr1=document.createAttribute("hello");

	`)
	if obj, err := baseobject.Get(js.Global(), "listattr1"); testingutils.AssertErr(t, err) {

		if namednodemap1, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if objattr, err := baseobject.Get(js.Global(), "attr1"); testingutils.AssertErr(t, err) {

				if attr, err := attr.NewFromJSObject(objattr); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, namednodemap1.SetNamedItem(attr))
					if item, err := namednodemap1.Item(0); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "[object Attr]", item.ToString_())
					}

				}

			}

		}
	}

}

func TestRemoveNamedItem(t *testing.T) {
	baseobject.Eval(`br=document.createElement("button")
	br.setAttribute("hello","world")
	listattrr=br.attributes
	`)
	if obj, err := baseobject.Get(js.Global(), "listattrr"); testingutils.AssertErr(t, err) {

		if namednodemap, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, namednodemap.RemoveNamedItem("hello"))
			_, err := namednodemap.Item(0)
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))

		}
	}

}

func TestSetNamedItemNS(t *testing.T) {
	baseobject.Eval(`bns1=document.createElement("button")
	listattrns1=bns1.attributes
	attrns1=document.createAttributeNS("namespace","hello")

	`)
	if obj, err := baseobject.Get(js.Global(), "listattrns1"); testingutils.AssertErr(t, err) {

		if namednodemap, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if objattr, err := baseobject.Get(js.Global(), "attrns1"); testingutils.AssertErr(t, err) {

				if attr, err := attr.NewFromJSObject(objattr); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, namednodemap.SetNamedItemNS(attr))
					if item, err := namednodemap.Item(0); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "[object Attr]", item.ToString_())
					}

				}

			}

		}
	}

}

func TestRemoveNamedItemNS(t *testing.T) {
	baseobject.Eval(`brns=document.createElement("button")
	brns.setAttributeNS("namespace","hello","world")
	listattrrns=brns.attributes
	`)
	if obj, err := baseobject.Get(js.Global(), "listattrrns"); testingutils.AssertErr(t, err) {

		if namednodemap, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, namednodemap.RemoveNamedItemNS("namespace", "hello"))
			_, err := namednodemap.Item(0)
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))

		}
	}

}
