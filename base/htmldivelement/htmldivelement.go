package htmldivelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmldivelementinterface js.Value

// HtmlDetailsElement struct
type HtmlDivElement struct {
	htmlelement.HtmlElement
}

type HtmlDivElementFrom interface {
	HtmlDivElement_() HtmlDivElement
}

func (h HtmlDivElement) HtmlDivElement_() HtmlDivElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldivelementinterface, err = baseobject.Get(js.Global(), "HTMLDivElement"); err != nil {
			htmldivelementinterface = js.Undefined()
		}
		baseobject.Register(htmldivelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldivelementinterface
}

func New(d document.Document) (HtmlDivElement, error) {
	var err error

	var h HtmlDivElement
	var e element.Element

	if e, err = d.CreateElement("div"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDivElement, error) {
	var h HtmlDivElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlDivElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDivElement, error) {
	var h HtmlDivElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlDivElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
