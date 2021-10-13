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

type DOMRectListFrom interface {
	DOMRectList_() DOMRectList
}

func (d DOMRectList) DOMRectList_() DOMRectList {
	return d
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domrectlistinterface, err = js.Global().GetWithErr("DOMRectList"); err != nil {
			domrectlistinterface = js.Undefined()
		}
		baseobject.Register(domrectlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domrectlistinterface
}

func NewFromJSObject(obj js.Value) (DOMRectList, error) {
	var d DOMRectList
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
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
