package htmlfieldsetelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/validitystate"
)

var singleton sync.Once

var htmlfieldsetelementinterface js.Value

//HtmlFieldSetElement struct
type HtmlFieldSetElement struct {
	htmlelement.HtmlElement
}

type HtmlFieldSetElementFrom interface {
	HtmlFieldSetElement_() HtmlFieldSetElement
}

func (h HtmlFieldSetElement) HtmlFieldSetElement_() HtmlFieldSetElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlfieldsetelementinterface, err = baseobject.Get(js.Global(), "HTMLFieldSetElement"); err != nil {
			htmlfieldsetelementinterface = js.Undefined()
		}
		baseobject.Register(htmlfieldsetelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlfieldsetelementinterface
}

func New(d document.Document) (HtmlFieldSetElement, error) {
	var err error

	var h HtmlFieldSetElement
	var e element.Element

	if e, err = d.CreateElement("fieldset"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlFieldSetElement, error) {
	var h HtmlFieldSetElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlFieldSetElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlFieldSetElement, error) {
	var h HtmlFieldSetElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlFieldSetElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlFieldSetElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlFieldSetElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlFieldSetElement) Elements() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("elements"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlFieldSetElement) Form() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("form"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlFieldSetElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlFieldSetElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlFieldSetElement) Type() (string, error) {

	return h.GetAttributeString("type")
}

func (h HtmlFieldSetElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlFieldSetElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.Get("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlFieldSetElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlFieldSetElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlFieldSetElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlFieldSetElement) SetCustomValidity(message string) error {

	_, err := h.Call("setCustomValidity", js.ValueOf(message))
	return err
}
