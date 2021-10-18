package baseobject

import (
	"syscall/js"
)

var registry map[string]func(js.Value) (interface{}, error)

//Register Register a construct func for type Object given
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

//Discover Discover the Object Given and return a Hogosuru Class if the construct is registered
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

//ObjectFrom Interface to check if Object is a BaseObject
type ObjectFrom interface {
	JSObject() js.Value
	BaseObject_() BaseObject
}

//BaseObject_ Return the current BaseObject
func (b BaseObject) BaseObject_() BaseObject {
	return b
}

//String return the string representation of the given Object
func String(object js.Value) string {
	return object.String()
}

//ToStringWithErr return the ToString representation of the given Object
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

//BaseObject Base Object where all hogosuru herited from
type BaseObject struct {
	object *js.Value
}

//NewFromJSObject Build a BaseObject from a Js Value Object given
func NewFromJSObject(obj js.Value) (BaseObject, error) {
	var o BaseObject

	o.object = &obj
	return o, nil

}

//Empty check if the struct is an empty Struct or have a JS Value attached
func (b BaseObject) Empty() bool {

	return b.object == nil
}

//Discover Use Discover of this struct
func (b BaseObject) Discover() (interface{}, error) {
	return Discover(b.JSObject())
}

//SetObject Set the JS value Object to this struct
func (b BaseObject) SetObject(object js.Value) BaseObject {

	b.object = &object

	return b
}

//JSObject Give the JS Value Object attach to this struct
func (b BaseObject) JSObject() js.Value {
	if b.object != nil {
		return *b.object
	} else {
		return js.Undefined()
	}

}

//String Get the current string representation of the js Value attached to this struct
func (b BaseObject) String() string {
	return String(*b.object)
}

//ToString Get the current toString representation of the js Value attached to this struct
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

//Value Equivalent to String()
func (b BaseObject) Value() string {
	return b.object.String()
}

//Length Length of the JS.Value attached of this strict
func (b BaseObject) Length() int {
	return b.object.Length()
}

//Bind Bind
func (b BaseObject) Bind(to BaseObject) (interface{}, error) {
	var err error
	var bindObj js.Value
	var gobj interface{}

	if bindObj, err = b.JSObject().CallWithErr("bind", to.JSObject()); err == nil {

		gobj, err = Discover(bindObj)

	}
	return gobj, err
}

//Implement Check if the stuct implement a given name method
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

func (b BaseObject) SetAttribute(attribute string, i interface{}) error {
	var obj interface{}

	if objGo, ok := i.(ObjectFrom); ok {
		obj = objGo

	} else {
		obj = js.ValueOf(i)
	}

	return b.JSObject().SetWithErr(attribute, obj)
}

func (b BaseObject) Export(name string) {
	js.Global().Set(name, b.object)
}

func (b BaseObject) GetAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var ret = ""

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {

		if obj.Type() == js.TypeString {
			ret = obj.String()
		} else {
			err = ErrObjectNotString
		}

	}
	return ret, err

}

func (b BaseObject) GetAttributeGlobal(attribute string) (interface{}, error) {

	var err error
	var obj js.Value
	var objGlobal interface{}

	if obj, err = b.JSObject().GetWithErr(attribute); err == nil {

		if obj.IsUndefined() {
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
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = ErrObjectNotNumber
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
			err = ErrObjectNotDouble
		}
	}

	return result, err
}

func (b BaseObject) SetAttributeDouble(attribute string, value float64) error {

	return b.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

//CallInt64 Call method given and return a 64bit int
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

//CallInt64 Call method given and return a bool int
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
	case js.TypeNull:
		return nil
	}

	obj, _ := Discover(object)

	return obj
}
