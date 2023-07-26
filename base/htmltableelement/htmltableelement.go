package htmltableelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlcollection"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/htmltablecaptionelement"
	"github.com/realPy/hogosuru/base/htmltablerowelement"
	"github.com/realPy/hogosuru/base/htmltablesectionelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltableelementinterface js.Value

// HtmlTableElement struct
type HtmlTableElement struct {
	htmlelement.HtmlElement
}

type HtmlTableElementFrom interface {
	HtmlTableElement_() HtmlTableElement
}

func (h HtmlTableElement) HtmlTableElement_() HtmlTableElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltableelementinterface, err = baseobject.Get(js.Global(), "HTMLTableElement"); err != nil {
			htmltableelementinterface = js.Undefined()
		}
		baseobject.Register(htmltableelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmltableelementinterface
}

func New(d document.Document) (HtmlTableElement, error) {
	var err error

	var h HtmlTableElement
	var e element.Element

	if e, err = d.CreateElement("table"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableElement, error) {
	var h HtmlTableElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableColElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableElement, error) {
	var h HtmlTableElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLTableColElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableElement) Caption() (htmltablecaptionelement.HtmlTableCaptionElement, error) {

	var err error
	var obj js.Value
	var caption htmltablecaptionelement.HtmlTableCaptionElement

	if obj, err = h.Get("caption"); err == nil {

		caption, err = htmltablecaptionelement.NewFromJSObject(obj)
	}

	return caption, err

}

func (h HtmlTableElement) SetCaption(caption htmltablecaptionelement.HtmlTableCaptionElement) error {

	return h.Set("caption", caption.JSObject())
}

func (h HtmlTableElement) getCollectionMethod(method string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get(method); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableElement) getElemMethod(method string) (htmltablesectionelement.HtmlTableSectionElement, error) {

	var err error
	var obj js.Value
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.Get(method); err == nil {

		elem, err = htmltablesectionelement.NewFromJSObject(obj)
	}

	return elem, err
}

func (h HtmlTableElement) Rows() (htmlcollection.HtmlCollection, error) {
	return h.getCollectionMethod("rows")

}

func (h HtmlTableElement) TBodies() (htmlcollection.HtmlCollection, error) {
	return h.getCollectionMethod("tBodies")
}

func (h HtmlTableElement) TFoot() (htmltablesectionelement.HtmlTableSectionElement, error) {
	return h.getElemMethod("tFoot")
}

func (h HtmlTableElement) THead() (htmltablesectionelement.HtmlTableSectionElement, error) {
	return h.getElemMethod("tHead")
}

func (h HtmlTableElement) CreateCaption() (htmltablecaptionelement.HtmlTableCaptionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablecaptionelement.HtmlTableCaptionElement

	if obj, err = h.Call("createCaption"); err == nil {
		elem, err = htmltablecaptionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTHead() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.Call("createTHead"); err == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTFoot() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.Call("createTFoot"); err == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteCaption() error {

	_, err := h.Call("deleteCaption")
	return err
}

func (h HtmlTableElement) DeleteTHead() error {

	_, err := h.Call("deleteTHead")
	return err
}

func (h HtmlTableElement) DeleteTFoot() error {

	_, err := h.Call("deleteTFoot")
	return err
}

func (h HtmlTableElement) InsertRow(index ...int) (htmltablerowelement.HtmlTableRowElement, error) {
	var obj js.Value
	var err error
	var elem htmltablerowelement.HtmlTableRowElement

	var arrayJS []interface{}

	if len(index) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(index[0]))
	}

	if obj, err = h.Call("insertRow", arrayJS...); err == nil {
		elem, err = htmltablerowelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteRow(index int) error {

	var err error
	_, err = h.Call("deleteRow", js.ValueOf(index))
	return err
}
