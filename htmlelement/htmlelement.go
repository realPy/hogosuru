package htmlelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/element"
)

var singleton sync.Once

var htmlelementinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HtmlInputElement struct
type HtmlElement struct {
	element.Element
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var htmlelementinstance JSInterface
		var err error
		if htmlelementinstance.objectInterface, err = js.Global().GetWithErr("HTMLElement"); err == nil {
			htmlelementinterface = &htmlelementinstance
		}
	})

	return htmlelementinterface
}

func New() (HtmlElement, error) {

	var h HtmlElement

	if hci := GetJSInterface(); hci != nil {
		h.BaseObject = h.SetObject(hci.objectInterface.New())
		return h, nil
	}
	return h, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (HtmlElement, error) {
	var h HtmlElement

	if hei := GetJSInterface(); hei != nil {
		if obj.InstanceOf(hei.objectInterface) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlElement
}
