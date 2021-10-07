package htmlstyleelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/stylesheet"
)

var singleton sync.Once

var htmlstylelementinterface js.Value

//HtmlStyleElement struct
type HtmlStyleElement struct {
	htmlelement.HtmlElement
}

type HtmlStyleElementFrom interface {
	HtmlStyleElement_() HtmlStyleElement
}

func (h HtmlStyleElement) HtmlStyleElement_() HtmlStyleElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlstylelementinterface, err = js.Global().GetWithErr("HTMLStyleElement"); err != nil {
			htmlstylelementinterface = js.Null()
		}
		baseobject.Register(htmlstylelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlstylelementinterface
}

func New(d document.Document) (HtmlStyleElement, error) {
	var err error

	var h HtmlStyleElement
	var e element.Element

	if e, err = d.CreateElement("style"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlStyleElement, error) {
	var h HtmlStyleElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLStyleElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlStyleElement, error) {
	var h HtmlStyleElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLStyleElement
}

func (h HtmlStyleElement) Media() (string, error) {
	return h.GetAttributeString("media")
}

func (h HtmlStyleElement) SetMedia(value string) error {
	return h.SetAttributeString("media", value)
}

func (h HtmlStyleElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlStyleElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlStyleElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlStyleElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlStyleElement) Sheet() (stylesheet.StyleSheet, error) {
	var err error
	var obj js.Value
	var s stylesheet.StyleSheet
	if obj, err = h.JSObject().GetWithErr("sheet"); err == nil {

		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {
			s, err = stylesheet.NewFromJSObject(obj)
		}
	}
	return s, err
}

func (h HtmlStyleElement) Scoped() (bool, error) {
	return h.GetAttributeBool("scoped")
}

func (h HtmlStyleElement) SetScoped(value bool) error {
	return h.SetAttributeBool("scoped", value)
}
