package abortcontroller

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/abortsignal"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var abortcontrollerinterface js.Value

//GetInterface get the JS interface abortcontroller
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
	var ai js.Value
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return a, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(ai) {
		return a, ErrNotAnAbortController
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func New() (AbortController, error) {
	var a AbortController
	var ai, obj js.Value
	var err error
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj, err = baseobject.New(ai); err != nil {
		return a, err
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func (a AbortController) Signal() (abortsignal.AbortSignal, error) {
	var err error
	var obj js.Value
	var as abortsignal.AbortSignal
	if obj, err = a.Get("signal"); err != nil {
		return as, err
	}
	if obj.IsUndefined() {
		return as, baseobject.ErrNotAnObject
	}
	return abortsignal.NewFromJSObject(obj)
}

func (a AbortController) Abort() error {
	var err error
	_, err = a.Call("abort")
	return err
}
