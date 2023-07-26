package htmltablecellelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`td= document.createElement("td")
	th= document.createElement("th")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if td, err := NewTd(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableCellElement", td.ConstructName_())
		}

		if th, err := NewTh(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableCellElement", th.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "td"); testingutils.AssertErr(t, err) {

		if td, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableCellElement", td.ConstructName_())
		}

	}

	if obj, err := baseobject.Get(js.Global(), "th"); testingutils.AssertErr(t, err) {

		if th, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableCellElement", th.ConstructName_())
		}

	}
}
