package htmlhrelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlhrelementinterface js.Value

//HtmlHrElement struct
type HtmlHrElement struct {
	htmlelement.HtmlElement
}

type HtmlHrElementFrom interface {
	HtmlHrElement_() HtmlHrElement
}

func (h HtmlHrElement) HtmlHrElement_() HtmlHrElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlhrelementinterface, err = baseobject.Get(js.Global(), "HTMLHRElement"); err != nil {
			htmlhrelementinterface = js.Undefined()
		}
		baseobject.Register(htmlhrelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlhrelementinterface
}

func New(d document.Document) (HtmlHrElement, error) {
	var err error

	var h HtmlHrElement
	var e element.Element

	if e, err = d.CreateElement("hr"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlHrElement, error) {
	var h HtmlHrElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlHrElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlHrElement, error) {
	var h HtmlHrElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlHrElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}
