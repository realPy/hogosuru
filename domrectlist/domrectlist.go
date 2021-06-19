package domrectlist

//

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domrect"
)

var singleton sync.Once

var domrectlistinterface js.Value

//DOMRectLists struct
type DOMRectList struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domrectlistinterface, err = js.Global().GetWithErr("DOMRectList"); err != nil {
			domrectlistinterface = js.Null()
		}
	})

	baseobject.Register(domrectlistinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return domrectlistinterface
}

func NewFromJSObject(obj js.Value) (DOMRectList, error) {
	var d DOMRectList
	var err error
	if dli := GetInterface(); !dli.IsNull() {
		if obj.InstanceOf(dli) {
			d.BaseObject = d.SetObject(obj)

		} else {
			err = ErrNotAnDOMRectList
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMRectList) Item(index int) (domrect.DOMRect, error) {

	return domrect.NewFromJSObject(d.JSObject().Index(index))
}