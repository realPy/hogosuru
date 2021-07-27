package htmlparagraphelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlparagraphelementinterface js.Value

//HtmlParagraphElement struct
type HtmlParagraphElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlparagraphelementinterface, err = js.Global().GetWithErr("HTMLParagraphElement"); err != nil {
			htmlparagraphelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlparagraphelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlparagraphelementinterface
}

func New(d document.Document) (HtmlParagraphElement, error) {
	var err error

	var h HtmlParagraphElement
	var e element.Element

	if e, err = d.CreateElement("p"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLParagraphElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlParagraphElement, error) {
	var h HtmlParagraphElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLParagraphElement
}