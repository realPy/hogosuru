package htmltableelement

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`t=document.createElement("table")
	c=document.createElement("caption")
	c.testContent="title"
	t.appendChild(c)
	thead=document.createElement("thead")
	t.appendChild(thead)
	r=document.createElement("row")
	t.insertRow(r)
	tfoot=document.createElement("tfoot")
	t.appendChild(tfoot)

	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if table, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLTableElement", table.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTableElement", table.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Rows", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Caption", "type": "constructnamechecking", "resultattempt": "HTMLTableCaptionElement"},
	{"method": "TBodies", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "TFoot", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "THead", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "CreateCaption", "type": "constructnamechecking", "resultattempt": "HTMLTableCaptionElement"},
	{"method": "CreateTHead", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
	{"method": "CreateTFoot", "type": "constructnamechecking", "resultattempt": "HTMLTableSectionElement"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "t"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, table, result)
			}

		}

	}
}

func TestInsertRow(t *testing.T) {

	baseobject.Eval(`tir=document.createElement("table")
	`)

	if obj, err := baseobject.Get(js.Global(), "tir"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if r, err := table.InsertRow(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTableRowElement", r.ConstructName_())
				if cr, err := table.Rows(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, cr.Length())
				}

			}
		}

	}
}

func TestDeleteRow(t *testing.T) {

	baseobject.Eval(`tdr=document.createElement("table")
	`)

	if obj, err := baseobject.Get(js.Global(), "tdr"); testingutils.AssertErr(t, err) {

		if table, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			table.InsertRow()
			table.InsertRow()
			if cr, err := table.Rows(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, 2, cr.Length())

				testingutils.AssertErr(t, table.DeleteRow(0))
				if cr2, err := table.Rows(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, cr2.Length())
				}

			}

		}

	}
}

func TestDeleteCaption(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if table, err := New(doc); testingutils.AssertErr(t, err) {

			table.CreateCaption()
			_, err := table.Caption()
			testingutils.AssertExpect(t, false, errors.Is(err, baseobject.ErrUndefinedValue))
			testingutils.AssertErr(t, table.DeleteCaption())
			_, err = table.Caption()
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
		}

	}

}

func TestDeleteTFoot(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if table, err := New(doc); testingutils.AssertErr(t, err) {

			table.CreateTFoot()
			_, err := table.TFoot()
			testingutils.AssertExpect(t, false, errors.Is(err, baseobject.ErrUndefinedValue))
			testingutils.AssertErr(t, table.DeleteTFoot())
			_, err = table.TFoot()
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
		}

	}

}

func TestDeleteTHead(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if table, err := New(doc); testingutils.AssertErr(t, err) {

			table.CreateTHead()
			_, err := table.THead()
			testingutils.AssertExpect(t, false, errors.Is(err, baseobject.ErrUndefinedValue))
			testingutils.AssertErr(t, table.DeleteTHead())
			_, err = table.THead()
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
		}

	}

}
