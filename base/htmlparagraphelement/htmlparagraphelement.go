package htmlparagraphelement

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

var htmlparagraphelementinterface js.Value

// HtmlParagraphElement struct
type HtmlParagraphElement struct {
	htmlelement.HtmlElement
}

type HtmlParagraphElementFrom interface {
	HtmlParagraphElement_() HtmlParagraphElement
}

func (h HtmlParagraphElement) HtmlParagraphElement_() HtmlParagraphElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlparagraphelementinterface, err = baseobject.Get(js.Global(), "HTMLParagraphElement"); err != nil {
			htmlparagraphelementinterface = js.Undefined()
		}
		baseobject.Register(htmlparagraphelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlparagraphelementinterface
}

func New(d document.Document) (HtmlParagraphElement, error) {
	var err error

	var h HtmlParagraphElement
	var e element.Element

	if e, err = d.CreateElement("p"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLParagraphElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLParagraphElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
