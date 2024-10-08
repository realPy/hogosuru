package htmlquoteelement

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

var htmlquoteelementinterface js.Value

// HtmlQuoteElement struct
type HtmlQuoteElement struct {
	htmlelement.HtmlElement
}

type HtmlQuoteElementFrom interface {
	HtmlQuoteElement_() HtmlQuoteElement
}

func (h HtmlQuoteElement) HtmlQuoteElement_() HtmlQuoteElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlquoteelementinterface, err = baseobject.Get(js.Global(), "HTMLQuoteElement"); err != nil {
			htmlquoteelementinterface = js.Undefined()
		}
		baseobject.Register(htmlquoteelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlquoteelementinterface
}

func New(d document.Document) (HtmlQuoteElement, error) {
	var err error

	var h HtmlQuoteElement
	var e element.Element

	if e, err = d.CreateElement("q"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewBlockQuote(d document.Document) (HtmlQuoteElement, error) {
	var err error

	var h HtmlQuoteElement
	var e element.Element

	if e, err = d.CreateElement("blockquote"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlQuoteElement, error) {
	var h HtmlQuoteElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLQuoteElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlQuoteElement, error) {
	var h HtmlQuoteElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLQuoteElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlQuoteElement) Cite() (string, error) {

	return h.GetAttributeString("cite")
}

func (h HtmlQuoteElement) SetCite(value string) error {
	return h.SetAttributeString("cite", value)
}
