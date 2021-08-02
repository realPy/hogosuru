package htmlspanelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlspanelementinterface js.Value

//HtmlSpanElement struct
type HtmlSpanElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlspanelementinterface, err = js.Global().GetWithErr("HTMLSpanElement"); err != nil {
			htmlspanelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlspanelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlspanelementinterface
}

func New(d document.Document) (HtmlSpanElement, error) {
	var err error

	var h HtmlSpanElement
	var e element.Element

	if e, err = d.CreateElement("span"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlSpanElement, error) {
	var h HtmlSpanElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLSpanElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlSpanElement, error) {
	var h HtmlSpanElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLSpanElement
}
