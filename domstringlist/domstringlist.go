package domstringlist

//

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var domstringlistinterface js.Value

//DOMRectLists struct
type DOMStringList struct {
	baseobject.BaseObject
}

type DOMStringListFrom interface {
	DOMStringList() DOMStringList
}

func (d DOMStringList) DOMStringList() DOMStringList {
	return d
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domstringlistinterface, err = js.Global().GetWithErr("DOMStringList"); err != nil {
			domstringlistinterface = js.Null()
		}
		baseobject.Register(domstringlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domstringlistinterface
}

func NewFromJSObject(obj js.Value) (DOMStringList, error) {
	var d DOMStringList
	var err error
	if dli := GetInterface(); !dli.IsNull() {
		if obj.InstanceOf(dli) {
			d.BaseObject = d.SetObject(obj)

		} else {
			err = ErrNotAnDOMStringList
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMStringList) Item(index int) js.Value {
	var obj js.Value
	obj = d.JSObject().Index(index)
	return obj
}

func (d DOMStringList) Contains(search string) (bool, error) {

	var err error
	var obj js.Value
	var result bool
	if obj, err = d.JSObject().CallWithErr("contains", js.ValueOf(search)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}
