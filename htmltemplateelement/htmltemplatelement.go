package htmltemplateelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/documentfragment"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmltemplateelementinterface js.Value

//HtmlTemplateElement struct
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

	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTemplateElement
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
