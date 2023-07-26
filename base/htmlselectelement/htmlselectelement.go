package htmlselectelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlcollection"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/htmlformelement"
	"github.com/realPy/hogosuru/base/htmloptionelement"
	"github.com/realPy/hogosuru/base/htmloptionscollection"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/validitystate"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlselectelementinterface js.Value

// HtmlSelectElement struct
type HtmlSelectElement struct {
	htmlelement.HtmlElement
}

type HtmlSelectElementFrom interface {
	HtmlSelectElement_() HtmlSelectElement
}

func (h HtmlSelectElement) HtmlSelectElement_() HtmlSelectElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlselectelementinterface, err = baseobject.Get(js.Global(), "HTMLSelectElement"); err != nil {
			htmlselectelementinterface = js.Undefined()
		}
		baseobject.Register(htmlselectelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlselectelementinterface
}

func New(d document.Document) (HtmlSelectElement, error) {
	var err error

	var h HtmlSelectElement
	var e element.Element

	if e, err = d.CreateElement("select"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSelectElement, error) {
	var h HtmlSelectElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLSelectElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSelectElement, error) {
	var h HtmlSelectElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLSelectElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlSelectElement) Autofocus() (bool, error) {
	return h.GetAttributeBool("autofocus")
}

func (h HtmlSelectElement) SetAutofocus(value bool) error {
	return h.SetAttributeBool("autofocus", value)
}

func (h HtmlSelectElement) Disabled() (bool, error) {
	return h.GetAttributeBool("disabled")
}

func (h HtmlSelectElement) SetDisabled(value bool) error {
	return h.SetAttributeBool("disabled", value)
}

func (h HtmlSelectElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var f htmlformelement.HtmlFormElement
	if obj, err = h.Get("form"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			f, err = htmlformelement.NewFromJSObject(obj)
		}
	}
	return f, err
}

func (h HtmlSelectElement) Length() (int, error) {
	return h.GetAttributeInt("length")
}

func (h HtmlSelectElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlSelectElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlSelectElement) Options() (htmloptionscollection.HtmlOptionsCollection, error) {

	var err error
	var obj js.Value
	var optioncollection htmloptionscollection.HtmlOptionsCollection

	if obj, err = h.Get("options"); err == nil {

		optioncollection, err = htmloptionscollection.NewFromJSObject(obj)
	}

	return optioncollection, err
}

func (h HtmlSelectElement) Multiple() (bool, error) {
	return h.GetAttributeBool("multiple")
}

func (h HtmlSelectElement) SetMultiple(value bool) error {
	return h.SetAttributeBool("multiple", value)
}

func (h HtmlSelectElement) Required() (bool, error) {
	return h.GetAttributeBool("required")
}

func (h HtmlSelectElement) SetRequired(value bool) error {
	return h.SetAttributeBool("required", value)
}

func (h HtmlSelectElement) SelectedIndex() (int, error) {
	return h.GetAttributeInt("selectedIndex")
}

func (h HtmlSelectElement) SetSelectedIndex(value int) error {
	return h.SetAttributeInt("selectedIndex", value)
}

func (h HtmlSelectElement) SelectedOptions() (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("selectedOptions"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (h HtmlSelectElement) Size() (int, error) {
	return h.GetAttributeInt("size")
}

func (h HtmlSelectElement) SetSize(value int) error {
	return h.SetAttributeInt("size", value)
}

func (h HtmlSelectElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlSelectElement) Validity() (validitystate.ValidityState, error) {
	var err error
	var obj js.Value
	var state validitystate.ValidityState

	if obj, err = h.Get("validity"); err == nil {

		state, err = validitystate.NewFromJSObject(obj)
	}
	return state, err

}

func (h HtmlSelectElement) Value() (string, error) {
	return h.GetAttributeString("value")
}

func (h HtmlSelectElement) SetValue(value string) error {
	return h.SetAttributeString("value", value)
}

func (h HtmlSelectElement) ValidationMessage() (string, error) {
	return h.GetAttributeString("validationMessage")
}

func (h HtmlSelectElement) WillValidate() (bool, error) {
	return h.GetAttributeBool("willValidate")
}

func (h HtmlSelectElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlSelectElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlSelectElement) SetCustomValidity(message string) error {

	_, err := h.Call("setCustomValidity", js.ValueOf(message))
	return err
}

func (h HtmlSelectElement) Add(elem htmloptionelement.HtmlOptionElement, before ...interface{}) error {
	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, elem.JSObject())

	for _, value := range before {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	_, err = h.Call("add", arrayJS...)
	return err
}

func (h HtmlSelectElement) Item(index int) (htmloptionelement.HtmlOptionElement, error) {

	var optelem htmloptionelement.HtmlOptionElement
	var jsobj js.Value
	var err error

	if jsobj, err = h.Call("item", js.ValueOf(index)); err == nil {
		optelem, err = htmloptionelement.NewFromJSObject(jsobj)
	}
	return optelem, err
}

func (h HtmlSelectElement) NamedItem(str string) (htmloptionelement.HtmlOptionElement, error) {

	var optelem htmloptionelement.HtmlOptionElement
	var jsobj js.Value
	var err error

	if jsobj, err = h.Call("namedItem", js.ValueOf(str)); err == nil {
		optelem, err = htmloptionelement.NewFromJSObject(jsobj)
	}
	return optelem, err
}

func (h HtmlSelectElement) Remove(index int) error {
	_, err := h.Call("remove", js.ValueOf(index))
	return err
}
