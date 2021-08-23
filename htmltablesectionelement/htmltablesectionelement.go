package htmltablesectionelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmltablerowelement"
)

var singleton sync.Once

var htmltablesectionelementinterface js.Value

//HtmlTableRowElement struct
type HtmlTableSectionElement struct {
	htmlelement.HtmlElement
}

type HtmlTableSectionElementFrom interface {
	HtmlTableSectionElement() HtmlTableSectionElement
}

func (h HtmlTableSectionElement) HtmlTableSectionElement() HtmlTableSectionElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablesectionelementinterface, err = js.Global().GetWithErr("HTMLTableSectionElement"); err != nil {
			htmltablesectionelementinterface = js.Null()
		}
		baseobject.Register(htmltablesectionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmltablesectionelementinterface
}

func NewTBody(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("tbody"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewTHead(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("thead"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewTFoot(d document.Document) (HtmlTableSectionElement, error) {
	var err error

	var h HtmlTableSectionElement
	var e element.Element

	if e, err = d.CreateElement("tfoot"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableSectionElement, error) {
	var h HtmlTableSectionElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableSectionElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableSectionElement, error) {
	var h HtmlTableSectionElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTableSectionElement
}

func (h HtmlTableSectionElement) Rows() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.JSObject().GetWithErr("rows"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableSectionElement) InsertRow(index ...int) (htmltablerowelement.HtmlTableRowElement, error) {
	var obj js.Value
	var err error
	var elem htmltablerowelement.HtmlTableRowElement

	var arrayJS []interface{}

	if len(index) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(index[0]))
	}

	if obj, err = h.JSObject().CallWithErr("insertRow", arrayJS...); err == nil {
		elem, err = htmltablerowelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableSectionElement) DeleteRow(index int) error {

	var err error
	_, err = h.JSObject().CallWithErr("deleteRow", js.ValueOf(index))
	return err
}