package htmlmapelement

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

var htmlmapelementinterface js.Value

// HtmlMapElement struct
type HtmlMapElement struct {
	htmlelement.HtmlElement
}

type HtmlMapElementFrom interface {
	HtmlMapElement_() HtmlMapElement
}

func (h HtmlMapElement) HtmlMapElement_() HtmlMapElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlmapelementinterface, err = baseobject.Get(js.Global(), "HTMLMapElement"); err != nil {
			htmlmapelementinterface = js.Undefined()
		}
		baseobject.Register(htmlmapelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlmapelementinterface
}

func New(d document.Document) (HtmlMapElement, error) {
	var err error

	var h HtmlMapElement
	var e element.Element

	if e, err = d.CreateElement("map"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMapElement, error) {
	var h HtmlMapElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLMapElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMapElement, error) {
	var h HtmlMapElement
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHTMLMapElement
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h HtmlMapElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlMapElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlMapElement) Areas() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = h.Get("areas"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
