package htmlembedelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlembedelementinterface js.Value

//HtmlEmbedElement struct
type HtmlEmbedElement struct {
	htmlelement.HtmlElement
}

type HtmlEmbedElementFrom interface {
	HtmlEmbedElement_() HtmlEmbedElement
}

func (h HtmlEmbedElement) HtmlEmbedElement_() HtmlEmbedElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlembedelementinterface, err = baseobject.Get(js.Global(), "HTMLEmbedElement"); err != nil {
			htmlembedelementinterface = js.Undefined()
		}
		baseobject.Register(htmlembedelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlembedelementinterface
}

func New(d document.Document) (HtmlEmbedElement, error) {
	var err error

	var h HtmlEmbedElement
	var e element.Element

	if e, err = d.CreateElement("embed"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlEmbedElement, error) {
	var h HtmlEmbedElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlEmbedElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlEmbedElement, error) {
	var h HtmlEmbedElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlEmbedElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlEmbedElement) Height() (string, error) {
	return h.GetAttributeString("height")
}

func (h HtmlEmbedElement) SetHeight(value string) error {
	return h.SetAttributeString("height", value)
}

func (h HtmlEmbedElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlEmbedElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlEmbedElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlEmbedElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlEmbedElement) Width() (string, error) {
	return h.GetAttributeString("width")
}

func (h HtmlEmbedElement) SetWidth(value string) error {
	return h.SetAttributeString("width", value)
}
