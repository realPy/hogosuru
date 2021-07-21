package baseobject

import (
	"syscall/js"
)

var registry map[string]func(js.Value) (interface{}, error)

func Register(inter js.Value, contruct func(js.Value) (interface{}, error)) {
	if registry == nil {
		registry = make(map[string]func(js.Value) (interface{}, error))
	}

	//registry[inter.Get("prototype").Call("toString").String()] = contruct
	registry[inter.Get("name").String()] = contruct
}

func Discover(obj js.Value) (BaseObject, error) {
	var err error
	var bobj BaseObject

	//if f, ok := registry[obj.Get("constructor").Get("prototype").Call("toString").String()]; ok {
	if f, ok := registry[obj.Get("constructor").Get("name").String()]; ok {
		var obji interface{}
		var ok bool
		obji, err = f(obj)
		if bobj, ok = obji.(BaseObject); !ok {
			err = ErrNotABaseObject
		}

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

func (o BaseObject) GetAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = o.JSObject().GetWithErr(attribute); err == nil {

		if obj.IsNull() {
			err = ErrNotAnObject

		} else {

			//valueStr, err = ToStringWithErr(obj)
			valueStr = obj.String()
		}
	}

	return valueStr, err

}

func (o BaseObject) GetAttributeGlobal(attribute string) (BaseObject, error) {

	var err error
	var obj js.Value
	var objGlobal BaseObject

	if obj, err = o.JSObject().GetWithErr(attribute); err == nil {

		if obj.IsNull() {
			err = ErrNotAnObject

		} else {

			objGlobal, err = Discover(obj)
		}
	}

	return objGlobal, err

}

func (o BaseObject) SetAttributeString(attribute string, value string) error {

	return o.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (o BaseObject) GetAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = o.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return ret, err
}

func (o BaseObject) SetAttributeBool(attribute string, value bool) error {

	return o.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (o BaseObject) GetAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	if obj, err = o.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Int()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err
}
func (o BaseObject) SetAttributeInt(attribute string, value int) error {

	return o.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (o BaseObject) GetAttributeDouble(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var result float64

	if obj, err = o.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Float()
		} else {
			err = ErrObjectNotNumber
		}
	}

	return result, err
}

func (o BaseObject) SetAttributeDouble(attribute string, value float64) error {

	return o.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (o BaseObject) CallInt64(method string) (int64, error) {

	var err error
	var obj js.Value
	var ret int64

	if obj, err = o.JSObject().CallWithErr(method); err == nil {
		if obj.Type() == js.TypeNumber {
			ret = int64(obj.Float())
		} else {
			err = ErrObjectNotNumber
		}
	}
	return ret, err
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
