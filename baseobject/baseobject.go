package baseobject

import (
	"syscall/js"
)

/*
var singleton sync.Once

var objinterface *JSInterface
*/
type ObjectFrom interface {
	JSObject() js.Value
}

/*
//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

type ObjectInterface struct {
	object js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var objinstance JSInterface
		var err error
		if objinstance.objectInterface, err = js.Global().GetWithErr("Object"); err == nil {
			objinterface = &objinstance
		}
	})

	return objinterface
}
*/

/*
func NewObject() (ObjectInterface, error) {
	var objectinstance ObjectInterface
	var err error
	if obji := GetJSInterface(); obji != nil {
		objectinstance.object = obji.objectInterface
		return objectinstance, nil
	}
	return objectinstance, err
}

func (o ObjectInterface) Type(object js.Value) (string, error) {
	var err error
	var pobject, strobject, typeobject js.Value

	if pobject, err = o.baseobject.GetWithErr("prototype"); err == nil {
		if strobject, err = pbaseobject.GetWithErr("toString"); err == nil {
			if typeobject, err = strbaseobject.CallWithErr("call", object); err == nil {
				return typebaseobject.String(), nil
			}
		}
	}
	return "", err
}

func (o ObjectInterface) Values(object js.Value) (js.Value, error) {

	if baseobject.Type() == js.TypeObject {

		if value, err := o.baseobject.CallWithErr("values", object); err == nil {
			return value, nil
		} else {
			return js.Value{}, err
		}

	}

	return js.Value{}, ErrNotAnObject
}

func (o ObjectInterface) Entries(object js.Value) (js.Value, error) {
	if baseobject.Type() == js.TypeObject {
		return o.baseobject.CallWithErr("entries", object)
	}

	return js.Value{}, ErrNotAnObject
}

type GOValue struct {
	value interface{}
}

func (h GOValue) JSObject() js.Value {
	return h.value.(js.Value)
}

func (g GOValue) Get(key string) GOValue {
	if g.IsGOMap() {
		return g.GOMap().value[key]
	}
	return g
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
	case GOMap:
		return fmt.Sprintf("%s", value)
	default:
		return "unknown"
	}

}

func (g GOValue) IsGOMap() bool {
	if _, ok := g.value.(GOMap); ok {
		return true
	}
	return false
}

func (g GOValue) IsInt() bool {
	if _, ok := g.value.(int); ok {
		return true
	}
	return false
}

func (g GOValue) IsObject() bool {
	if _, ok := g.value.(js.Value); ok {
		return true
	}
	return false
}

func (g GOValue) Int() int {
	return g.value.(int)
}

func (g GOValue) Object() js.Value {
	return g.value.(js.Value)
}

func (g GOValue) GOMap() GOMap {
	return g.value.(GOMap)
}

func NewGOValue(object js.Value) GOValue {

	switch baseobject.Type() {
	case js.TypeNumber:
		val := baseobject.Float()
		if val == float64(int(val)) {
			return GOValue{value: baseobject.Int()}
		} else {
			return GOValue{value: baseobject.Float()}
		}
	case js.TypeString:
		return GOValue{value: baseobject.String()}
	case js.TypeBoolean:
		return GOValue{value: baseobject.Bool()}
	case js.TypeObject:
		return GOValue{value: object}
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
*/

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

	return "", ErrNotAnObject
}

/*------------------------------------------------------*/

type BaseObject struct {
	object js.Value
}

func NewFromJSObject(obj js.Value) (BaseObject, error) {
	var o BaseObject

	o.object = obj
	return o, nil

}

func (o BaseObject) SetObject(object js.Value) BaseObject {
	o.object = object
	return o
}

func (o BaseObject) JSObject() js.Value {
	return o.object
}

func (o BaseObject) String() string {
	return String(o.object)
}

func (o BaseObject) ToString() (string, error) {
	var value js.Value
	var err error
	if o.JSObject().Type() == js.TypeObject {
		if value, err = o.JSObject().CallWithErr("toString"); err == nil {
			return value.String(), nil
		} else {
			return "", err
		}

	}

	return "", ErrNotAnObject
}

func (o BaseObject) Value() string {
	return o.object.String()
}

func (o BaseObject) Length() int {
	return o.object.Length()
}

func (o BaseObject) Export(name string) {
	js.Global().Set(name, o.object)
}

func Eval(str string) (js.Value, error) {

	return js.Global().CallWithErr("eval", str)

}

func GoValue(object js.Value) interface{} {

	switch object.Type() {
	case js.TypeNumber:
		val := object.Float()
		if val == float64(int(val)) {
			return object.Int()
		} else {
			return object.Float()
		}
	case js.TypeString:
		return object.String()
	case js.TypeBoolean:
		return object.Bool()
	}
	if obj, err := NewFromJSObject(object); err == nil {
		return obj
	}
	return nil
}