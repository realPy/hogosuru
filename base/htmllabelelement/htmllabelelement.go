package htmllabelelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/htmlformelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmllabelelementinterface js.Value

// HtmlLabelElement struct
type HtmlLabelElement struct {
	htmlelement.HtmlElement
}

type HtmlLabelElementFrom interface {
	HtmlLabelElement_() HtmlLabelElement
}

func (h HtmlLabelElement) HtmlLabelElement_() HtmlLabelElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmllabelelementinterface, err = baseobject.Get(js.Global(), "HTMLLabelElement"); err != nil {
			htmllabelelementinterface = js.Undefined()
		}
		baseobject.Register(htmllabelelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllabelelementinterface
}

func New(d document.Document) (HtmlLabelElement, error) {
	var err error

	var h HtmlLabelElement
	var e element.Element

	if e, err = d.CreateElement("label"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLabelElement, error) {
	var h HtmlLabelElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLLabelElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLabelElement, error) {
	var h HtmlLabelElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLLabelElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLabelElement) Control() (htmlelement.HtmlElement, error) {
	var err error
	var obj js.Value
	var htmlelem htmlelement.HtmlElement

	if obj, err = h.Get("control"); err == nil {

		htmlelem, err = htmlelement.NewFromJSObject(obj)
	}

	return htmlelem, err
}

func (h HtmlLabelElement) Form() (htmlformelement.HtmlFormElement, error) {
	var err error
	var obj js.Value
	var formelem htmlformelement.HtmlFormElement

	if obj, err = h.Get("form"); err == nil {

		formelem, err = htmlformelement.NewFromJSObject(obj)
	}

	return formelem, err
}

func (h HtmlLabelElement) HtmlFor() (string, error) {
	return h.GetAttributeString("htmlFor")
}

func (h HtmlLabelElement) SetHtmlFor(value string) error {
	return h.SetAttributeString("htmlFor", value)
}
