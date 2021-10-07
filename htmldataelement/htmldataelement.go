package htmldataelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmldataelementinterface js.Value

//HtmlDataElement struct
type HtmlDataElement struct {
	htmlelement.HtmlElement
}

type HtmlDataElementFrom interface {
	HtmlDataElement_() HtmlDataElement
}

func (h HtmlDataElement) HtmlDataElement_() HtmlDataElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldataelementinterface, err = js.Global().GetWithErr("HTMLDataElement"); err != nil {
			htmldataelementinterface = js.Null()
		}
		baseobject.Register(htmldataelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldataelementinterface
}

func New(d document.Document) (HtmlDataElement, error) {
	var err error

	var h HtmlDataElement
	var e element.Element

	if e, err = d.CreateElement("data"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDataElement, error) {
	var h HtmlDataElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlDataElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDataElement, error) {
	var h HtmlDataElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlDataElement
}

func (h HtmlDataElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlDataElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}
