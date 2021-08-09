package htmlformelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlformelementinterface js.Value

//HtmlBaseElement struct
type HtmlFormElement struct {
	htmlelement.HtmlElement
}

type HtmlFormElementFrom interface {
	HtmlFormElement() HtmlFormElement
}

func (h HtmlFormElement) HtmlFormElement() HtmlFormElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlformelementinterface, err = js.Global().GetWithErr("HTMLFormElement"); err != nil {
			htmlformelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlformelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlformelementinterface
}

func New(d document.Document) (HtmlFormElement, error) {
	var err error

	var h HtmlFormElement
	var e element.Element

	if e, err = d.CreateElement("form"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlFormElement, error) {
	var h HtmlFormElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlFormElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlFormElement, error) {
	var h HtmlFormElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlFormElement
}

func (h HtmlFormElement) Name() (string, error) {

	return h.GetAttributeString("name")
}

func (h HtmlFormElement) SetName(name string) error {
	return h.SetAttributeString("name", name)
}

func (h HtmlFormElement) Method() (string, error) {

	return h.GetAttributeString("method")
}

func (h HtmlFormElement) SetMethod(value string) error {
	return h.SetAttributeString("method", value)
}

func (h HtmlFormElement) Target() (string, error) {

	return h.GetAttributeString("target")
}

func (h HtmlFormElement) SetTarget(value string) error {
	return h.SetAttributeString("target", value)
}

func (h HtmlFormElement) Action() (string, error) {

	return h.GetAttributeString("action")
}

func (h HtmlFormElement) SetAction(value string) error {
	return h.SetAttributeString("action", value)
}

func (h HtmlFormElement) Encoding() (string, error) {

	return h.GetAttributeString("encoding")
}

func (h HtmlFormElement) SetEncoding(value string) error {
	return h.SetAttributeString("encoding", value)
}

func (h HtmlFormElement) Enctype() (string, error) {
	return h.GetAttributeString("enctype")
}

func (h HtmlFormElement) SetEnctype(value string) error {
	return h.SetAttributeString("enctype", value)
}

func (h HtmlFormElement) AcceptCharset() (string, error) {
	return h.GetAttributeString("acceptCharset")
}

func (h HtmlFormElement) SetAcceptCharset(value string) error {
	return h.SetAttributeString("acceptCharset", value)
}

func (h HtmlFormElement) Autocomplete() (string, error) {
	return h.GetAttributeString("autocomplete")
}

func (h HtmlFormElement) SetAutocomplete(value string) error {
	return h.SetAttributeString("autocomplete", value)
}

func (h HtmlFormElement) NoValidate() (bool, error) {
	return h.GetAttributeBool("noValidate")
}

func (h HtmlFormElement) SetNoValidate(value bool) error {
	return h.SetAttributeBool("noValidate", value)
}

func (h HtmlFormElement) CheckValidity() (bool, error) {

	return h.CallBool("checkValidity")
}

func (h HtmlFormElement) ReportValidity() (bool, error) {

	return h.CallBool("reportValidity")
}

func (h HtmlFormElement) RequestSubmit(elem ...baseobject.BaseObject) error {

	var arrayJS []interface{}

	if len(elem) > 0 {
		arrayJS = append(arrayJS, elem[0].JSObject())
	}

	_, err := h.JSObject().CallWithErr("requestSubmit", arrayJS...)
	return err
}

func (h HtmlFormElement) Reset() error {
	_, err := h.JSObject().CallWithErr("reset")
	return err
}

func (h HtmlFormElement) Submit() error {
	_, err := h.JSObject().CallWithErr("submit")
	return err
}
