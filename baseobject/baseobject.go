package baseobject

import (
	"errors"
	"syscall/js"
)

var setFunc js.Value
var getFunc js.Value
var callFunc js.Value
var invokeFunc js.Value
var newFunc js.Value
var errorInterface js.Value

func SetSyscall() {
	//Set Set and get function
	eval_(`hSet = (obj, set , value) => { try { Reflect.set(obj,set,value) ; return }catch(err){ return err } }`)
	eval_(`hGet = (obj, get ) => { try { return [true,Reflect.get(obj,get)] }catch(err){ return [false,err] } }`)
	eval_(`hCall = (obj,method,args) => { try { func=Reflect.get(obj,method); return [true,Reflect.apply(func,obj,args)] } catch (err) { return [false,err] } }`)
	eval_(`hInvoke = (func,args) => { try { return [true,Reflect.apply(func,undefined,args)] } catch (err) { return [false,err] } }`)
	eval_(`hNew= (func,args) => { try { return [true,Reflect.construct(func,args)] } catch (err) { return [false,err] } }`)
	setFunc = js.Global().Get("hSet")
	getFunc = js.Global().Get("hGet")
	callFunc = js.Global().Get("hCall")
	invokeFunc = js.Global().Get("hInvoke")
	newFunc = js.Global().Get("hNew")
	errorInterface = js.Global().Get("Error")
}

func Set(obj js.Value, name string, val interface{}) error {
	ret := setFunc.Invoke(obj, js.ValueOf(name), val)
	if !ret.IsUndefined() {
		return errors.New(ret.Get("message").String())
	}
	return nil
}

func Get(obj js.Value, i interface{}) (js.Value, error) {
	var invokvar interface{}
	if s, ok := i.(string); ok {
		invokvar = js.ValueOf(s)
	} else {
		invokvar = i
	}
	ret := getFunc.Invoke(obj, invokvar)
	if ret.Index(0).Bool() {
		return ret.Index(1), nil
	}
	return ret.Index(1), errors.New(ret.Index(1).Get("message").String())
}

func GetIndex(obj js.Value, index int) (js.Value, error) {
	ret := getFunc.Invoke(obj, js.ValueOf(index))
	if ret.Index(0).Bool() {
		return ret.Index(1), nil
	}
	return ret.Index(1), errors.New(ret.Index(1).Get("message").String())
}

func New(obj js.Value, args ...interface{}) (js.Value, error) {
	var jsargs []interface{}
	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := newFunc.Invoke(obj, jsargs)
	if ret.Index(0).Bool() {
		return ret.Index(1), nil
	}
	return ret.Index(1), errors.New(ret.Index(1).Get("message").String())
}

func Call(obj js.Value, name string, args ...interface{}) (js.Value, error) {
	var jsargs []interface{}
	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := callFunc.Invoke(obj, js.ValueOf(name), jsargs)
	if ret.Index(0).Bool() {
		return ret.Index(1), nil
	}
	return ret.Index(1), errors.New(ret.Index(1).Get("message").String())
}

func Invoke(f js.Value, args ...interface{}) (js.Value, error) {
	var jsargs []interface{}
	for _, arg := range args {
		jsargs = append(jsargs, js.ValueOf(arg))
	}
	ret := invokeFunc.Invoke(f, jsargs)
	if ret.Index(0).Bool() {
		return ret.Index(1), nil
	}
	return ret.Index(1), errors.New(ret.Index(1).Get("message").String())
}

func CopyBytesToGo(dst []byte, src js.Value) (int, error) {
	return js.CopyBytesToGo(dst, src), nil
}
func CopyBytesToJS(dst js.Value, src []byte) (int, error) {
	return js.CopyBytesToJS(dst, src), nil
}

var registry map[string]func(js.Value) (interface{}, error)

//Register Register a construct func for type Object given
func Register(inter js.Value, contruct func(js.Value) (interface{}, error)) error {
	var err error
	var constructname string
	if registry == nil {
		registry = make(map[string]func(js.Value) (interface{}, error))
	}
	//registry[inter.Get("prototype").Call("toString").String()] = contruct
	if constructname, err = GetFuncName(inter); err != nil {
		return err
	}
	registry[constructname] = contruct
	return nil
}

func GetFuncName(inter js.Value) (string, error) {
	var obj js.Value
	var err error
	if obj, err = Get(inter, "name"); err == nil && !obj.IsUndefined() {
		return obj.String(), nil
	}
	return "", ErrUnableGetFunctName
}

