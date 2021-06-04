package htmlinputelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/filelist"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/validitystate"
)

var singleton sync.Once

var htmlinputelementinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HtmlInputElement struct
type HtmlInputElement struct {
	htmlelement.HtmlElement
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var htmlinputelementinstance JSInterface
		var err error
		if htmlinputelementinstance.objectInterface, err = js.Global().GetWithErr("HTMLInputElement"); err == nil {
			htmlinputelementinterface = &htmlinputelementinstance
		}
	})

	return htmlinputelementinterface
}

func New() (HtmlInputElement, error) {

	var h HtmlInputElement

	if hci := GetJSInterface(); hci != nil {
		h.BaseObject = h.SetObject(hci.objectInterface.New())
		return h, nil
	}
	return h, ErrNotImplemented
}

func NewFromElement(elem element.Element) (HtmlInputElement, error) {
	var h HtmlInputElement
	var err error

	if ai := GetJSInterface(); ai != nil {
		if elem.BaseObject.JSObject().InstanceOf(ai.objectInterface) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlInputElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlInputElement, error) {
	var h HtmlInputElement

	if hei := GetJSInterface(); hei != nil {
		if obj.InstanceOf(hei.objectInterface) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlInputElement
}

func (h HtmlInputElement) getAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {

		valueStr = obj.String()
	}

	return valueStr, err

}
func (h HtmlInputElement) setAttributeString(attribute string, value string) error {

	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (h HtmlInputElement) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}

func (h HtmlInputElement) setAttributeBool(attribute string, value bool) error {

	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

//Properties related to the parent form

func (h HtmlInputElement) Form() (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = h.JSObject().GetWithErr("form"); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}
	return elem, err
}

func (h HtmlInputElement) FormAction() (string, error) {
	return h.getAttributeString("formAction")
}

func (h HtmlInputElement) SetFormAction(value string) error {
	return h.setAttributeString("formAction", value)
}

func (h HtmlInputElement) FormEncType() (string, error) {
	return h.getAttributeString("formEncType")
}

func (h HtmlInputElement) SetFormEncType(value string) error {
	return h.setAttributeString("formEncType", value)
}

func (h HtmlInputElement) FormMethod() (string, error) {
	return h.getAttributeString("formMethod")
}

func (h HtmlInputElement) SetFormMethod(value string) error {
	return h.setAttributeString("formMethod", value)
}

func (h HtmlInputElement) FormNoValidate() (bool, error) {
	return h.getAttributeBool("formNoValidate")
}

func (h HtmlInputElement) SetFormNoValidate(value bool) error {
	return h.setAttributeBool("formNoValidate", value)
}

func (h HtmlInputElement) FormTarget() (string, error) {
	return h.getAttributeString("formTarget")
}

func (h HtmlInputElement) SetFormTarget(value string) error {
	return h.setAttributeString("formTarget", value)
}

//Properties that apply to any type of input element that is not hidden

func (h HtmlInputElement) Name() (string, error) {
	return h.getAttributeString("name")
}

func (h HtmlInputElement) SetName(value string) error {
	return h.setAttributeString("name", value)
}

func (h HtmlInputElement) Type() (string, error) {
	return h.getAttributeString("type")
}

func (h HtmlInputElement) SetType(value string) error {
	return h.setAttributeString("type", value)
}

func (h HtmlInputElement) Disable() (bool, error) {
	return h.getAttributeBool("disable")
}

func (h HtmlInputElement) SetDisable(value bool) error {
	return h.setAttributeBool("disable", value)
}

func (h HtmlInputElement) Autofocus() (bool, error) {
	return h.getAttributeBool("autofocus")
}

func (h HtmlInputElement) SetAutofocus(value bool) error {
	return h.setAttributeBool("autofocus", value)
}

func (h HtmlInputElement) Required() (bool, error) {
	return h.getAttributeBool("required")
}

func (h HtmlInputElement) SetRequired(value bool) error {
	return h.setAttributeBool("required", value)
}

func (h HtmlInputElement) Value() (string, error) {
	return h.getAttributeString("value")
}

func (h HtmlInputElement) SetValue(value string) error {
	return h.setAttributeString("value", value)
}

func (h HtmlInputElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.JSObject().GetWithErr("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlInputElement) ValidationMessage() (string, error) {
	return h.getAttributeString("validationMessage")
}

func (h HtmlInputElement) WillValidate() (bool, error) {
	return h.getAttributeBool("willValidate")
}

// Properties that apply only to elements of type "checkbox" or "radio"
func (h HtmlInputElement) Checked() (bool, error) {
	return h.getAttributeBool("checked")
}

func (h HtmlInputElement) SetChecked(value bool) error {
	return h.setAttributeBool("checked", value)
}

func (h HtmlInputElement) DefaultChecked() (bool, error) {
	return h.getAttributeBool("defaultChecked")
}

func (h HtmlInputElement) SetDefaultChecked(value bool) error {
	return h.setAttributeBool("defaultChecked", value)
}

func (h HtmlInputElement) Indeterminate() (bool, error) {
	return h.getAttributeBool("indeterminate")
}

func (h HtmlInputElement) SetIndeterminate(value bool) error {
	return h.setAttributeBool("indeterminate", value)
}

// Properties that apply only to elements of type "file"

func (h HtmlInputElement) Files() (filelist.FileList, error) {
	var files js.Value
	var err error
	if files, err = h.JSObject().GetWithErr("files"); err == nil {
		return filelist.NewFromJSObject(files)
	}
	return filelist.FileList{}, err
}
