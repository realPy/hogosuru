package window

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/location"
)

var singleton sync.Once

var locationinterface js.Value

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if locationinterface, err = js.Global().GetWithErr("Window"); err != nil {
			locationinterface = js.Null()
		}

	})
	baseobject.Register(locationinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})
	return locationinterface
}

type Window struct {
	eventtarget.EventTarget
}

func NewFromJSObject(obj js.Value) (Window, error) {
	var w Window

	if wi := GetInterface(); !wi.IsNull() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func New() (Window, error) {

	var err error
	var w Window
	var windowObj js.Value
	if windowObj, err = js.Global().GetWithErr("window"); err == nil {

		w, err = NewFromJSObject(windowObj)

	}
	return w, err
}

func (w Window) Location() (location.Location, error) {
	var err error
	var obj js.Value
	var l location.Location

	if obj, err = w.JSObject().GetWithErr("location"); err == nil {
		l, err = location.NewFromJSObject(obj)
	}

	return l, err
}
