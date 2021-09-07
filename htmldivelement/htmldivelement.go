package htmldivelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmldivelementinterface js.Value

//HtmlDetailsElement struct
type HtmlDivElement struct {
	htmlelement.HtmlElement
}

type HtmlDivElementFrom interface {
	HtmlDivElement() HtmlDivElement
}

func (h HtmlDivElement) HtmlDivElement() HtmlDivElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldivelementinterface, err = js.Global().GetWithErr("HTMLDivElement"); err != nil {
			htmldivelementinterface = js.Null()
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

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlDivElement
}
