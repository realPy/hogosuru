package htmltextareaelement

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

var htmltextareaelementinterface js.Value

//HtmlTextAreaElement struct
type HtmlTextAreaElement struct {
	htmlelement.HtmlElement
}

type HtmlTextAreaElementFrom interface {
	HtmlTextAreaElement() HtmlTextAreaElement
}

func (h HtmlTextAreaElement) HtmlTextAreaElement() HtmlTextAreaElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltextareaelementinterface, err = js.Global().GetWithErr("HTMLTextAreaElement"); err != nil {
			htmltextareaelementinterface = js.Null()
		}

	})

	baseobject.Register(htmltextareaelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmltextareaelementinterface
}

func New(d document.Document) (HtmlTextAreaElement, error) {
	var err error

	var h HtmlTextAreaElement
	var e element.Element

	if e, err = d.CreateElement("textarea"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTextAreaElement, error) {
	var h HtmlTextAreaElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTextAreaElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTextAreaElement, error) {
	var h HtmlTextAreaElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTextAreaElement
}

func (h HtmlTextAreaElement) AccessKey() (string, error) {
	return h.GetAttributeString("accessKey")
}

func (h HtmlTextAreaElement) SetAccessKey(value string) error {
	return h.SetAttributeString("accessKey", value)
}

func (h HtmlTextAreaElement) Autocapitalize() (string, error) {
	return h.GetAttributeString("autocapitalize")
}

func (h HtmlTextAreaElement) SetAutocapitalize(value string) error {
	return h.SetAttributeString("autocapitalize", value)
}

func (h HtmlTextAreaElement) Autocomplete() (string, error) {
	return h.GetAttributeString("autocomplete")
}

func (h HtmlTextAreaElement) SetAutocomplete(value string) error {
	return h.SetAttributeString("autocomplete", value)
}

func (h HtmlTextAreaElement) Autofocus() (bool, error) {
	return h.GetAttributeBool("autofocus")
}

func (h HtmlTextAreaElement) SetAutofocus(value bool) error {
	return h.SetAttributeBool("autofocus", value)
}

func (h HtmlTextAreaElement) Cols() (int, error) {
	return h.GetAttributeInt("cols")
}

func (h HtmlTextAreaElement) SetCols(value int) error {
	return h.SetAttributeInt("cols", value)
}

func (h HtmlTextAreaElement) DefaultValue() (bool, error) {
	return h.GetAttributeBool("defaultValue")
}

func (h HtmlTextAreaElement) SetDefaultValue(value bool) error {
	return h.SetAttributeBool("defaultValue", value)
}

func (h HtmlTextAreaElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlTextAreaElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlTextAreaElement) Form() (htmlformelement.HtmlFormElement, error) {
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

func (h HtmlTextAreaElement) MaxLength() (int, error) {
	return h.GetAttributeInt("maxLength")
}

func (h HtmlTextAreaElement) SetMaxLength(value int) error {
	return h.SetAttributeInt("maxLength", value)
}

func (h HtmlTextAreaElement) MinLength() (int, error) {
	return h.GetAttributeInt("minLength")
}

func (h HtmlTextAreaElement) SetMinLength(value int) error {
	return h.SetAttributeInt("minLength", value)
}

func (h HtmlTextAreaElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlTextAreaElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlTextAreaElement) Placeholder() (string, error) {
	return h.GetAttributeString("placeholder")
}

func (h HtmlTextAreaElement) SetPlaceholder(value string) error {
	return h.SetAttributeString("placeholder", value)
}

func (h HtmlTextAreaElement) ReadOnly() (bool, error) {
	return h.GetAttributeBool("readOnly")
}

func (h HtmlTextAreaElement) SetReadOnly(value bool) error {
	return h.SetAttributeBool("readOnly", value)
}

func (h HtmlTextAreaElement) Required() (bool, error) {
	return h.GetAttributeBool("required")
}

func (h HtmlTextAreaElement) SetRequired(value bool) error {
	return h.SetAttributeBool("required", value)
}

func (h HtmlTextAreaElement) Rows() (int, error) {
	return h.GetAttributeInt("rows")
}

func (h HtmlTextAreaElement) SetRows(value int) error {
	return h.SetAttributeInt("rows", value)
}

func (h HtmlTextAreaElement) SetSelectionStart(value int) error {
	return h.SetAttributeInt("selectionStart", value)
}

func (h HtmlTextAreaElement) SelectionEnd() (int, error) {
	return h.GetAttributeInt("selectionEnd")
}

func (h HtmlTextAreaElement) SetSelectionEnd(value int) error {
	return h.SetAttributeInt("selectionEnd", value)
}

func (h HtmlTextAreaElement) SelectionDirection() (string, error) {
	return h.GetAttributeString("selectionDirection")
}

func (h HtmlTextAreaElement) SetSelectionDirection(value string) error {
	return h.SetAttributeString("selectionDirection", value)
}

func (h HtmlTextAreaElement) TabIndex() (int, error) {
	return h.GetAttributeInt("tabIndex")
}

func (h HtmlTextAreaElement) SetIndex(value int) error {
	return h.SetAttributeInt("tabIndex", value)
}

func (h HtmlTextAreaElement) TextLength() (int, error) {
	return h.GetAttributeInt("textLength")
}

func (h HtmlTextAreaElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlTextAreaElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlTextAreaElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.JSObject().GetWithErr("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlTextAreaElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlTextAreaElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}

func (h HtmlTextAreaElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlTextAreaElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlTextAreaElement) Wrap() (string, error) {
	return h.GetAttributeString("wrap")
}

func (h HtmlTextAreaElement) SetWrap(value string) error {
	return h.SetAttributeString("wrap", value)
}

func (h HtmlTextAreaElement) Labels() (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = h.JSObject().GetWithErr("labels"); err == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}

	return nlist, err
}

func (h HtmlTextAreaElement) Blur() error {
	_, err := h.JSObject().CallWithErr("blur")
	return err
}

func (h HtmlTextAreaElement) Focus() error {
	_, err := h.JSObject().CallWithErr("focus")
	return err
}

func (h HtmlTextAreaElement) Select() error {
	_, err := h.JSObject().CallWithErr("select")
	return err
}

func (h HtmlTextAreaElement) SetRangeText(replacement string, options ...interface{}) error {

	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(replacement))

	for _, option := range options {
		if objGo, ok := option.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(option))
		}
	}
	_, err = h.JSObject().CallWithErr("setRangeText", arrayJS...)
	return err
}

func (h HtmlTextAreaElement) SetSelectionRange(selectionStart, selectionEnd string, selectionDirection ...string) error {

	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(selectionStart))
	arrayJS = append(arrayJS, js.ValueOf(selectionEnd))

	if len(selectionDirection) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(selectionDirection[0]))
	}
	_, err = h.JSObject().CallWithErr("setSelectionRange", arrayJS...)
	return err
}

func (h HtmlTextAreaElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlTextAreaElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlTextAreaElement) SetCustomValidity(message string) error {

	_, err := h.JSObject().CallWithErr("setCustomValidity", js.ValueOf(message))
	return err
}
