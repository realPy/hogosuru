package baseobject

import (
	"syscall/js"
)

var registry map[string]func(js.Value) (interface{}, error)

func Register(inter js.Value, contruct func(js.Value) (interface{}, error)) error {
	var obj js.Value
	var err error
	if registry == nil {
		registry = make(map[string]func(js.Value) (interface{}, error))
	}

	//registry[inter.Get("prototype").Call("toString").String()] = contruct
	if obj, err = inter.GetWithErr("name"); err == nil {
		registry[obj.String()] = contruct
	}
	return err
}

func Discover(obj js.Value) (interface{}, error) {
	var err error
	var bobj interface{}
	var objname js.Value
	var objconstructor js.Value

	if objconstructor, err = obj.GetWithErr("constructor"); err == nil {

		if objname, err = objconstructor.GetWithErr("name"); err == nil {
			if f, ok := registry[objname.String()]; ok {
				var obji interface{}
				var ok bool

				if obji, err = f(obj); err == nil {
					if bobj, ok = obji.(ObjectFrom); !ok {
						err = ErrNotABaseObject
					}
				}

			} else {
				bobj, err = NewFromJSObject(obj)
			}

		} else {
			bobj, err = NewFromJSObject(obj)
		}

	} else {
		bobj, err = NewFromJSObject(obj)
	}

	return bobj, err
}

type ObjectFrom interface {
	JSObject() js.Value
	BaseObject_() BaseObject
}

func (b BaseObject) BaseObject_() BaseObject {
	return b
}

func String(object js.Value) string {
	return object.String()
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
	object *js.Value
}

func NewFromJSObject(obj js.Value) (BaseObject, error) {
	var o BaseObject

	o.object = &obj
	return o, nil

}

func (b BaseObject) Empty() bool {

	return b.object == nil
}

func (b BaseObject) Discover() (interface{}, error) {
	return Discover(b.JSObject())
}

func (b BaseObject) SetObject(object js.Value) BaseObject {

	b.object = &object

	return b
}

func (b BaseObject) JSObject() js.Value {
	if b.object != nil {
		return *b.object
	} else {
		return js.Null()
	}

}

func (b BaseObject) String() string {
	return String(*b.object)
}

func (b BaseObject) ToString() (string, error) {
	var value js.Value
	var err error
	if b.JSObject().Type() == js.TypeObject {
		if value, err = b.JSObject().CallWithErr("toString"); err == nil {
			return value.String(), nil
		} else {
			return "", err
		}

	}

	return "", ErrNotAnObject
}

func (b BaseObject) Value() string {
	return b.object.String()
}

func (b BaseObject) Length() int {
	return b.object.Length()
}

func (b BaseObject) Bind(to BaseObject) (interface{}, error) {
	var err error
	var bindObj js.Value
	var gobj interface{}

	if bindObj, err = b.JSObject().CallWithErr("bind", to.JSObject()); err == nil {

		gobj, err = Discover(bindObj)

	}
	return gobj, err
}

func (b BaseObject) Implement(method string) (bool, error) {

	var obj js.Value

	var err error

	if obj, err = b.JSObject().GetWithErr(method); err == nil {

		if obj.Type() == js.TypeFunction {
			return true, nil
		}

	}

	return false, err
}

func (b BaseObject) Class() (string, error) {
	var err error
	var objconstructor, objname js.Value
	var classname string

	if objconstructor, err = b.JSObject().GetWithErr("constructor"); err == nil {

		if objname, err = objconstructor.GetWithErr("name"); err == nil {
			classname = objname.String()
		}

	}
	return classname, err
}

func (b BaseObject) SetFunc(attribute string, f func(this js.Value, args []js.Value) interface{}) error {
	return b.JSObject().SetWithErr(attribute, js.FuncOf(f))
}

func (b BaseObject) Export(name string) {
	js.Global().Set(name, b.object)
}

func (b BaseObject) GetAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {

		if obj.IsNull() {
			err = ErrNotAnObject

		} else {

			//valueStr, err = ToStringWithErr(obj)
			valueStr = obj.String()
		}
	}

	return valueStr, err

}

func (b BaseObject) GetAttributeGlobal(attribute string) (interface{}, error) {

	var err error
	var obj js.Value
	var objGlobal interface{}

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {

		if obj.IsNull() {
			err = ErrNotAnObject

		} else {

			objGlobal, err = Discover(obj)
		}
	}

	return objGlobal, err

}

func (b BaseObject) SetAttributeString(attribute string, value string) error {

	return b.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}

	return ret, err
}

func (b BaseObject) SetAttributeBool(attribute string, value bool) error {

	return b.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Int()
		} else {
			err = ErrObjectNotBool
		}
	}

	return result, err
}
func (b BaseObject) SetAttributeInt(attribute string, value int) error {

	return b.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeDouble(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var result float64

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Float()
		} else {
			err = ErrObjectNotNumber
		}
	}

	return result, err
}

func (b BaseObject) SetAttributeDouble(attribute string, value float64) error {

	return b.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (b BaseObject) CallInt64(method string) (int64, error) {

	var err error
	var obj js.Value
	var ret int64

	if obj, err = b.JSObject().CallWithErr(method); err == nil {
		if obj.Type() == js.TypeNumber {
			ret = int64(obj.Float())
		} else {
			err = ErrObjectNotNumber
		}
	}
	return ret, err
}

func (b BaseObject) CallBool(method string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = b.JSObject().CallWithErr(method); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = ErrObjectNotBool
		}
	}
	return result, err
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

	obj, _ := Discover(object)

	return obj
}
