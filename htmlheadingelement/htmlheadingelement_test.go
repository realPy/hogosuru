package htmlheadingelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`h= document.createElement("h1")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if h, err := NewH1(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH2(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH3(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH4(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

		if h, err := NewH5(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}
		if h, err := NewH6(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "h"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLHeadingElement", h.ConstructName_())
		}

	}
}
