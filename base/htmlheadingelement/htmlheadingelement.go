package htmlheadingelement

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

var htmlheadingelementinterface js.Value

// HtmlHeadingElement struct
type HtmlHeadingElement struct {
	htmlelement.HtmlElement
}

type HtmlHeadingElementFrom interface {
	HtmlHeadingElement_() HtmlHeadingElement
}

func (h HtmlHeadingElement) HtmlHeadingElement_() HtmlHeadingElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlheadingelementinterface, err = baseobject.Get(js.Global(), "HTMLHeadingElement"); err != nil {
			htmlheadingelementinterface = js.Undefined()
		}
		baseobject.Register(htmlheadingelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlheadingelementinterface
}

func NewH1(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h1"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH2(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h2"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH3(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h3"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH4(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h4"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH5(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h5"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewH6(d document.Document) (HtmlHeadingElement, error) {
	var err error

	var h HtmlHeadingElement
	var e element.Element

	if e, err = d.CreateElement("h6"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHeadingElement, error) {
	var h HtmlHeadingElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlHeadingElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHeadingElement, error) {
	var h HtmlHeadingElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlHeadingElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
