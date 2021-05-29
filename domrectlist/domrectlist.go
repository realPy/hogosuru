package domrectlist

//

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domrect"
)

var singleton sync.Once

var domrectlistinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//DOMRectLists struct
type DOMRectList struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var domrectlistinstance JSInterface
		var err error
		if domrectlistinstance.objectInterface, err = js.Global().GetWithErr("DOMRectList"); err == nil {
			domrectlistinterface = &domrectlistinstance
		}
	})

	return domrectlistinterface
}

func NewFromJSObject(obj js.Value) (DOMRectList, error) {
	var d DOMRectList
	var err error
	if dli := GetJSInterface(); dli != nil {
		if obj.InstanceOf(dli.objectInterface) {
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
