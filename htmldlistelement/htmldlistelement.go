package htmldlistelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmldlistinterface js.Value

//HtmlDListElement struct
type HtmlDListElement struct {
	htmlelement.HtmlElement
}

type HtmlDListElementFrom interface {
	HtmlDListElement_() HtmlDListElement
}

func (h HtmlDListElement) HtmlDivElement_() HtmlDListElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldlistinterface, err = js.Global().GetWithErr("HTMLDListElement"); err != nil {
			htmldlistinterface = js.Null()
		}
		baseobject.Register(htmldlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmldlistinterface
}

func New(d document.Document) (HtmlDListElement, error) {
	var err error

	var h HtmlDListElement
	var e element.Element

	if e, err = d.CreateElement("dl"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDListElement, error) {
	var h HtmlDListElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlDListElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDListElement, error) {
	var h HtmlDListElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlDListElement
}
