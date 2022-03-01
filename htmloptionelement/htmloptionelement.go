package htmloptionelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

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
		if htmloptionelementinterface, err = baseobject.Get(js.Global(), "HTMLOptionElement"); err != nil {
			htmloptionelementinterface = js.Undefined()
		}
		baseobject.Register(htmloptionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmloptionelementinterface
}

func Option(text string, opts ...interface{}) (HtmlOptionElement, error) {

	var err error
	var h HtmlOptionElement
	var arrayjs []interface{}

	arrayjs = append(arrayjs, js.ValueOf(text))

	for opt := range opts {
		arrayjs = append(arrayjs, js.ValueOf(opt))

	}

	if hci, err := baseobject.Get(js.Global(), "Option"); err == nil {
		var obj js.Value
		if obj, err = baseobject.New(hci, arrayjs...); err == nil {

			h, err = NewFromJSObject(obj)
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
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

	if hci := GetInterface(); !hci.IsUndefined() {
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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLOptionElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
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

	if obj, err = h.Get("form"); err == nil {

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

func (h HtmlOptionElement) SetLabel(value string) error {
	return h.SetAttributeString("label", value)
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
