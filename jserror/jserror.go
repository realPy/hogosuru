package jserror

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var errorinterface js.Value

//JSError JSError struct
type JSError struct {
	baseobject.BaseObject
}

type JSErrorFrom interface {
	JSError_() JSError
}

func (e JSError) DomException_() JSError {
	return e
}

//GetJSInterface get the Error interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if errorinterface, err = js.Global().GetWithErr("Error"); err != nil {
			errorinterface = js.Undefined()
		}

		baseobject.Register(errorinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})
	return errorinterface
}

func New(value interface{}) (JSError, error) {
	var e JSError
	var obj interface{}

	switch value.(type) {
	case string:
		obj = value
	case error:
		obj = value.(error).Error()
	}

	if ei := GetInterface(); !ei.IsUndefined() {
		e.BaseObject = e.SetObject(ei.New(obj))
		return e, nil
	}
	return e, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (JSError, error) {
	var e JSError
	var err error
	if ei := GetInterface(); !ei.IsUndefined() {
		if obj.InstanceOf(ei) {
			e.BaseObject = e.SetObject(obj)
		} else {
			err = ErrNotAnError
		}
	} else {
		err = ErrNotImplemented
	}

	return e, err
}
