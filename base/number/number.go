package number

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterfaceNumber)
}

var singletonnumber sync.Once

var numberinterface js.Value

// GetInterfaceNumber get the JS interface Array
func GetInterfaceNumber() js.Value {

	singletonnumber.Do(func() {

		var err error
		if numberinterface, err = baseobject.Get(js.Global(), "Number"); err != nil {
			numberinterface = js.Undefined()
		}
	})

	return numberinterface
}

// Number struct
type Number struct {
	baseobject.BaseObject
}

type NumberFrom interface {
	Number_() Number
}

func (n Number) Number_() Number {
	return n
}

func IsInteger(obj js.Value) (bool, error) {
	var err error
	var result bool
	if ni := GetInterfaceNumber(); !ni.IsUndefined() {

		if obj, err := baseobject.Call(ni, "isInteger", obj); err == nil {

			if obj.Type() == js.TypeBoolean {
				result = obj.Bool()
			} else {
				err = baseobject.ErrObjectNotBool
			}
		}

	}
	return result, err
}
