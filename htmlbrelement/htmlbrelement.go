package htmlbrelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlbrelementinterface js.Value

//HtmlBrElement struct
type HtmlBRElement struct {
	htmlelement.HtmlElement
}

type HtmlBRElementFrom interface {
	HtmlBRElement() HtmlBRElement
}

func (h HtmlBRElement) HtmlBRElement() HtmlBRElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlbrelementinterface, err = js.Global().GetWithErr("HTMLBRElement"); err != nil {
			htmlbrelementinterface = js.Null()
		}
		baseobject.Register(htmlbrelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlbrelementinterface
}

func New(d document.Document) (HtmlBRElement, error) {
	var err error

	var h HtmlBRElement
	var e element.Element

	if e, err = d.CreateElement("br"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlBRElement, error) {
	var h HtmlBRElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlBrElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlBRElement, error) {
	var h HtmlBRElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlBrElement
}
