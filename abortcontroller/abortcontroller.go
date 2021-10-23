package abortcontroller

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/abortsignal"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var abortcontrollerinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if abortcontrollerinterface, err = baseobject.Get(js.Global(), "AbortController"); err != nil {
			abortcontrollerinterface = js.Undefined()
		}
		baseobject.Register(abortcontrollerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return abortcontrollerinterface
}

//AbortController struct
type AbortController struct {
	baseobject.BaseObject
}

type AbortControllerFrom interface {
	AbortController_() AbortController
}

func (a AbortController) AbortController_() AbortController {
	return a
}

func NewFromJSObject(obj js.Value) (AbortController, error) {
	var a AbortController
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			a.BaseObject = a.SetObject(obj)
			return a, nil
		} else {
			err = ErrNotAnAbortController
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func New() (AbortController, error) {

	var a AbortController

	if ai := GetInterface(); !ai.IsUndefined() {

		a.BaseObject = a.SetObject(ai.New())
		return a, nil
	}
	return a, ErrNotImplemented
}

func (a AbortController) Signal() (abortsignal.AbortSignal, error) {
	var err error
	var obj js.Value
	var as abortsignal.AbortSignal
	if obj, err = a.Get("signal"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			as, err = abortsignal.NewFromJSObject(obj)
		}
	}
	return as, err
}

func (a AbortController) Abort() error {
	var err error
	_, err = a.Call("abort")
	return err
}
