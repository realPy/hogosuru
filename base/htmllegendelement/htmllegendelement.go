package htmllegendelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/htmlformelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmllegendelementinterface js.Value

// HtmlLegendElement struct
type HtmlLegendElement struct {
	htmlelement.HtmlElement
}

type HtmlLegendElementFrom interface {
	HtmlLegendElement_() HtmlLegendElement
}

func (h HtmlLegendElement) HtmlLegendElement_() HtmlLegendElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmllegendelementinterface, err = baseobject.Get(js.Global(), "HTMLLegendElement"); err != nil {
			htmllegendelementinterface = js.Undefined()
		}
		baseobject.Register(htmllegendelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllegendelementinterface
}

func New(d document.Document) (HtmlLegendElement, error) {
	var err error

	var h HtmlLegendElement
	var e element.Element

	if e, err = d.CreateElement("legend"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLegendElement, error) {
	var h HtmlLegendElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLLegendElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLegendElement, error) {
	var h HtmlLegendElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLLegendElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLegendElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var formelem htmlformelement.HtmlFormElement

	if obj, err = h.Get("form"); err == nil {

		formelem, err = htmlformelement.NewFromJSObject(obj)
	}

	return formelem, err
}

func (h HtmlLegendElement) AccessKey() (string, error) {
	return h.GetAttributeString("accessKey")
}

func (h HtmlLegendElement) SetAccessKey(value string) error {
	return h.SetAttributeString("accessKey", value)
}
