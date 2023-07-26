package htmlinputelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/date"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/filelist"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/nodelist"
	"github.com/realPy/hogosuru/base/validitystate"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlinputelementinterface js.Value

// HtmlInputElement struct
type HtmlInputElement struct {
	htmlelement.HtmlElement
}

type HtmlInputElementFrom interface {
	HtmlInputElement_() HtmlInputElement
}

func (h HtmlInputElement) HtmlInputElement_() HtmlInputElement {
	return h
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlinputelementinterface, err = baseobject.Get(js.Global(), "HTMLInputElement"); err != nil {
			htmlinputelementinterface = js.Undefined()
		}
		baseobject.Register(htmlinputelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmlinputelementinterface
}

func New(d document.Document) (HtmlInputElement, error) {
	var err error

	var h HtmlInputElement
	var e element.Element

	if e, err = d.CreateElement("input"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlInputElement, error) {
	var h HtmlInputElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlInputElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

//Properties related to the parent form

func (h HtmlInputElement) Form() (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = h.Get("form"); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}
	return elem, err
}

func (h HtmlInputElement) FormAction() (string, error) {
	return h.GetAttributeString("formAction")
}

func (h HtmlInputElement) SetFormAction(value string) error {
	return h.SetAttributeString("formAction", value)
}

func (h HtmlInputElement) FormEncType() (string, error) {
	return h.GetAttributeString("formEncType")
}

func (h HtmlInputElement) SetFormEncType(value string) error {
	return h.SetAttributeString("formEncType", value)
}

func (h HtmlInputElement) FormMethod() (string, error) {
	return h.GetAttributeString("formMethod")
}

func (h HtmlInputElement) SetFormMethod(value string) error {
	return h.SetAttributeString("formMethod", value)
}

func (h HtmlInputElement) FormNoValidate() (bool, error) {
	return h.GetAttributeBool("formNoValidate")
}

func (h HtmlInputElement) SetFormNoValidate(value bool) error {
	return h.SetAttributeBool("formNoValidate", value)
}

func (h HtmlInputElement) FormTarget() (string, error) {
	return h.GetAttributeString("formTarget")
}

func (h HtmlInputElement) SetFormTarget(value string) error {
	return h.SetAttributeString("formTarget", value)
}

//Properties that apply to any type of input element that is not hidden

func (h HtmlInputElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlInputElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlInputElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlInputElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlInputElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlInputElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlInputElement) Autofocus() (bool, error) {
	return h.GetAttributeBool("autofocus")
}

func (h HtmlInputElement) SetAutofocus(value bool) error {
	return h.SetAttributeBool("autofocus", value)
}

func (h HtmlInputElement) Required() (bool, error) {
	return h.GetAttributeBool("required")
}

func (h HtmlInputElement) SetRequired(value bool) error {
	return h.SetAttributeBool("required", value)
}

func (h HtmlInputElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlInputElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}

func (h HtmlInputElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.Get("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlInputElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlInputElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

// Properties that apply only to elements of type "checkbox" or "radio"
func (h HtmlInputElement) Checked() (bool, error) {
	return h.GetAttributeBool("checked")
}

func (h HtmlInputElement) SetChecked(value bool) error {
	return h.SetAttributeBool("checked", value)
}

func (h HtmlInputElement) DefaultChecked() (bool, error) {
	return h.GetAttributeBool("defaultChecked")
}

func (h HtmlInputElement) SetDefaultChecked(value bool) error {
	return h.SetAttributeBool("defaultChecked", value)
}

func (h HtmlInputElement) Indeterminate() (bool, error) {
	return h.GetAttributeBool("indeterminate")
}

func (h HtmlInputElement) SetIndeterminate(value bool) error {
	return h.SetAttributeBool("indeterminate", value)
}

// Properties that apply only to elements of type "image"

func (h HtmlInputElement) Alt() (string, error) {
	return h.GetAttributeString("alt")
}

func (h HtmlInputElement) SetAlt(value string) error {
	return h.SetAttributeString("alt", value)
}

// MSN  is wrong return type is int not string

func (h HtmlInputElement) Height() (int, error) {
	return h.GetAttributeInt("height")
}

func (h HtmlInputElement) SetHeight(value int) error {
	return h.SetAttributeInt("height", value)
}

func (h HtmlInputElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlInputElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlInputElement) Width() (int, error) {
	return h.GetAttributeInt("width")
}

func (h HtmlInputElement) SetWidth(value int) error {
	return h.SetAttributeInt("width", value)
}

// Properties that apply only to elements of type "file"

func (h HtmlInputElement) Accept() (string, error) {
	return h.GetAttributeString("accept")
}

func (h HtmlInputElement) SetAccept(value string) error {
	return h.SetAttributeString("accept", value)
}

func (h HtmlInputElement) Files() (filelist.FileList, error) {
	var files js.Value
	var err error
	if files, err = h.Get("files"); err == nil {
		return filelist.NewFromJSObject(files)
	}
	return filelist.FileList{}, err
}

// Properties that apply only to text/number-containing or elements

func (h HtmlInputElement) Autocomplete() (string, error) {
	return h.GetAttributeString("autocomplete")
}

func (h HtmlInputElement) SetAutocomplete(value string) error {
	return h.SetAttributeString("autocomplete", value)
}

func (h HtmlInputElement) Max() (string, error) {
	return h.GetAttributeString("max")
}

func (h HtmlInputElement) SetMax(value string) error {
	return h.SetAttributeString("max", value)
}

func (h HtmlInputElement) MaxLength() (int, error) {
	return h.GetAttributeInt("maxLength")
}

func (h HtmlInputElement) SetMaxLength(value int) error {
	return h.SetAttributeInt("maxLength", value)
}

func (h HtmlInputElement) Min() (string, error) {
	return h.GetAttributeString("min")
}

func (h HtmlInputElement) SetMin(value string) error {
	return h.SetAttributeString("min", value)
}

func (h HtmlInputElement) MinLength() (int, error) {
	return h.GetAttributeInt("minLength")
}

func (h HtmlInputElement) SetMinLength(value int) error {
	return h.SetAttributeInt("minLength", value)
}

func (h HtmlInputElement) Pattern() (string, error) {
	return h.GetAttributeString("pattern")
}

func (h HtmlInputElement) SetPattern(value string) error {
	return h.SetAttributeString("pattern", value)
}

func (h HtmlInputElement) Placeholder() (string, error) {
	return h.GetAttributeString("placeholder")
}

func (h HtmlInputElement) SetPlaceholder(value string) error {
	return h.SetAttributeString("placeholder", value)
}

func (h HtmlInputElement) ReadOnly() (bool, error) {
	return h.GetAttributeBool("readOnly")
}

func (h HtmlInputElement) SetReadOnly(value bool) error {
	return h.SetAttributeBool("readOnly", value)
}

func (h HtmlInputElement) SelectionStart() (int, error) {
	return h.GetAttributeInt("selectionStart")
}

func (h HtmlInputElement) SetSelectionStart(value int) error {
	return h.SetAttributeInt("selectionStart", value)
}

func (h HtmlInputElement) SelectionEnd() (int, error) {
	return h.GetAttributeInt("selectionEnd")
}

func (h HtmlInputElement) SetSelectionEnd(value int) error {
	return h.SetAttributeInt("selectionEnd", value)
}

func (h HtmlInputElement) SelectionDirection() (string, error) {
	return h.GetAttributeString("selectionDirection")
}

func (h HtmlInputElement) SetSelectionDirection(value string) error {
	return h.SetAttributeString("selectionDirection", value)
}

func (h HtmlInputElement) Size() (int, error) {
	return h.GetAttributeInt("size")
}

func (h HtmlInputElement) SetSize(value int) error {
	return h.SetAttributeInt("size", value)
}

//  Properties not yet categorized

func (h HtmlInputElement) DefaultValue() (string, error) {
	return h.GetAttributeString("defaultValue")
}

func (h HtmlInputElement) SetDefaultValue(value string) error {
	return h.SetAttributeString("defaultValue", value)
}

func (h HtmlInputElement) DirName() (string, error) {
	return h.GetAttributeString("dirName")
}

func (h HtmlInputElement) SetDirName(value string) error {
	return h.SetAttributeString("dirName", value)
}

func (h HtmlInputElement) List() (htmlelement.HtmlElement, error) {
	var obj js.Value
	var err error
	var elem htmlelement.HtmlElement
	if obj, err = h.Get("list"); err == nil {

		elem, err = htmlelement.NewFromJSObject(obj)
	}
	return elem, err
}

func (h HtmlInputElement) Multiple() (bool, error) {
	return h.GetAttributeBool("multiple")
}

func (h HtmlInputElement) SetMultiple(value bool) error {
	return h.SetAttributeBool("multiple", value)
}

func (h HtmlInputElement) Labels() (nodelist.NodeList, error) {
	var obj js.Value
	var err error
	var arr nodelist.NodeList
	if obj, err = h.Get("labels"); err == nil {

		arr, err = nodelist.NewFromJSObject(obj)
	}
	return arr, err
}

func (h HtmlInputElement) Step() (string, error) {
	return h.GetAttributeString("step")
}

func (h HtmlInputElement) SetStep(value string) error {
	return h.SetAttributeString("step", value)
}

func (h HtmlInputElement) ValueAsDate() (date.Date, error) {
	var obj js.Value
	var err error
	var arr date.Date
	if obj, err = h.Get("valueAsDate"); err == nil {

		arr, err = date.NewFromJSObject(obj)
	}
	return arr, err
}

func (h HtmlInputElement) ValueAsNumber() (float64, error) {
	return h.GetAttributeDouble("valueAsNumber")
}