//Discover Discover the Object Given and return a Hogosuru Class if the construct is registered
func Discover(obj js.Value) (interface{}, error) {
	var err error
	var bobj interface{}
	var objname js.Value
	var objconstructor js.Value
	if objconstructor, err = Get(obj, "constructor"); err != nil {
		return NewFromJSObject(obj)
	}
	if objname, err = Get(objconstructor, "name"); err != nil {
		return NewFromJSObject(obj)
	}
	if f, ok := registry[objname.String()]; ok {
		var obji interface{}
		var ok bool
		if obji, err = f(obj); err != nil {
			return NewFromJSObject(obj)
		}
		if bobj, ok = obji.(ObjectFrom); !ok {
			return bobj, ErrNotABaseObject
		}
		return bobj, nil
	}
	return NewFromJSObject(obj)
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
	var err error
	var value js.Value
	if object.Type() != js.TypeObject {
		return "", ErrNotAnObject
	}
	if value, err = Call(object, "toString"); err != nil {
		return "", err
	}
	return value.String(), nil
}

/*------------------------------------------------------*/
//BaseObject Base Object where all hogosuru herited from
type BaseObject struct {
	object *js.Value
}

//NewFromJSObject Build a BaseObject from a Js Value Object given
func NewFromJSObject(obj js.Value) (BaseObject, error) {
	var o BaseObject
	if obj.IsUndefined() {
		return o, ErrUndefinedValue
	}
	o.object = &obj
	return o, nil
}

//Empty check if the struct is an empty Struct or have a JS Value attached
func (b BaseObject) Empty() bool {
	return b.object == nil
}

//Get Get Value of object and handle err
func (b BaseObject) Get(name string) (js.Value, error) {
	return Get(b.JSObject(), name)
}

//Get Get Value of object and handle err
func (b BaseObject) GetIndex(index int) (js.Value, error) {
	return GetIndex(b.JSObject(), index)
}

//Set Set Value of object and handle err
func (b BaseObject) Set(name string, value interface{}) error {
	return Set(b.JSObject(), name, value)
}

//Call
func (b BaseObject) Call(name string, args ...interface{}) (js.Value, error) {
	return Call(b.JSObject(), name, args...)
}

//Discover Use Discover of this struct
func (b BaseObject) Discover() (interface{}, error) {
	return Discover(b.JSObject())
}

//ConstructName Get the construct name
func (b BaseObject) ConstructName() (string, error) {
	var construct js.Value
	var err error
	if construct, err = b.Get("constructor"); err != nil {
		return "", nil
	}
	if !construct.IsUndefined() {
		return GetFuncName(construct)
	}
	return "", ErrUnableGetConstruct
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
	}
	return js.Undefined()
}

//String Get the current string representation of the js Value attached to this struct
func (b BaseObject) String() string {
	return String(*b.object)
}

//ToString Get the current toString representation of the js Value attached to this struct
func (b BaseObject) ToString() (string, error) {
	var value js.Value
	var err error
	if b.JSObject().Type() != js.TypeObject {
		return "", ErrNotAnObject
	}
	if value, err = b.Call("toString"); err == nil {
		return value.String(), nil
	}
	return "", err

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
	if bindObj, err = b.Call("bind", to.JSObject()); err == nil {
		return Discover(bindObj)
	}
	return gobj, err
}

//Implement Check if the stuct implement a given name method
func (b BaseObject) Implement(method string) (bool, error) {
	var obj js.Value
	var err error
	if obj, err = b.Get(method); err != nil {
		return false, err
	}
	if obj.Type() == js.TypeFunction {
		return true, nil
	}
	return false, nil
}

func (b BaseObject) Class() (string, error) {
	var err error
	var objconstructor, objname js.Value
	if objconstructor, err = b.Get("constructor"); err != nil {
		return "", err
	}
	if objname, err = Get(objconstructor, "name"); err != nil {
		return "", err
	}
	return objname.String(), nil
}

func (b BaseObject) SetFunc(attribute string, f func(this js.Value, args []js.Value) interface{}) error {
	return b.Set(attribute, js.FuncOf(f))
}

func (b BaseObject) SetAttribute(attribute string, i interface{}) error {
	return b.Set(attribute, GetJsValueOf(i))
}

func (b BaseObject) Export(name string) {
	js.Global().Set(name, b.object)
}

