package htmlhtmlelement

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

var htmlhtmlelementinterface js.Value

// HtmlHtmlElement struct
type HtmlHtmlElement struct {
	htmlelement.HtmlElement
}

type HtmlHtmlElementFrom interface {
	HtmlHtmlElement_() HtmlHtmlElement
}

func (h HtmlHtmlElement) HtmlHtmlElement_() HtmlHtmlElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlhtmlelementinterface, err = baseobject.Get(js.Global(), "HTMLHtmlElement"); err != nil {
			htmlhtmlelementinterface = js.Undefined()
		}
		baseobject.Register(htmlhtmlelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlhtmlelementinterface
}

func New(d document.Document) (HtmlHtmlElement, error) {
	var err error

	var h HtmlHtmlElement
	var e element.Element

	if e, err = d.CreateElement("html"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHtmlElement, error) {
	var h HtmlHtmlElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlHtmlElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHtmlElement, error) {
	var h HtmlHtmlElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlHtmlElement
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return h, err
}
