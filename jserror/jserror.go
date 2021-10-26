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

func (e JSError) JSError_() JSError {
	return e
}

//GetInterface get the Error interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if errorinterface, err = baseobject.Get(js.Global(), "Error"); err != nil {
			errorinterface = js.Undefined()
		}

		baseobject.Register(errorinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})
	return errorinterface
}

func New(values ...interface{}) (JSError, error) {
	var e JSError
	var objs []interface{}

	if len(values) == 1 {
		switch values[0].(type) {
		case string:
			objs = append(objs, values[0])
		case error:
			objs = append(objs, values[0].(error).Error())
		}
	}

	if ei := GetInterface(); !ei.IsUndefined() {
		e.BaseObject = e.SetObject(ei.New(objs...))
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

func (j JSError) Message() (string, error) {
	return j.GetAttributeString("message")
}

func (j JSError) SetMessage(value string) error {
	return j.SetAttributeString("message", value)
}

func (j JSError) Name() (string, error) {
	return j.GetAttributeString("name")
}
func (j JSError) SetName(value string) error {
	return j.SetAttributeString("name", value)
}

func (j JSError) Stack() (string, error) {
	return j.GetAttributeString("stack")
}
