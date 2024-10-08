package htmlcollection

// https://developer.mozilla.org/fr/docs/Web/API/HTMLCollection

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var htmlcollectioninterface js.Value

// HTMLCollection struct
type HtmlCollection struct {
	baseobject.BaseObject
}

type HtmlCollectionFrom interface {
	HtmlCollection_() HtmlCollection
}

func (h HtmlCollection) HtmlCollection_() HtmlCollection {
	return h
}

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if htmlcollectioninterface, err = baseobject.Get(js.Global(), "HTMLCollection"); err != nil {
			htmlcollectioninterface = js.Undefined()
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
	if fli := GetInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if obj.InstanceOf(fli) {
				h.BaseObject = h.SetObject(obj)
			} else {
				err = ErrNotAnHTMLCollection
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlCollection) Item(index int) (interface{}, error) {

	var i interface{}
	var err error
	obj := h.JSObject().Index(index)
	if !obj.IsUndefined() && !obj.IsNull() {
		i, err = baseobject.Discover(obj)
	}

	return i, err

}
