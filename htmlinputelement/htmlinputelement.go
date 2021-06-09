package htmlinputelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/date"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/filelist"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/nodelist"
	"github.com/realPy/hogosuru/validitystate"
)

var singleton sync.Once

var htmlinputelementinterface js.Value

//HtmlInputElement struct
type HtmlInputElement struct {
	htmlelement.HtmlElement
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlinputelementinterface, err = js.Global().GetWithErr("HTMLInputElement"); err != nil {
			htmlinputelementinterface = js.Null()
		}

	})

	return htmlinputelementinterface
}

func New() (HtmlInputElement, error) {

	var h HtmlInputElement

	if hci := GetInterface(); !hci.IsNull() {
		h.BaseObject = h.SetObject(hci.New())
		return h, nil
	}
	return h, ErrNotImplemented
}

func NewFromElement(elem element.Element) (HtmlInputElement, error) {
	var h HtmlInputElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

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

func (h HtmlInputElement) getAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}

	return result, err
}

func (h HtmlInputElement) setAttributeInt(attribute string, value int) error {
	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (h HtmlInputElement) getAttributeDouble(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var result float64

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Float()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}

	return result, err
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

// Properties that apply only to elements of type "image"

func (h HtmlInputElement) Alt() (string, error) {
	return h.getAttributeString("alt")
}

func (h HtmlInputElement) SetAlt(value string) error {
	return h.setAttributeString("alt", value)
}

func (h HtmlInputElement) Height() (string, error) {
	return h.getAttributeString("height")
}

func (h HtmlInputElement) SetHeight(value string) error {
	return h.setAttributeString("height", value)
}

func (h HtmlInputElement) Src() (string, error) {
	return h.getAttributeString("src")
}

func (h HtmlInputElement) SetSrc(value string) error {
	return h.setAttributeString("src", value)
}

func (h HtmlInputElement) Width() (string, error) {
	return h.getAttributeString("width")
}

func (h HtmlInputElement) SetWidth(value string) error {
	return h.setAttributeString("width", value)
}

// Properties that apply only to elements of type "file"

func (h HtmlInputElement) Accept() (string, error) {
	return h.getAttributeString("accept")
}

func (h HtmlInputElement) SetAccept(value string) error {
	return h.setAttributeString("accept", value)
}

func (h HtmlInputElement) Files() (filelist.FileList, error) {
	var files js.Value
	var err error
	if files, err = h.JSObject().GetWithErr("files"); err == nil {
		return filelist.NewFromJSObject(files)
	}
	return filelist.FileList{}, err
}

// Properties that apply only to text/number-containing or elements

func (h HtmlInputElement) Autocomplete() (string, error) {
	return h.getAttributeString("autocomplete")
}

func (h HtmlInputElement) SetAutocomplete(value string) error {
	return h.setAttributeString("autocomplete", value)
}

func (h HtmlInputElement) Max() (string, error) {
	return h.getAttributeString("max")
}

func (h HtmlInputElement) SetMax(value string) error {
	return h.setAttributeString("max", value)
}

func (h HtmlInputElement) MaxLength() (int, error) {
	return h.getAttributeInt("maxLength")
}

func (h HtmlInputElement) SetMaxLength(value int) error {
	return h.setAttributeInt("maxLength", value)
}

func (h HtmlInputElement) Min() (string, error) {
	return h.getAttributeString("min")
}

func (h HtmlInputElement) SetMin(value string) error {
	return h.setAttributeString("min", value)
}

func (h HtmlInputElement) MinLength() (int, error) {
	return h.getAttributeInt("minLength")
}

func (h HtmlInputElement) SetMinLength(value int) error {
	return h.setAttributeInt("minLength", value)
}

func (h HtmlInputElement) Pattern() (string, error) {
	return h.getAttributeString("pattern")
}

func (h HtmlInputElement) SetPattern(value string) error {
	return h.setAttributeString("pattern", value)
}

func (h HtmlInputElement) Placeholder() (string, error) {
	return h.getAttributeString("placeholder")
}

func (h HtmlInputElement) SetPlaceholder(value string) error {
	return h.setAttributeString("placeholder", value)
}

func (h HtmlInputElement) ReadOnly() (bool, error) {
	return h.getAttributeBool("readOnly")
}

func (h HtmlInputElement) SetReadOnly(value bool) error {
	return h.setAttributeBool("readOnly", value)
}

func (h HtmlInputElement) SelectionStart() (int, error) {
	return h.getAttributeInt("selectionStart")
}

func (h HtmlInputElement) SetSelectionStart(value int) error {
	return h.setAttributeInt("selectionStart", value)
}

func (h HtmlInputElement) SelectionEnd() (int, error) {
	return h.getAttributeInt("selectionEnd")
}

func (h HtmlInputElement) SetSelectionEnd(value int) error {
	return h.setAttributeInt("selectionEnd", value)
}

func (h HtmlInputElement) SelectionDirection() (string, error) {
	return h.getAttributeString("selectionDirection")
}

func (h HtmlInputElement) SetSelectionDirection(value string) error {
	return h.setAttributeString("selectionDirection", value)
}

func (h HtmlInputElement) Size() (int, error) {
	return h.getAttributeInt("size")
}

func (h HtmlInputElement) SetSize(value int) error {
	return h.setAttributeInt("size", value)
}

//  Properties not yet categorized

func (h HtmlInputElement) DefaultValue() (string, error) {
	return h.getAttributeString("defaultValue")
}

func (h HtmlInputElement) SetDefaultValue(value string) error {
	return h.setAttributeString("defaultValue", value)
}

func (h HtmlInputElement) DirName() (string, error) {
	return h.getAttributeString("dirName")
}

func (h HtmlInputElement) SetDirName(value string) error {
	return h.setAttributeString("dirName", value)
}

func (h HtmlInputElement) AccessKey() (string, error) {
	return h.getAttributeString("accessKey")
}

func (h HtmlInputElement) SetAccessKey(value string) error {
	return h.setAttributeString("accessKey", value)
}

func (h HtmlInputElement) List() (htmlelement.HtmlElement, error) {
	var obj js.Value
	var err error
	var elem htmlelement.HtmlElement
	if obj, err = h.JSObject().GetWithErr("list"); err == nil {

		elem, err = htmlelement.NewFromJSObject(obj)
	}
	return elem, err
}

func (h HtmlInputElement) Multiple() (bool, error) {
	return h.getAttributeBool("multiple")
}

func (h HtmlInputElement) SetMultiple(value bool) error {
	return h.setAttributeBool("multiple", value)
}

func (h HtmlInputElement) Labels() (nodelist.NodeList, error) {
	var obj js.Value
	var err error
	var arr nodelist.NodeList
	if obj, err = h.JSObject().GetWithErr("labels"); err == nil {

		arr, err = nodelist.NewFromJSObject(obj)
	}
	return arr, err
}

func (h HtmlInputElement) Step() (string, error) {
	return h.getAttributeString("step")
}

func (h HtmlInputElement) SetStep(value string) error {
	return h.setAttributeString("step", value)
}

func (h HtmlInputElement) ValueAsDate() (date.Date, error) {
	var obj js.Value
	var err error
	var arr date.Date
	if obj, err = h.JSObject().GetWithErr("valueAsDate"); err == nil {

		arr, err = date.NewFromJSObject(obj)
	}
	return arr, err
}

func (h HtmlInputElement) ValueAsNumber() (float64, error) {
	return h.getAttributeDouble("valueAsNumber")
}
