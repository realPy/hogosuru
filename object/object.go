package object

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type ObjectInterface struct {
	objectInterface js.Value
}

func NewObjectInterface() (ObjectInterface, error) {
	var objectinstance ObjectInterface
	var err error
	objectinstance.objectInterface, err = js.Global().GetWithErr("Object")

	return objectinstance, err
}

func (o ObjectInterface) Type(object js.Value) (string, error) {
	var err error
	var pobject, strobject, typeobject js.Value

	if pobject, err = o.objectInterface.GetWithErr("prototype"); err == nil {
		if strobject, err = pobject.GetWithErr("toString"); err == nil {
			if typeobject, err = strobject.CallWithErr("call", object); err == nil {
				return typeobject.String(), nil
			}
		}
	}
	return "", err
}

func (o ObjectInterface) Values(object js.Value) (js.Value, error) {

	if object.Type() == js.TypeObject {

		if value, err := o.objectInterface.CallWithErr("values", object); err == nil {
			return value, nil
		} else {
			return js.Value{}, err
		}

	}

	return js.Value{}, fmt.Errorf("The given value must be an object")
}

func (o ObjectInterface) Entries(object js.Value) (js.Value, error) {
	if object.Type() == js.TypeObject {
		return o.objectInterface.CallWithErr("entries", object)
	}

	return js.Value{}, fmt.Errorf("The given value must be an object")
}

type GOValue struct {
	value interface{}
}

func (g GOValue) String() string {

	switch value := g.value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case string:
		return fmt.Sprintf("%s", value)
	case float64:
		return fmt.Sprintf("%f", value)
	case bool:
		return fmt.Sprintf("%t", value)
	default:
		return "unknown"
	}

}

func NewGOValue(object js.Value) GOValue {

	switch object.Type() {
	case js.TypeNumber:
		val := object.Float()
		if val == float64(int(val)) {
			return GOValue{value: object.Int()}
		} else {
			return GOValue{value: object.Float()}
		}
	case js.TypeString:
		return GOValue{value: object.String()}
	case js.TypeBoolean:
		return GOValue{value: object.Bool()}
	}
	return GOValue{}
}

func Pair(keypair js.Value) (GOValue, GOValue) {
	if keypair.Type() == js.TypeObject {
		if keypair.Length() == 2 {

			key := keypair.Index(0)
			value := keypair.Index(1)
			return NewGOValue(key), NewGOValue(value)

		}

	}
	return GOValue{}, GOValue{}
}

func String(object js.Value) string {
	s, _ := StringWithErr(object)
	return s
}

//String return the string javascript value represent the object
func StringWithErr(object js.Value) (string, error) {

	if object.Type() == js.TypeObject {
		if value, err := object.CallWithErr("toString"); err == nil {
			return value.String(), nil
		} else {
			return "", err
		}

	}

	return "", fmt.Errorf("The given value must be an object")
}
