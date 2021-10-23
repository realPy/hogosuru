package htmlelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/cssstyledeclaration"
	"github.com/realPy/hogosuru/element"
)

var singleton sync.Once

var htmlelementinterface js.Value

//HtmlInputElement struct
type HtmlElement struct {
	element.Element
}

type HtmlElementFrom interface {
	HtmlElement_() HtmlElement
}

func (h HtmlElement) HtmlElement_() HtmlElement {
	return h
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlelementinterface, err = baseobject.Get(js.Global(), "HTMLElement"); err != nil {
			htmlelementinterface = js.Undefined()
		}
		baseobject.Register(htmlelementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlelementinterface
}

func NewFromJSObject(obj js.Value) (HtmlElement, error) {
	var h HtmlElement
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			h.BaseObject = h.SetObject(obj)

		} else {
			err = ErrNotAnHtmlElement
		}

	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromElement(elem element.Element) (HtmlElement, error) {
	var h HtmlElement
	var err error

	if ai := GetInterface(); !ai.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(ai) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlElement) AccessKey() (string, error) {
	return h.GetAttributeString("accessKey")
}

func (h HtmlElement) SetAccessKey(value string) error {
	return h.SetAttributeString("accessKey", value)
}

func (h HtmlElement) AccessKeyLabel() (string, error) {
	return h.GetAttributeString("accessKeyLabel")
}

func (h HtmlElement) SetAccessKeyLabel(value string) error {
	return h.SetAttributeString("accessKeyLabel", value)
}

func (h HtmlElement) ClassName() (string, error) {
	return h.GetAttributeString("className")
}

func (h HtmlElement) SetClassName(classname string) error {
	return h.SetAttributeString("className", classname)
}

func (h HtmlElement) ContentEditable() (string, error) {
	return h.GetAttributeString("contentEditable")
}

func (h HtmlElement) SetContentEditable(value string) error {
	return h.SetAttributeString("contentEditable", value)
}

func (h HtmlElement) IsContentEditable() (bool, error) {
	return h.GetAttributeBool("isContentEditable")
}

func (h HtmlElement) Dataset(name string) (interface{}, error) {
	var err error
	var obj, objv js.Value
	var ret interface{}

	if obj, err = h.Get("dataset"); err == nil {

		if objv, err = baseobject.Get(obj, name); err == nil {
			ret = baseobject.GoValue(objv)
		}

	}

	return ret, err

}

func (h HtmlElement) SetDataset(name string, value interface{}) error {

	var err error
	var obj js.Value

	if obj, err = h.Get("dataset"); err == nil {
		err = baseobject.Set(obj, name, js.ValueOf(value))

	}
	return err
}

func (h HtmlElement) Dir() (string, error) {
	return h.GetAttributeString("dir")
}

func (h HtmlElement) Hidden() (bool, error) {
	return h.GetAttributeBool("hidden")
}

func (h HtmlElement) SetHidden(value bool) error {
	return h.SetAttributeBool("hidden", value)
}

func (h HtmlElement) SetDir(value string) error {
	return h.SetAttributeString("dir", value)
}

func (h HtmlElement) Lang() (string, error) {
	return h.GetAttributeString("lang")
}

func (h HtmlElement) SetLang(value string) error {
	return h.SetAttributeString("lang", value)
}

func (h HtmlElement) OffsetHeight() (int, error) {
	return h.GetAttributeInt("offsetHeight")
}

func (h HtmlElement) SetOffsetHeight(value int) error {
	return h.SetAttributeInt("offsetHeight", value)
}

func (h HtmlElement) OffsetLeft() (int, error) {
	return h.GetAttributeInt("offsetLeft")
}

func (h HtmlElement) SetOffsetLeft(value int) error {
	return h.SetAttributeInt("offsetLeft", value)
}

func (h HtmlElement) OffsetParent() (baseobject.BaseObject, error) {
	var err error
	var obj js.Value
	var ret baseobject.BaseObject

	if obj, err = h.Get("offsetParent"); err == nil {
		if !obj.IsUndefined() {
			ret, err = baseobject.NewFromJSObject(obj)
		} else {
			err = baseobject.ErrNotAnObject
		}

	}
	return ret, err
}

func (h HtmlElement) OffsetTop() (int, error) {
	return h.GetAttributeInt("offsetTop")
}

func (h HtmlElement) SetOffsetTop(value int) error {
	return h.SetAttributeInt("offsetTop", value)
}

func (h HtmlElement) OffsetWidth() (int, error) {
	return h.GetAttributeInt("offsetWidth")
}

func (h HtmlElement) SetOffsetWidth(value int) error {
	return h.SetAttributeInt("offsetWidth", value)
}

func (h HtmlElement) Title() (string, error) {
	return h.GetAttributeString("title")
}

func (h HtmlElement) SetTitle(value string) error {
	return h.SetAttributeString("title", value)
}

func (h HtmlElement) Blur() error {
	_, err := h.Call("blur")
	return err
}

func (h HtmlElement) Click() error {
	_, err := h.Call("click")
	return err
}

func (h HtmlElement) Focus() error {
	_, err := h.Call("focus")
	return err
}

func (h HtmlElement) Style() (cssstyledeclaration.CSSStyleDeclaration, error) {
	var err error
	var obj js.Value
	var ret cssstyledeclaration.CSSStyleDeclaration

	if obj, err = h.Get("style"); err == nil {

		if !obj.IsUndefined() {
			ret, err = cssstyledeclaration.NewFromJSObject(obj)
		} else {
			err = baseobject.ErrNotAnObject
		}

	}
	return ret, err
}
