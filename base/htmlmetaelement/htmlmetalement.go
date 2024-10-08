package htmlmetaelement

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

var htmlmetaelementinterface js.Value

// HtmlMetaElement struct
type HtmlMetaElement struct {
	htmlelement.HtmlElement
}

type HtmlMetaElementFrom interface {
	HtmlMetaElement_() HtmlMetaElement
}

func (h HtmlMetaElement) HtmlMetaElement_() HtmlMetaElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlmetaelementinterface, err = baseobject.Get(js.Global(), "HTMLMetaElement"); err != nil {
			htmlmetaelementinterface = js.Undefined()
		}
		baseobject.Register(htmlmetaelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlmetaelementinterface
}

func New(d document.Document) (HtmlMetaElement, error) {
	var err error

	var h HtmlMetaElement
	var e element.Element

	if e, err = d.CreateElement("meta"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMetaElement, error) {
	var h HtmlMetaElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLMetaElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMetaElement, error) {
	var h HtmlMetaElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLMetaElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlMetaElement) Content() (string, error) {
	return h.GetAttributeString("content")
}

func (h HtmlMetaElement) SetContent(value string) error {
	return h.SetAttributeString("content", value)
}

func (h HtmlMetaElement) HttpEquiv() (string, error) {
	return h.GetAttributeString("httpEquiv")
}

func (h HtmlMetaElement) SetHttpEquiv(value string) error {
	return h.SetAttributeString("httpEquiv", value)
}

func (h HtmlMetaElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlMetaElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}
