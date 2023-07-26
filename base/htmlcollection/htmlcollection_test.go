package htmlcollection

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`ch= document.children
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ch"); testingutils.AssertErr(t, err) {

		if c, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLCollection", c.ConstructName_())
		}

	}
}

func TestItem(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ch"); testingutils.AssertErr(t, err) {

		if c, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if n, err := c.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "HTMLHtmlElement", n.(baseobject.BaseObject).ConstructName_())

			}
		}

	}
}
