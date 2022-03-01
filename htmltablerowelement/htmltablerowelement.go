package htmltablerowelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmltablecellelement"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltablerowelementinterface js.Value

//HtmlTableRowElement struct
type HtmlTableRowElement struct {
	htmlelement.HtmlElement
}

type HtmlTableRowElementFrom interface {
	HtmlTableRowElement_() HtmlTableRowElement
}

func (h HtmlTableRowElement) HtmlTableRowElement_() HtmlTableRowElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablerowelementinterface, err = baseobject.Get(js.Global(), "HTMLTableRowElement"); err != nil {
			htmltablerowelementinterface = js.Undefined()
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

	if hci := GetInterface(); !hci.IsUndefined() {
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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLTableRowElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableRowElement) Cells() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("cells"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
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

	if obj, err = h.Call("insertCell", arrayJS...); err == nil {
		elem, err = htmltablecellelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableRowElement) DeleteCell(index int) error {

	var err error
	_, err = h.Call("deleteCell", js.ValueOf(index))

	return err
}
