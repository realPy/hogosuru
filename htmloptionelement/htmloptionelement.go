package htmloptionelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmloptionelementinterface js.Value

//HtmlOptionElement struct
type HtmlOptionElement struct {
	htmlelement.HtmlElement
}

type HtmlOptionElementFrom interface {
	HtmlOptionElement_() HtmlOptionElement
}

func (h HtmlOptionElement) HtmlOptionElement_() HtmlOptionElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmloptionelementinterface, err = js.Global().GetWithErr("HTMLOptionElement"); err != nil {
			htmloptionelementinterface = js.Null()
		}
		baseobject.Register(htmloptionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmloptionelementinterface
}

func New(d document.Document) (HtmlOptionElement, error) {
	var err error

	var h HtmlOptionElement
	var e element.Element

	if e, err = d.CreateElement("option"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlOptionElement, error) {
	var h HtmlOptionElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLOptionElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlOptionElement, error) {
	var h HtmlOptionElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLOptionElement
}

func (h HtmlOptionElement) DefaultSelected() (bool, error) {
	return h.GetAttributeBool("defaultSelected")
}

func (h HtmlOptionElement) SetDefaultSelected(value bool) error {
	return h.SetAttributeBool("defaultSelected", value)
}

func (h HtmlOptionElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlOptionElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlOptionElement) Form() (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = h.JSObject().GetWithErr("form"); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}
	return elem, err
}

func (h HtmlOptionElement) Index() (int, error) {
	return h.GetAttributeInt("index")
}

func (h HtmlOptionElement) Label() (string, error) {
	return h.GetAttributeString("label")
}

func (h HtmlOptionElement) Selected() (bool, error) {
	return h.GetAttributeBool("selected")
}

func (h HtmlOptionElement) SetSelected(value bool) error {
	return h.SetAttributeBool("selected", value)
}

func (h HtmlOptionElement) Text() (string, error) {

	return h.GetAttributeString("text")
}

func (h HtmlOptionElement) SetText(value string) error {
	return h.SetAttributeString("text", value)
}

func (h HtmlOptionElement) Value() (string, error) {

	return h.GetAttributeString("value")
}

func (h HtmlOptionElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}
