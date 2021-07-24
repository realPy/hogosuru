package htmlbuttonelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/htmlformelement"
	"github.com/realPy/hogosuru/nodelist"
	"github.com/realPy/hogosuru/validitystate"
)

var singleton sync.Once

var htmlbuttonelementinterface js.Value

//HtmlButtonElement struct
type HtmlButtonElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlbuttonelementinterface, err = js.Global().GetWithErr("HTMLButtonElement"); err != nil {
			htmlbuttonelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlbuttonelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlbuttonelementinterface
}

func New(d document.Document) (HtmlButtonElement, error) {
	var err error

	var h HtmlButtonElement
	var e element.Element

	if e, err = d.CreateElement("button"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlButtonElement, error) {
	var h HtmlButtonElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlButtonElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlButtonElement, error) {
	var h HtmlButtonElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlButtonElement
}

func (h HtmlButtonElement) Autofocus() (bool, error) {
	return h.GetAttributeBool("autofocus")
}

func (h HtmlButtonElement) SetAutofocus(value bool) error {
	return h.SetAttributeBool("autofocus", value)
}

func (h HtmlButtonElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlButtonElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlButtonElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var f htmlformelement.HtmlFormElement
	if obj, err = h.JSObject().GetWithErr("form"); err == nil {

		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {
			f, err = htmlformelement.NewFromJSObject(obj)
		}
	}
	return f, err
}

func (h HtmlButtonElement) FormAction() (string, error) {
	return h.GetAttributeString("formAction")
}

func (h HtmlButtonElement) SetFormAction(value string) error {
	return h.SetAttributeString("formAction", value)
}

func (h HtmlButtonElement) FormEncType() (string, error) {
	return h.GetAttributeString("formEncType")
}

func (h HtmlButtonElement) SetFormEncType(value string) error {
	return h.SetAttributeString("formEncType", value)
}

func (h HtmlButtonElement) FormMethod() (string, error) {
	return h.GetAttributeString("formMethod")
}

func (h HtmlButtonElement) SetFormMethod(value string) error {
	return h.SetAttributeString("formMethod", value)
}

func (h HtmlButtonElement) FormNoValidate() (bool, error) {
	return h.GetAttributeBool("formNoValidate")
}

func (h HtmlButtonElement) SetFormNoValidate(value bool) error {
	return h.SetAttributeBool("formNoValidate", value)
}

func (h HtmlButtonElement) FormTarget() (string, error) {
	return h.GetAttributeString("formTarget")
}

func (h HtmlButtonElement) SetFormTarget(value string) error {
	return h.SetAttributeString("formTarget", value)
}

func (h HtmlButtonElement) Labels() (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = h.JSObject().GetWithErr("labels"); err == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}

	return nlist, err
}

func (h HtmlButtonElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlButtonElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlButtonElement) TabIndex() (int, error) {
	return h.GetAttributeInt("tabIndex")
}

func (h HtmlButtonElement) SetIndex(value int) error {
	return h.SetAttributeInt("tabIndex", value)
}

func (h HtmlButtonElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlButtonElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlButtonElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.JSObject().GetWithErr("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlButtonElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlButtonElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlButtonElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlButtonElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}
