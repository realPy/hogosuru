package htmloptionscollection

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/htmlcollection"
)

var singleton sync.Once

var htmloptionscollectioninterface js.Value

//HTMLOptionsCollection struct
type HtmlOptionsCollection struct {
	htmlcollection.HtmlCollection
}

type HtmlOptionsCollectionFrom interface {
	HtmlOptionsCollection_() HtmlOptionsCollection
}

func (h HtmlOptionsCollection) HTMLOptionsCollection_() HtmlOptionsCollection {
	return h
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if htmloptionscollectioninterface, err = baseobject.Get(js.Global(), "HTMLOptionsCollection"); err != nil {
			htmloptionscollectioninterface = js.Undefined()
		}
		baseobject.Register(htmloptionscollectioninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return htmloptionscollectioninterface
}

func NewFromJSObject(obj js.Value) (HtmlOptionsCollection, error) {
	var h HtmlOptionsCollection
	var err error
	if fli := GetInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fli) {
				h.BaseObject = h.SetObject(obj)
			} else {
				err = ErrNotAnHTMLOptionsCollection
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func (h HtmlOptionsCollection) length() (int, error) {

	return h.GetAttributeInt("length")

}
