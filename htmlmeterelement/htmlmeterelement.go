package htmlmeterelement

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

var htmlmeterelementinterface js.Value

//HtmlMeterElement struct
type HtmlMeterElement struct {
	htmlelement.HtmlElement
}

type HtmlMeterElementFrom interface {
	HtmlMeterElement() HtmlMeterElement
}

func (h HtmlMeterElement) HtmlMeterElement() HtmlMeterElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlmeterelementinterface, err = js.Global().GetWithErr("HTMLMeterElement"); err != nil {
			htmlmeterelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlmeterelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlmeterelementinterface
}

func New(d document.Document) (HtmlMeterElement, error) {
	var err error

	var h HtmlMeterElement
	var e element.Element

	if e, err = d.CreateElement("meter"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlMeterElement, error) {
	var h HtmlMeterElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHTMLMeterElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlMeterElement, error) {
	var h HtmlMeterElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHTMLMeterElement
}

func (h HtmlMeterElement) High() (float64, error) {
	return h.GetAttributeDouble("high")
}

func (h HtmlMeterElement) SetHigh(value float64) error {
	return h.SetAttributeDouble("high", value)
}

func (h HtmlMeterElement) Low() (float64, error) {
	return h.GetAttributeDouble("low")
}

func (h HtmlMeterElement) SetLow(value float64) error {
	return h.SetAttributeDouble("low", value)
}

func (h HtmlMeterElement) Max() (float64, error) {
	return h.GetAttributeDouble("max")
}

func (h HtmlMeterElement) SetMax(value float64) error {
	return h.SetAttributeDouble("max", value)
}

func (h HtmlMeterElement) Min() (float64, error) {
	return h.GetAttributeDouble("min")
}

func (h HtmlMeterElement) SetMin(value float64) error {
	return h.SetAttributeDouble("min", value)
}

func (h HtmlMeterElement) Optimum() (float64, error) {
	return h.GetAttributeDouble("optimum")
}

func (h HtmlMeterElement) SetOptimum(value float64) error {
	return h.SetAttributeDouble("optimum", value)
}

func (h HtmlMeterElement) Value() (float64, error) {
	return h.GetAttributeDouble("value")
}

func (h HtmlMeterElement) SetValue(value float64) error {
	return h.SetAttributeDouble("value", value)
}

func (h HtmlMeterElement) Labels() (nodelist.NodeList, error) {
	var obj js.Value
	var err error
	var arr nodelist.NodeList
	if obj, err = h.JSObject().GetWithErr("labels"); err == nil {

		arr, err = nodelist.NewFromJSObject(obj)
	}
	return arr, err
}
