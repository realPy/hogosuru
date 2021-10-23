package htmltimeelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmltimeelementinterface js.Value

//HtmlTimeElement struct
type HtmlTimeElement struct {
	htmlelement.HtmlElement
}

type HtmlTimeElementFrom interface {
	HtmlTimeElement_() HtmlTimeElement
}

func (h HtmlTimeElement) HtmlTimeElement_() HtmlTimeElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltimeelementinterface, err = js.Global().GetWithErr("HTMLTimeElement"); err != nil {
			htmltimeelementinterface = js.Undefined()
		}
		baseobject.Register(htmltimeelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltimeelementinterface
}

func New(d document.Document) (HtmlTimeElement, error) {
	var err error

	var h HtmlTimeElement
	var e element.Element

	if e, err = d.CreateElement("time"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTimeElement, error) {
	var h HtmlTimeElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTimeElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTimeElement, error) {
	var h HtmlTimeElement

	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTimeElement
}

func (h HtmlTimeElement) DateTime() (string, error) {
	return h.GetAttributeString("dateTime")
}

func (h HtmlTimeElement) SetDateTime(value string) error {
	return h.SetAttributeString("dateTime", value)
}
