package htmlbodyelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlbodyelementinterface js.Value

//HtmlBodyElement struct
type HtmlBodyElement struct {
	htmlelement.HtmlElement
}

type HtmlBodyElementFrom interface {
	HtmlBodyElement_() HtmlBodyElement
}

func (h HtmlBodyElement) HtmlBodyElement_() HtmlBodyElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlbodyelementinterface, err = js.Global().GetWithErr("HTMLBodyElement"); err != nil {
			htmlbodyelementinterface = js.Null()
		}
		baseobject.Register(htmlbodyelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlbodyelementinterface
}

func New(d document.Document) (HtmlBodyElement, error) {
	var err error

	var h HtmlBodyElement
	var e element.Element

	if e, err = d.CreateElement("body"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlBodyElement, error) {
	var h HtmlBodyElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlBodyElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlBodyElement, error) {
	var h HtmlBodyElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlBodyElement
}
