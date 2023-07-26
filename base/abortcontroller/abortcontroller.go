package abortcontroller

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/abortsignal"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var abortcontrollerinterface js.Value

// GetInterface get the JS interface abortcontroller
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

// AbortController struct
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
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.BaseObject = a.SetObject(obj)

			} else {
				err = ErrNotAnAbortController
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func New() (AbortController, error) {

	var a AbortController
	var obj js.Value
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = baseobject.New(ai); err == nil {
			a.BaseObject = a.SetObject(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return a, err
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
