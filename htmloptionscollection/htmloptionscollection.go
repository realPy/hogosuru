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
type HTMLOptionsCollection struct {
	htmlcollection.HTMLCollection
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if htmloptionscollectioninterface, err = js.Global().GetWithErr("HTMLOptionsCollection"); err != nil {
			htmloptionscollectioninterface = js.Null()
		}

	})
	baseobject.Register(htmloptionscollectioninterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmloptionscollectioninterface
}

func NewFromJSObject(obj js.Value) (HTMLOptionsCollection, error) {
	var h HTMLOptionsCollection
	var err error
	if fli := GetInterface(); !fli.IsNull() {
		if obj.InstanceOf(fli) {
			h.BaseObject = h.SetObject(obj)
		}
	} else {
		err = ErrNotAnHTMLOptionsCollection
	}

	return h, err
}

func (h HTMLOptionsCollection) length() (int, error) {

	return h.GetAttributeInt("length")

}
