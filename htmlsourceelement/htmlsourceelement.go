package htmlsourceelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlsourceelementinterface js.Value

//HtmlSourceElement struct
type HtmlSourceElement struct {
	htmlelement.HtmlElement
}

type HtmlSourceElementFrom interface {
	HtmlSourceElement_() HtmlSourceElement
}

func (h HtmlSourceElement) HtmlSourceElement_() HtmlSourceElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlsourceelementinterface, err = baseobject.Get(js.Global(), "HTMLSourceElement"); err != nil {
			htmlsourceelementinterface = js.Undefined()
		}
		baseobject.Register(htmlsourceelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlsourceelementinterface
}

func New(d document.Document) (HtmlSourceElement, error) {
	var err error

	var h HtmlSourceElement
	var e element.Element

	if e, err = d.CreateElement("source"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSourceElement, error) {
	var h HtmlSourceElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLSourceElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSourceElement, error) {
	var h HtmlSourceElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLSourceElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlSourceElement) Media() (string, error) {
	return h.GetAttributeString("media")
}

func (h HtmlSourceElement) SetMedia(value string) error {
	return h.SetAttributeString("media", value)
}

func (h HtmlSourceElement) Sizes() (string, error) {
	return h.GetAttributeString("sizes")
}

func (h HtmlSourceElement) SetSizes(value string) error {
	return h.SetAttributeString("sizes", value)
}

func (h HtmlSourceElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlSourceElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlSourceElement) SrcSet() (string, error) {
	return h.GetAttributeString("srcSet")
}

func (h HtmlSourceElement) SetSrcSet(value string) error {
	return h.SetAttributeString("srcSet", value)
}

func (h HtmlSourceElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlSourceElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}
