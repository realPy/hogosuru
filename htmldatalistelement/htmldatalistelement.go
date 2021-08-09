package htmldatalistelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmldatalistelementinterface js.Value

//HtmlDataElement struct
type HtmlDataListElement struct {
	htmlelement.HtmlElement
}

type HtmlDataListElementFrom interface {
	HtmlDataListElement() HtmlDataListElement
}

func (h HtmlDataListElement) HtmlDataListElement() HtmlDataListElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldatalistelementinterface, err = js.Global().GetWithErr("HTMLDataListElement"); err != nil {
			htmldatalistelementinterface = js.Null()
		}

	})

	baseobject.Register(htmldatalistelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmldatalistelementinterface
}

func New(d document.Document) (HtmlDataListElement, error) {
	var err error

	var h HtmlDataListElement
	var e element.Element

	if e, err = d.CreateElement("datalist"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDataListElement, error) {
	var h HtmlDataListElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlDataListElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDataListElement, error) {
	var h HtmlDataListElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlDataListElement
}

func (h HtmlDataListElement) Options() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.JSObject().CallWithErr("options"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
