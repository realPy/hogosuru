package htmllinkelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/domtokenlist"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/stylesheet"
)

var singleton sync.Once

var htmllinklementinterface js.Value

//HtmlLinkElement struct
type HtmlLinkElement struct {
	htmlelement.HtmlElement
}

type HtmlLinkElementFrom interface {
	HtmlLinkElement() HtmlLinkElement
}

func (h HtmlLinkElement) HtmlLinkElement() HtmlLinkElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmllinklementinterface, err = js.Global().GetWithErr("HTMLLinkElement"); err != nil {
			htmllinklementinterface = js.Null()
		}
		baseobject.Register(htmllinklementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllinklementinterface
}

func New(d document.Document) (HtmlLinkElement, error) {
	var err error

	var h HtmlLinkElement
	var e element.Element

	if e, err = d.CreateElement("link"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLinkElement, error) {
	var h HtmlLinkElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlLinkElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLinkElement, error) {
	var h HtmlLinkElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlLinkElement
}

func (h HtmlLinkElement) As() (string, error) {
	return h.GetAttributeString("as")
}

func (h HtmlLinkElement) SetAs(value string) error {
	return h.SetAttributeString("as", value)
}

func (s HtmlLinkElement) Disabled() (bool, error) {
	return s.GetAttributeBool("disabled")
}

func (s HtmlLinkElement) SetDisabled(value bool) error {
	return s.SetAttributeBool("disabled", value)
}

func (h HtmlLinkElement) Media() (string, error) {
	return h.GetAttributeString("media")
}

func (h HtmlLinkElement) SetMedia(value string) error {
	return h.SetAttributeString("media", value)
}

func (h HtmlLinkElement) Href() (string, error) {
	return h.GetAttributeString("href")
}

func (h HtmlLinkElement) SetHref(value string) error {
	return h.SetAttributeString("href", value)
}

func (h HtmlLinkElement) Hreflang() (string, error) {
	return h.GetAttributeString("hreflang")
}

func (h HtmlLinkElement) SetHreflang(value string) error {
	return h.SetAttributeString("hreflang", value)
}

func (h HtmlLinkElement) ReferrerPolicy() (string, error) {
	return h.GetAttributeString("referrerPolicy")
}

func (h HtmlLinkElement) SetReferrerPolicy(value string) error {
	return h.SetAttributeString("referrerPolicy", value)
}

func (h HtmlLinkElement) Rel() (string, error) {
	return h.GetAttributeString("rel")
}

func (h HtmlLinkElement) SetRel(value string) error {
	return h.SetAttributeString("rel", value)
}

func (h HtmlLinkElement) RelList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj, err = h.JSObject().GetWithErr("relList"); err == nil {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (h HtmlLinkElement) Sheet() (stylesheet.StyleSheet, error) {
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

func (h HtmlLinkElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlLinkElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}
