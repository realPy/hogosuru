package abortsignal

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
)

var singleton sync.Once

var abortsignalinterface js.Value

//GetInterface get the JS interface abortsignal
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if abortsignalinterface, err = baseobject.Get(js.Global(), "AbortSignal"); err != nil {
			abortsignalinterface = js.Undefined()
		}
		baseobject.Register(abortsignalinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return abortsignalinterface
}

//AbortSignal struct
type AbortSignal struct {
	eventtarget.EventTarget
}

type AbortSignalFrom interface {
	AbortSignal_() AbortSignal
}

func NewFromJSObject(obj js.Value) (AbortSignal, error) {
	var a AbortSignal
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			a.BaseObject = a.SetObject(obj)
			return a, nil
		} else {
			err = ErrNotAnAbortSignal
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a AbortSignal) Aborted() (bool, error) {

	return a.GetAttributeBool("aborted")

}

func (a AbortSignal) Abort() (AbortSignal, error) {
	var err error
	var obj js.Value
	var as AbortSignal
	if obj, err = a.Call("abort"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			as, err = NewFromJSObject(obj)
		}
	}
	return as, err
}
