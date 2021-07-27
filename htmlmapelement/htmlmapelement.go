package htmlmapelement

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

var htmlmapelementinterface js.Value

//HtmlMapElement struct
type HtmlMapElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlmapelementinterface, err = js.Global().GetWithErr("HTMLMapElement"); err != nil {
			htmlmapelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlmapelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
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

	if hci := GetInterface(); !hci.IsNull() {
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

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLMapElement
}

func (h HtmlMapElement) Name() (string, error) {
	return h.GetAttributeString("name")
}

func (h HtmlMapElement) SetName(value string) error {
	return h.SetAttributeString("name", value)
}

func (h HtmlMapElement) Areas() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = h.JSObject().GetWithErr("areas"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
