package htmldatalistelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/htmlcollection"
	"github.com/realPy/hogosuru/base/htmlelement"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmldatalistelementinterface js.Value

// HtmlDataElement struct
type HtmlDataListElement struct {
	htmlelement.HtmlElement
}

type HtmlDataListElementFrom interface {
	HtmlDataListElement_() HtmlDataListElement
}

func (h HtmlDataListElement) HtmlDataListElement_() HtmlDataListElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldatalistelementinterface, err = baseobject.Get(js.Global(), "HTMLDataListElement"); err != nil {
			htmldatalistelementinterface = js.Undefined()
		}
		baseobject.Register(htmldatalistelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
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

	if hci := GetInterface(); !hci.IsUndefined() {
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
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHtmlDataListElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlDataListElement) Options() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("options"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
