package validitystate

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var validitystateinterface js.Value

//ValidityState struct
type ValidityState struct {
	baseobject.BaseObject
}

type ValidityStateFrom interface {
	ValidityState_() ValidityState
}

func (v ValidityState) ValidityState_() ValidityState {
	return v
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if validitystateinterface, err = baseobject.Get(js.Global(), "ValidityState"); err != nil {
			validitystateinterface = js.Undefined()
		}
		baseobject.Register(validitystateinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return validitystateinterface
}

func NewFromJSObject(obj js.Value) (ValidityState, error) {
	var v ValidityState
	var hei js.Value
	if hei = GetInterface(); hei.IsUndefined() {
		return v, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return v, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(hei) {
		return v, ErrNotAnValidityState
	}
	v.BaseObject = v.SetObject(obj)
	return v, nil
}

func (v ValidityState) BadInput() (bool, error) {
	return v.GetAttributeBool("badInput")
}

func (v ValidityState) CustomError() (bool, error) {
	return v.GetAttributeBool("customError")
}

func (v ValidityState) PatternMismatch() (bool, error) {
	return v.GetAttributeBool("patternMismatch")
}

func (v ValidityState) RangeOverflow() (bool, error) {
	return v.GetAttributeBool("rangeOverflow")
}

func (v ValidityState) RangeUnderflow() (bool, error) {
	return v.GetAttributeBool("rangeUnderflow")
}

func (v ValidityState) StepMismatch() (bool, error) {
	return v.GetAttributeBool("stepMismatch")
}

func (v ValidityState) TooLong() (bool, error) {
	return v.GetAttributeBool("tooLong")
}

func (v ValidityState) TooShort() (bool, error) {
	return v.GetAttributeBool("tooShort")
}

func (v ValidityState) TypeMismatch() (bool, error) {
	return v.GetAttributeBool("typeMismatch")
}

func (v ValidityState) Valid() (bool, error) {
	return v.GetAttributeBool("valid")
}

func (v ValidityState) ValueMissing() (bool, error) {
	return v.GetAttributeBool("valueMissing")
}
