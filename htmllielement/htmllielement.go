package htmllielement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmllielementinterface js.Value

//HtmlLIElement struct
type HtmlLIElement struct {
	htmlelement.HtmlElement
}

type HtmlLIElementFrom interface {
	HtmlLIElement_() HtmlLIElement
}

func (h HtmlLIElement) HtmlLIElement_() HtmlLIElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmllielementinterface, err = baseobject.Get(js.Global(), "HTMLLIElement"); err != nil {
			htmllielementinterface = js.Undefined()
		}
		baseobject.Register(htmllielementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmllielementinterface
}

func New(d document.Document) (HtmlLIElement, error) {
	var err error

	var h HtmlLIElement
	var e element.Element

	if e, err = d.CreateElement("li"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlLIElement, error) {
	var h HtmlLIElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLLIElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlLIElement, error) {
	var h HtmlLIElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLLIElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlLIElement) Value() (int, error) {
	return h.GetAttributeInt("value")
}

func (h HtmlLIElement) SetAccessKey(value int) error {
	return h.SetAttributeInt("value", value)
}
