package htmlpreelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlpreelementinterface js.Value

//HtmlPreElement struct
type HtmlPreElement struct {
	htmlelement.HtmlElement
}

type HtmlPreElementFrom interface {
	HtmlPreElement() HtmlPreElement
}

func (h HtmlPreElement) HtmlPreElement() HtmlPreElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlpreelementinterface, err = js.Global().GetWithErr("HTMLPreElement"); err != nil {
			htmlpreelementinterface = js.Null()
		}

		baseobject.Register(htmlpreelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlpreelementinterface
}

func New(d document.Document) (HtmlPreElement, error) {
	var err error

	var h HtmlPreElement
	var e element.Element

	if e, err = d.CreateElement("p"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlPreElement, error) {
	var h HtmlPreElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLPreElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlPreElement, error) {
	var h HtmlPreElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLPreElement
}