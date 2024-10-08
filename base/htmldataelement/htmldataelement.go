package htmldataelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmldataelementinterface js.Value

// HtmlDataElement struct
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
		if htmldataelementinterface, err = baseobject.Get(js.Global(), "HTMLDataElement"); err != nil {
			htmldataelementinterface = js.Undefined()
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

	if hci := GetInterface(); !hci.IsUndefined() {

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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlDataElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlDataElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlDataElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}
