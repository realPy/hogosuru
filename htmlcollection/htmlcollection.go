package htmlcollection

// https://developer.mozilla.org/fr/docs/Web/API/FileList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var htmlcollectioninterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HTMLCollection struct
type HTMLCollection struct {
	object.Object
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

func NewFromJSObject(obj js.Value) HTMLCollection {
	var h HTMLCollection

	if fli := GetJSInterface(); fli != nil {
		if obj.InstanceOf(fli.objectInterface) {
			h.Object = h.SetObject(obj)
		}
	} else {
		h.Error = &ErrNotAnHTMLCollection
	}

	return h
}

func (h HTMLCollection) Item(index int) element.Element {
	var elem element.Element
	if h.Error == nil {
		elem = element.NewFromJSObject(h.JSObject().Index(index))
	}
	return elem
}
