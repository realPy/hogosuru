package htmldlistelement

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

var htmldlistinterface js.Value

// HtmlDListElement struct
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
		if htmldlistinterface, err = baseobject.Get(js.Global(), "HTMLDListElement"); err != nil {
			htmldlistinterface = js.Undefined()
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

	if hci := GetInterface(); !hci.IsUndefined() {
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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlDListElement
			}
		}
	} else {
		err = ErrNotImplemented

	}
	return h, err
}
