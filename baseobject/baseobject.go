package baseobject

import (
	"syscall/js"
)

var registry map[string]func(js.Value) (interface{}, error)

func Register(inter js.Value, contruct func(js.Value) (interface{}, error)) {
	if registry == nil {
		registry = make(map[string]func(js.Value) (interface{}, error))
	}

	registry[inter.Get("prototype").Call("toString").String()] = contruct
}

func Discover(obj js.Value) (interface{}, error) {
	var err error
	var bobj interface{}

	if f, ok := registry[obj.Get("constructor").Get("prototype").Call("toString").String()]; ok {

		bobj, err = f(obj)
	} else {
		bobj, err = NewFromJSObject(obj)
	}

	return bobj, err
}

type ObjectFrom interface {
	JSObject() js.Value
}

func String(object js.Value) string {
	return object.String()

	/*	s, _ := StringWithErr(object)
		return s*/
}

//String return the string javascript value represent the object
func ToStringWithErr(object js.Value) (string, error) {

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
