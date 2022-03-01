package htmltablecolelement

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

var htmltablecolelementinterface js.Value

//HtmlTableColElement struct
type HtmlTableColElement struct {
	htmlelement.HtmlElement
}

type HtmlTableColElementFrom interface {
	HtmlTableColElement_() HtmlTableColElement
}

func (h HtmlTableColElement) HtmlTableColElement_() HtmlTableColElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablecolelementinterface, err = baseobject.Get(js.Global(), "HTMLTableColElement"); err != nil {
			htmltablecolelementinterface = js.Undefined()
		}
		baseobject.Register(htmltablecolelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecolelementinterface
}

func New(d document.Document) (HtmlTableColElement, error) {
	var err error

	var h HtmlTableColElement
	var e element.Element

	if e, err = d.CreateElement("col"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableColElement, error) {
	var h HtmlTableColElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableColElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableColElement, error) {
	var h HtmlTableColElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLTableColElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlTableColElement) Span() (int, error) {
	return h.GetAttributeInt("span")
}

func (h HtmlTableColElement) SetSpan(value int) error {
	return h.SetAttributeInt("span", value)
}
