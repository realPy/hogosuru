package htmltableelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmltablecaptionelement"
	"github.com/realPy/hogosuru/htmltablerowelement"
	"github.com/realPy/hogosuru/htmltablesectionelement"
)

var singleton sync.Once

var htmltableelementinterface js.Value

//HtmlTableElement struct
type HtmlTableElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltableelementinterface, err = js.Global().GetWithErr("HTMLTableElement"); err != nil {
			htmltableelementinterface = js.Null()
		}

	})

	baseobject.Register(htmltableelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
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

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTableColElement
}

func (h HtmlTableElement) Caption() (string, error) {
	return h.GetAttributeString("caption")
}

func (h HtmlTableElement) SetCaption(value string) error {
	return h.SetAttributeString("caption", value)
}

func (h HtmlTableElement) getCollectionMethod(method string) (htmlcollection.HTMLCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = h.JSObject().GetWithErr(method); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlTableElement) getElemMethod(method string) (htmltablesectionelement.HtmlTableSectionElement, error) {

	var err error
	var obj js.Value
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.JSObject().GetWithErr(method); err == nil {

		elem, err = htmltablesectionelement.NewFromJSObject(obj)
	}

	return elem, err
}

func (h HtmlTableElement) Rows() (htmlcollection.HTMLCollection, error) {
	return h.getCollectionMethod("rows")

}

func (h HtmlTableElement) TBodies() (htmlcollection.HTMLCollection, error) {
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

	if obj, err = h.JSObject().CallWithErr("createCaption"); err == nil {
		elem, err = htmltablecaptionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTHead() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.JSObject().CallWithErr("createTHead"); err == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) CreateTFoot() (htmltablesectionelement.HtmlTableSectionElement, error) {
	var obj js.Value
	var err error
	var elem htmltablesectionelement.HtmlTableSectionElement

	if obj, err = h.JSObject().CallWithErr("createTFoot"); err == nil {
		elem, err = htmltablesectionelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteCaption() error {

	_, err := h.JSObject().CallWithErr("deleteCaption")
	return err
}

func (h HtmlTableElement) DeleteTHead() error {

	_, err := h.JSObject().CallWithErr("deleteTHead")
	return err
}

func (h HtmlTableElement) DeleteTFoot() error {

	_, err := h.JSObject().CallWithErr("deleteTFoot")
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

	if obj, err = h.JSObject().CallWithErr("insertRow", arrayJS...); err == nil {
		elem, err = htmltablerowelement.NewFromJSObject(obj)

	}
	return elem, err
}

func (h HtmlTableElement) DeleteRow(index int) error {

	var err error
	_, err = h.JSObject().CallWithErr("deleteRow", js.ValueOf(index))
	return err
}
