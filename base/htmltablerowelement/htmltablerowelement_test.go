package htmltablerowelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`tr=document.createElement("tr")
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if tr, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableRowElement", tr.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "tr"); testingutils.AssertErr(t, err) {

		if tr, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableRowElement", tr.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "RowIndex", "resultattempt": -1},
	{"method": "SectionRowIndex", "resultattempt": -1},
	{"method": "Cells", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "tr"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertCell(t *testing.T) {

	baseobject.Eval(`tri=document.createElement("tr")
	`)

	if obj, err := baseobject.Get(js.Global(), "tri"); testingutils.AssertErr(t, err) {

		if tr, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if cell, err := tr.InsertCell(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTableCellElement", cell.ConstructName_())
				if cells, err := tr.Cells(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, cells.Length())
				}

			}
		}

	}
}

func TestDeleteCell(t *testing.T) {

	baseobject.Eval(`trd=document.createElement("tr")
	`)

	if obj, err := baseobject.Get(js.Global(), "trd"); testingutils.AssertErr(t, err) {

		if tr, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			tr.InsertCell()
			tr.InsertCell()
			if cells, err := tr.Cells(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, 2, cells.Length())

				testingutils.AssertErr(t, tr.DeleteCell(0))
				if cells2, err := tr.Cells(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, cells2.Length())
				}

			}
		}

	}
}
