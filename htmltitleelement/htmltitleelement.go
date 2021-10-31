package htmltitleelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmltitleelementinterface js.Value

//HtmlTemplatelement struct
type HtmlTitleElement struct {
	htmlelement.HtmlElement
}

type HtmlTitleElementFrom interface {
	HtmlTitleElement_() HtmlTitleElement
}

func (h HtmlTitleElement) HtmlTitleElement_() HtmlTitleElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltitleelementinterface, err = baseobject.Get(js.Global(), "HTMLTitleElement"); err != nil {
			htmltitleelementinterface = js.Undefined()
		}
		baseobject.Register(htmltitleelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltitleelementinterface
}

func New(d document.Document) (HtmlTitleElement, error) {
	var err error

	var h HtmlTitleElement
	var e element.Element

	if e, err = d.CreateElement("title"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTitleElement, error) {
	var h HtmlTitleElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTitleElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTitleElement, error) {
	var h HtmlTitleElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLTitleElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTitleElement) Text() (string, error) {
	return h.GetAttributeString("text")
}

func (h HtmlTitleElement) SetText(value string) error {
	return h.SetAttributeString("text", value)
}
