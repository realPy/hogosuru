package htmltablecaptionelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmltablecaptionelementinterface js.Value

//HtmlTableCaptionElement struct
type HtmlTableCaptionElement struct {
	htmlelement.HtmlElement
}

type HtmlTableCaptionElementFrom interface {
	HtmlTableCaptionElement_() HtmlTableCaptionElement
}

func (h HtmlTableCaptionElement) HtmlTableCaptionElement_() HtmlTableCaptionElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablecaptionelementinterface, err = js.Global().GetWithErr("HTMLTableCaptionElement"); err != nil {
			htmltablecaptionelementinterface = js.Undefined()
		}

		baseobject.Register(htmltablecaptionelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecaptionelementinterface
}

func New(d document.Document) (HtmlTableCaptionElement, error) {
	var err error

	var h HtmlTableCaptionElement
	var e element.Element

	if e, err = d.CreateElement("caption"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableCaptionElement, error) {
	var h HtmlTableCaptionElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableCaptionElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableCaptionElement, error) {
	var h HtmlTableCaptionElement

	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTableCaptionElement
}
