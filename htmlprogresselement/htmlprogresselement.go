package htmlprogresselement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/nodelist"
)

var singleton sync.Once

var htmlprogresselementinterface js.Value

//HtmlProgressElement struct
type HtmlProgressElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlprogresselementinterface, err = js.Global().GetWithErr("HTMLProgressElement"); err != nil {
			htmlprogresselementinterface = js.Null()
		}

	})

	baseobject.Register(htmlprogresselementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlprogresselementinterface
}

func New(d document.Document) (HtmlProgressElement, error) {
	var err error

	var h HtmlProgressElement
	var e element.Element

	if e, err = d.CreateElement("progress"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlProgressElement, error) {
	var h HtmlProgressElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlProgressElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlProgressElement, error) {
	var h HtmlProgressElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlProgressElement
}

func (h HtmlProgressElement) Max() (float64, error) {
	return h.GetAttributeDouble("max")
}

func (h HtmlProgressElement) SetMax(value float64) error {
	return h.SetAttributeDouble("max", value)
}

func (h HtmlProgressElement) Position() (float64, error) {
	return h.GetAttributeDouble("position")
}

func (h HtmlProgressElement) Value() (float64, error) {
	return h.GetAttributeDouble("value")
}

func (h HtmlProgressElement) SetValue(value float64) error {
	return h.SetAttributeDouble("value", value)
}

func (h HtmlProgressElement) Labels() (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = h.JSObject().GetWithErr("labels"); err == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}

	return nlist, err
}
