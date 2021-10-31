package htmlheadelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlheadelementinterface js.Value

//HtmlHeadElement struct
type HtmlHeadElement struct {
	htmlelement.HtmlElement
}

type HtmlHeadElementFrom interface {
	HtmlHeadElement_() HtmlHeadElement
}

func (h HtmlHeadElement) HtmlHeadElement_() HtmlHeadElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlheadelementinterface, err = baseobject.Get(js.Global(), "HTMLHeadElement"); err != nil {
			htmlheadelementinterface = js.Undefined()
		}
		baseobject.Register(htmlheadelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlheadelementinterface
}

func New(d document.Document) (HtmlHeadElement, error) {
	var err error

	var h HtmlHeadElement
	var e element.Element

	if e, err = d.CreateElement("head"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHeadElement, error) {
	var h HtmlHeadElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlHeadElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHeadElement, error) {
	var h HtmlHeadElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlHeadElement
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return h, err
}
