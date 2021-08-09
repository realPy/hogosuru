package htmlbaseelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlbaseelementinterface js.Value

//HtmlBaseElement struct
type HtmlBaseElement struct {
	htmlelement.HtmlElement
}

type HtmlBaseElementFrom interface {
	HtmlBaseElement() HtmlBaseElement
}

func (h HtmlBaseElement) HtmlBaseElement() HtmlBaseElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlbaseelementinterface, err = js.Global().GetWithErr("HTMLBaseElement"); err != nil {
			htmlbaseelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlbaseelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
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

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlBaseElement
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
