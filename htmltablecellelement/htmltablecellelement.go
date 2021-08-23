package htmltablecellelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmltablecellelementinterface js.Value

//HtmlTableCellElement struct
type HtmlTableCellElement struct {
	htmlelement.HtmlElement
}

type HtmlTableCellElementFrom interface {
	HtmlTableCellElement() HtmlTableCellElement
}

func (h HtmlTableCellElement) HtmlTableCellElement() HtmlTableCellElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmltablecellelementinterface, err = js.Global().GetWithErr("HTMLTableCellElement"); err != nil {
			htmltablecellelementinterface = js.Null()
		}
		baseobject.Register(htmltablecellelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmltablecellelementinterface
}

func New(d document.Document) (HtmlTableCellElement, error) {
	var err error

	var h HtmlTableCellElement
	var e element.Element

	if e, err = d.CreateElement("td"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlTableCellElement, error) {
	var h HtmlTableCellElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLTableCellElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlTableCellElement, error) {
	var h HtmlTableCellElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLTableCellElement
}
