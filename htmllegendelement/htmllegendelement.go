package htmllegendelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmlformelement"
)

var singleton sync.Once

var htmllegendelementinterface js.Value

//HtmlLegendElement struct
type HtmlLegendElement struct {
	htmlelement.HtmlElement
}

type HtmlLegendElementFrom interface {
	HtmlLegendElement() HtmlLegendElement
}

func (h HtmlLegendElement) HtmlLegendElement() HtmlLegendElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmllegendelementinterface, err = js.Global().GetWithErr("HTMLLegendElement"); err != nil {
			htmllegendelementinterface = js.Null()
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

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLLegendElement
}

func (h HtmlLegendElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var formelem htmlformelement.HtmlFormElement

	if obj, err = h.JSObject().GetWithErr("form"); err == nil {

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
