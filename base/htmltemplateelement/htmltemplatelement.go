package htmltemplateelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/documentfragment"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmltemplateelementinterface js.Value

// HtmlTemplateElement struct
type HtmlTemplateElement struct {
	htmlelement.HtmlElement
}

type HtmlTemplateElementFrom interface {
	HtmlTemplateElement_() HtmlTemplateElement
}

func (h HtmlTemplateElement) HtmlTemplateElement_() HtmlTemplateElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltemplateelementinterface, err = baseobject.Get(js.Global(), "HTMLTemplateElement"); err != nil {
			htmltemplateelementinterface = js.Undefined()
		}
		baseobject.Register(htmltemplateelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltemplateelementinterface
}

func New(d document.Document) (HtmlTemplateElement, error) {
	var err error

	var h HtmlTemplateElement
	var e element.Element

	if e, err = d.CreateElement("template"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTemplateElement, error) {
	var h HtmlTemplateElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTemplateElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTemplateElement, error) {
	var h HtmlTemplateElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLTemplateElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTemplateElement) Content() (documentfragment.DocumentFragment, error) {
	var err error
	var obj js.Value
	var fragment documentfragment.DocumentFragment

	if obj, err = h.Get("content"); err == nil {

		fragment, err = documentfragment.NewFromJSObject(obj)
	}

	return fragment, err
}
