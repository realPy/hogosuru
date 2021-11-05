package htmltablesectionelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`tbody=document.createElement("tbody")
	tfoot=document.createElement("tfoot")
	thead=document.createElement("thead")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if tr, err := NewTBody(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}
		if tr, err := NewTFoot(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}
		if tr, err := NewTHead(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableSectionElement", tr.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "tbody"); testingutils.AssertErr(t, err) {

		if tbody, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableSectionElement", tbody.ConstructName_())
		}

	}

	if obj, err := baseobject.Get(js.Global(), "thead"); testingutils.AssertErr(t, err) {

		if thead, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableSectionElement", thead.ConstructName_())
		}

	}

	if obj, err := baseobject.Get(js.Global(), "tfoot"); testingutils.AssertErr(t, err) {

		if tfoot, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableSectionElement", tfoot.ConstructName_())
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Rows", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "tbody"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertRow(t *testing.T) {

	baseobject.Eval(`tbodyi=document.createElement("tbody")
	`)

	if obj, err := baseobject.Get(js.Global(), "tbodyi"); testingutils.AssertErr(t, err) {

		if tbody, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if r, err := tbody.InsertRow(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTableRowElement", r.ConstructName_())
				if cr, err := tbody.Rows(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, cr.Length())
				}

			}
		}

	}
}

func TestDeleteRow(t *testing.T) {

	baseobject.Eval(`tbodyr=document.createElement("tbody")
	`)

	if obj, err := baseobject.Get(js.Global(), "tbodyr"); testingutils.AssertErr(t, err) {

		if tbody, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			tbody.InsertRow()
			tbody.InsertRow()
			if rows, err := tbody.Rows(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, 2, rows.Length())

				testingutils.AssertErr(t, tbody.DeleteRow(0))
				if rows, err := tbody.Rows(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, rows.Length())
				}

			}

		}

	}
}
