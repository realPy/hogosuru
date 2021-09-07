package htmltablerowelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmltablecellelement"
)

var singleton sync.Once

var htmltablerowelementinterface js.Value

//HtmlTableRowElement struct
type HtmlTableRowElement struct {
	htmlelement.HtmlElement
}

type HtmlTableRowElementFrom interface {
	HtmlTableRowElement() HtmlTableRowElement
}

func (h HtmlTableRowElement) HtmlTableRowElement() HtmlTableRowElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablerowelementinterface, err = js.Global().GetWithErr("HTMLTableRowElement"); err != nil {
			htmltablerowelementinterface = js.Null()
		}
		baseobject.Register(htmltablerowelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablerowelementinterface
}

func New(d document.Document) (HtmlTableRowElement, error) {
	var err error

	var h HtmlTableRowElement
	var e element.Element

	if e, err = d.CreateElement("tr"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableRowElement, error) {
	var h HtmlTableRowElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableRowElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableRowElement, error) {
	var h HtmlTableRowElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTableRowElement
}

func (h HtmlTableRowElement) Cells(method string) (htmlelement.HtmlElement, error) {

	var err error
	var obj js.Value
	var elem htmlelement.HtmlElement

	if obj, err = h.JSObject().GetWithErr("cells"); err == nil {

		elem, err = htmlelement.NewFromJSObject(obj)
	}

	return elem, err
}

func (h HtmlTableRowElement) RowIndex() (int, error) {
	return h.GetAttributeInt("rowIndex")
}

func (h HtmlTableRowElement) SectionRowIndex() (int, error) {
	return h.GetAttributeInt("sectionRowIndex")
}

func (h HtmlTableRowElement) InsertCell(index ...int) (htmltablecellelement.HtmlTableCellElement, error) {
	var obj js.Value
	var err error
	var elem htmltablecellelement.HtmlTableCellElement
	var arrayJS []interface{}

	if len(index) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(index[0]))
	}

	if obj, err = h.JSObject().CallWithErr("insertCell", arrayJS...); err == nil {
		elem, err = htmltablecellelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableRowElement) DeleteCell(index int) error {

	var err error
	_, err = h.JSObject().CallWithErr("deleteCell", js.ValueOf(index))

	return err
}
