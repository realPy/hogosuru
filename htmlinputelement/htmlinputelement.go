package htmlinputelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/filelist"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlinputelementinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HtmlInputElement struct
type HtmlInputElement struct {
	htmlelement.HtmlElement
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var htmlinputelementinstance JSInterface
		var err error
		if htmlinputelementinstance.objectInterface, err = js.Global().GetWithErr("HTMLInputElement"); err == nil {
			htmlinputelementinterface = &htmlinputelementinstance
		}
	})

	return htmlinputelementinterface
}

func New() (HtmlInputElement, error) {

	var h HtmlInputElement

	if hci := GetJSInterface(); hci != nil {
		h.BaseObject = h.SetObject(hci.objectInterface.New())
		return h, nil
	}
	return h, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (HtmlInputElement, error) {
	var h HtmlInputElement

	if hei := GetJSInterface(); hei != nil {
		if obj.InstanceOf(hei.objectInterface) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlInputElement
}

func (h HtmlInputElement) getStringAttribute(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {

		valueStr = obj.String()
	}

	return valueStr, err

}

func (h HtmlInputElement) Value() (string, error) {
	return h.getStringAttribute("value")
}

func (h HtmlInputElement) Files() (filelist.FileList, error) {
	var files js.Value
	var err error
	if files, err = h.JSObject().GetWithErr("files"); err == nil {
		return filelist.NewFromJSObject(files)
	}
	return filelist.FileList{}, err
}
