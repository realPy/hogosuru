package htmlcollection

// https://developer.mozilla.org/fr/docs/Web/API/HTMLCollection

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
)

var singleton sync.Once

var htmlcollectioninterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HTMLCollection struct
type HTMLCollection struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var htmlcollectioninstance JSInterface
		var err error
		if htmlcollectioninstance.objectInterface, err = js.Global().GetWithErr("HTMLCollection"); err == nil {
			htmlcollectioninterface = &htmlcollectioninstance
		}
	})

	return htmlcollectioninterface
}

func NewFromJSObject(obj js.Value) (HTMLCollection, error) {
	var h HTMLCollection
	var err error
	if fli := GetJSInterface(); fli != nil {
		if obj.InstanceOf(fli.objectInterface) {
			h.BaseObject = h.SetObject(obj)
		}
	} else {
		err = ErrNotAnHTMLCollection
	}

	return h, err
}

func (h HTMLCollection) Item(index int) (element.Element, error) {

	return element.NewFromJSObject(h.JSObject().Index(index))

}