package htmliframelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmliframelementinterface js.Value

//HtmlHeadingElement struct
type HtmlIFrameElement struct {
	htmlelement.HtmlElement
}

type HtmlIFrameElementFrom interface {
	HtmlIFrameElement_() HtmlIFrameElement
}

func (h HtmlIFrameElement) HtmlIFrameElement_() HtmlIFrameElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmliframelementinterface, err = baseobject.Get(js.Global(), "HTMLIFrameElement"); err != nil {
			htmliframelementinterface = js.Undefined()
		}

		baseobject.Register(htmliframelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmliframelementinterface
}

func New(d document.Document) (HtmlIFrameElement, error) {
	var err error

	var h HtmlIFrameElement
	var e element.Element

	if e, err = d.CreateElement("iframe"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlIFrameElement, error) {
	var h HtmlIFrameElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlIFrameElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlIFrameElement, error) {
	var h HtmlIFrameElement

	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlIFrameElement
}

func (h HtmlIFrameElement) AllowPaymentRequest() (bool, error) {
	return h.GetAttributeBool("allowPaymentRequest")
}

func (h HtmlIFrameElement) SetAllowPaymentRequest(value bool) error {
	return h.SetAttributeBool("allowPaymentRequest", value)
}

func (h HtmlIFrameElement) ContentDocument() (document.Document, error) {
	var err error
	var obj js.Value
	var doc document.Document

	if obj, err = h.Get("contentDocument"); err == nil {

		doc, err = document.NewFromJSObject(obj)
	}

	return doc, err
}

func (h HtmlIFrameElement) Height() (string, error) {
	return h.GetAttributeString("height")
}

func (h HtmlIFrameElement) SetHeight(value string) error {
	return h.SetAttributeString("height", value)
}

func (h HtmlIFrameElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlIFrameElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlIFrameElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlIFrameElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlIFrameElement) Srcdoc() (string, error) {
	return h.GetAttributeString("srcdoc")
}

func (h HtmlIFrameElement) SetSrcdoc(value string) error {
	return h.SetAttributeString("srcdoc", value)
}

func (h HtmlIFrameElement) Width() (string, error) {
	return h.GetAttributeString("width")
}

func (h HtmlIFrameElement) SetWidth(value string) error {
	return h.SetAttributeString("width", value)
}
