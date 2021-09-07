package htmlcollection

// https://developer.mozilla.org/fr/docs/Web/API/HTMLCollection

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var htmlcollectioninterface js.Value

//HTMLCollection struct
type HtmlCollection struct {
	baseobject.BaseObject
}

type HtmlCollectionFrom interface {
	HtmlCollection() HtmlCollection
}

func (h HtmlCollection) HtmlCollection() HtmlCollection {
	return h
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if htmlcollectioninterface, err = js.Global().GetWithErr("HTMLCollection"); err != nil {
			htmlcollectioninterface = js.Null()
		}
		baseobject.Register(htmlcollectioninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlcollectioninterface
}

func NewFromJSObject(obj js.Value) (HtmlCollection, error) {
	var h HtmlCollection
	var err error
	if fli := GetInterface(); !fli.IsNull() {
		if obj.InstanceOf(fli) {
			h.BaseObject = h.SetObject(obj)
		}
	} else {
		err = ErrNotAnHTMLCollection
	}

	return h, err
}

func (h HtmlCollection) Item(index int) js.Value {

	return h.JSObject().Index(index)

}
