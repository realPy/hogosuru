package validitystate

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var validitystateinterface js.Value

//HtmlInputElement struct
type ValidityState struct {
	baseobject.BaseObject
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if validitystateinterface, err = js.Global().GetWithErr("ValidityState"); err != nil {
			validitystateinterface = js.Null()
		}

	})

	return validitystateinterface
}

func NewFromJSObject(obj js.Value) (ValidityState, error) {
	var v ValidityState

	if hei := GetInterface(); !hei.IsNull() {
		if obj.InstanceOf(hei) {

			v.BaseObject = v.SetObject(obj)
			return v, nil
		}
	}
	return v, ErrNotAnValidityState
}

func (h ValidityState) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}

func (v ValidityState) BadInput() (bool, error) {
	return v.getAttributeBool("badInput")
}

func (v ValidityState) CustomError() (bool, error) {
	return v.getAttributeBool("customError")
}

func (v ValidityState) PatternMismatch() (bool, error) {
	return v.getAttributeBool("patternMismatch")
}

func (v ValidityState) RangeOverflow() (bool, error) {
	return v.getAttributeBool("rangeOverflow")
}

func (v ValidityState) RangeUnderflow() (bool, error) {
	return v.getAttributeBool("rangeUnderflow")
}

func (v ValidityState) StepMismatch() (bool, error) {
	return v.getAttributeBool("stepMismatch")
}

func (v ValidityState) TooLong() (bool, error) {
	return v.getAttributeBool("tooLong")
}

func (v ValidityState) TooShort() (bool, error) {
	return v.getAttributeBool("tooShort")
}

func (v ValidityState) TypeMismatch() (bool, error) {
	return v.getAttributeBool("typeMismatch")
}

func (v ValidityState) Valid() (bool, error) {
	return v.getAttributeBool("valid")
}

func (v ValidityState) ValueMissing() (bool, error) {
	return v.getAttributeBool("valueMissing")
}