func (b BaseObject) GetAttributeString(attribute string) (string, error) {
	var err error
	var obj js.Value
	if obj, err = b.Get(attribute); err != nil {
		return "", err
	}
	if obj.IsUndefined() {
		return "", ErrUndefinedValue
	}
	if obj.IsNull() {
		return "", ErrUndefinedValue
	}
	if obj.Type() == js.TypeString {
		return obj.String(), nil
	}
	return "", ErrObjectNotString
}

func (b BaseObject) GetAttributeGlobal(attribute string) (interface{}, error) {
	var err error
	var obj js.Value
	var objGlobal interface{}
	if obj, err = b.Get(attribute); err != nil {
		return objGlobal, err
	}
	if obj.IsUndefined() {
		return objGlobal, ErrUndefinedValue
	}
	return GoValue(obj)
}

func (b BaseObject) SetAttributeString(attribute string, value string) error {
	return b.Set(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeBool(attribute string) (bool, error) {
	var err error
	var obj js.Value
	if obj, err = b.Get(attribute); err != nil {
		return false, err
	}
	if obj.Type() == js.TypeBoolean {
		return obj.Bool(), nil
	}
	return false, ErrObjectNotBool
}

func (b BaseObject) SetAttributeBool(attribute string, value bool) error {
	return b.Set(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeInt(attribute string) (int, error) {
	var err error
	var obj js.Value
	var ret int
	if obj, err = b.Get(attribute); err != nil {
		return ret, err
	}
	if obj.IsUndefined() {
		return ret, ErrUndefinedValue
	}
	if obj.Type() == js.TypeNumber {
		return obj.Int(), nil
	}
	return ret, ErrObjectNotNumber
}

func (b BaseObject) GetAttributeInt64(attribute string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if obj, err = b.Get(attribute); err != nil {
		return ret, err
	}
	if obj.IsUndefined() {
		return ret, ErrUndefinedValue
	}
	if obj.Type() == js.TypeNumber {
		return int64(obj.Float()), nil
	}
	return ret, ErrObjectNotNumber
}

func (b BaseObject) SetAttributeInt(attribute string, value int) error {
	return b.Set(attribute, js.ValueOf(value))
}

func (b BaseObject) GetAttributeDouble(attribute string) (float64, error) {
	var err error
	var obj js.Value
	var result float64
	if obj, err = b.Get(attribute); err != nil {
		return result, err
	}
	if obj.IsUndefined() {
		return result, ErrUndefinedValue
	}
	if obj.Type() == js.TypeNumber {
		return obj.Float(), nil
	}
	return result, ErrObjectNotDouble
}

func (b BaseObject) SetAttributeDouble(attribute string, value float64) error {
	return b.Set(attribute, js.ValueOf(value))
}

//CallInt64 Call method given and return a 64bit int
func (b BaseObject) CallInt64(method string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if obj, err = b.Call(method); err != nil {
		return ret, err
	}
	if obj.Type() == js.TypeNumber {
		return int64(obj.Float()), nil
	}
	return ret, ErrObjectNotNumber
}

//CallInt Call method given and return int
func (b BaseObject) CallInt(method string) (int, error) {
	var err error
	var obj js.Value
	var ret int
	if obj, err = b.Call(method); err != nil {
		return ret, err
	}
	if obj.Type() == js.TypeNumber {
		return obj.Int(), nil
	}
	return ret, ErrObjectNotNumber
}

//CallInt64 Call method given and return a bool int
func (b BaseObject) CallBool(method string) (bool, error) {
	var err error
	var obj js.Value
	if obj, err = b.Call(method); err != nil {
		return false, err
	}
	if obj.Type() == js.TypeBoolean {
		return obj.Bool(), nil
	}
	return false, ErrObjectNotBool
}

func eval_(str string) {
	js.Global().Call("eval", str)
}

func Eval(str string) (js.Value, error) {
	f := js.Global().Get("eval")
	return Invoke(f, str)
}

func GoValue(object js.Value) (interface{}, error) {
	var err error
	switch object.Type() {
	case js.TypeNumber:
		val := object.Float()
		if val == float64(int(val)) {
			return object.Int(), nil
		}
		return object.Float(), nil
	case js.TypeString:
		return object.String(), nil
	case js.TypeBoolean:
		return object.Bool(), nil
	case js.TypeNull:
		return nil, nil
	}
	obj, err := Discover(object)
	return obj, err
}

func GetJsValueOf(i interface{}) js.Value {
	if objGo, ok := i.(ObjectFrom); ok {
		return objGo.JSObject()
	}
	return js.ValueOf(i)
}
