package htmlimageelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/promise"
)

var singleton sync.Once

var htmlimageelementinterface js.Value

//HtmlImageElement struct
type HtmlImageElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlimageelementinterface, err = js.Global().GetWithErr("HTMLImageElement"); err != nil {
			htmlimageelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlimageelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlimageelementinterface
}

func New(d document.Document) (HtmlImageElement, error) {
	var err error

	var h HtmlImageElement
	var e element.Element

	if e, err = d.CreateElement("img"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlImageElement, error) {
	var h HtmlImageElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmImageElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlImageElement, error) {
	var h HtmlImageElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmImageElement
}

func (h HtmlImageElement) Alt() (string, error) {
	return h.GetAttributeString("alt")
}

func (h HtmlImageElement) SetAlt(value string) error {
	return h.SetAttributeString("alt", value)
}

func (h HtmlImageElement) Complete() (bool, error) {
	return h.GetAttributeBool("complete")
}

func (h HtmlImageElement) CrossOrigin() (string, error) {
	return h.GetAttributeString("crossOrigin")
}

func (h HtmlImageElement) SetCrossOrigin(value string) error {
	return h.SetAttributeString("crossOrigin", value)
}

func (h HtmlImageElement) CurrentSrc() (string, error) {
	return h.GetAttributeString("currentSrc")
}

func (h HtmlImageElement) Decoding() (string, error) {
	return h.GetAttributeString("decoding")
}

func (h HtmlImageElement) SetDecoding(value string) error {
	return h.SetAttributeString("decoding", value)
}

func (h HtmlImageElement) Loading() (string, error) {
	return h.GetAttributeString("loading")
}

func (h HtmlImageElement) SetLoading(value string) error {
	return h.SetAttributeString("loading", value)
}

func (h HtmlImageElement) NaturalHeight() (int, error) {
	return h.GetAttributeInt("naturalHeight")
}

func (h HtmlImageElement) NaturalWidth() (int, error) {
	return h.GetAttributeInt("naturalWidth")
}

func (h HtmlImageElement) Src() (string, error) {
	return h.GetAttributeString("src")
}

func (h HtmlImageElement) SetSrc(value string) error {
	return h.SetAttributeString("src", value)
}

func (h HtmlImageElement) X() (int, error) {
	return h.GetAttributeInt("x")
}

func (h HtmlImageElement) Y() (int, error) {
	return h.GetAttributeInt("y")
}

func (h HtmlImageElement) Decode() (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = h.JSObject().CallWithErr("decode"); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
