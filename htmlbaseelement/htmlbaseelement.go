package htmlbaseelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlbaseelementinterface js.Value

//HtmlBaseElement struct
type HtmlBaseElement struct {
	htmlelement.HtmlElement
}

type HtmlBaseElementFrom interface {
	HtmlBaseElement_() HtmlBaseElement
}

func (h HtmlBaseElement) HtmlBaseElement_() HtmlBaseElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlbaseelementinterface, err = baseobject.Get(js.Global(), "HTMLBaseElement"); err != nil {
			htmlbaseelementinterface = js.Undefined()
		}
		baseobject.Register(htmlbaseelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlbaseelementinterface
}

func New(d document.Document) (HtmlBaseElement, error) {
	var err error

	var h HtmlBaseElement
	var e element.Element

	if e, err = d.CreateElement("base"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlBaseElement, error) {
	var h HtmlBaseElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {

		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlBaseElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlBaseElement, error) {
	var h HtmlBaseElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlBaseElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlBaseElement) Href() (string, error) {
	return h.GetAttributeString("href")
}

func (h HtmlBaseElement) SetHref(value string) error {
	return h.SetAttributeString("href", value)
}

func (h HtmlBaseElement) Target() (string, error) {
	return h.GetAttributeString("target")
}

func (h HtmlBaseElement) SetTarget(value string) error {
	return h.SetAttributeString("target", value)
}
