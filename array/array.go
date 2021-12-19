package array

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

var singleton sync.Once

var arrayinterface js.Value

//GetInterface get the JS interface Array
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if arrayinterface, err = baseobject.Get(js.Global(), "Array"); err != nil {
			arrayinterface = js.Undefined()
		}
		baseobject.Register(arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return arrayinterface
}

//Array struct
type Array struct {
	baseobject.BaseObject
}

type ArrayFrom interface {
	Array_() Array
}

func (a Array) Array_() Array {
	return a
}

func NewEmpty(size int) (Array, error) {
	var a Array
	var ai, obj js.Value
	var err error
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj, err = baseobject.New(ai, js.ValueOf(size)); err != nil {
		return a, err
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func From(iterable interface{}, f ...func(interface{}) interface{}) (Array, error) {
	var a Array
	var err error
	var ai, obj js.Value
	var jsfunc js.Func
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	var opts []interface{} = []interface{}{baseobject.GetJsValueOf(iterable)}
	if f != nil && len(f) == 1 {
		jsfunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			b := f[0](baseobject.GoValue_(args[0]))
			return js.ValueOf(b)
		})
		opts = append(opts, jsfunc)
	}
	if obj, err = baseobject.Call(ai, "from", opts...); err != nil {
		return a, err
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func Of(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}
	var ai js.Value
	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	a.BaseObject = a.SetObject(ai.Call("of", arrayJS...))
	return a, nil
}

func New(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}
	var ai, obj js.Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj, err = baseobject.New(ai, arrayJS...); err != nil {
		return a, err
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func NewFromJSObject(obj js.Value) (Array, error) {
	var a Array
	var err error
	var ai js.Value
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return a, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(ai) {
		return a, ErrNotAnArray
	}
	a.BaseObject = a.SetObject(obj)
	return a, err
}

func (a Array) Length() (int, error) {
	return a.GetAttributeInt("length")
}

func (a Array) Concat(a2 Array) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	if obj, err = a.Call("concat", a2.JSObject()); err != nil {
		return newArr, err
	}
	return NewFromJSObject(obj)

}

func (a Array) CopyWithin(cible int, opts ...int) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{} = []interface{}{js.ValueOf(cible)}
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	if obj, err = a.Call("copyWithin", arrayJS...); err != nil {
		return newArr, err
	}
	return NewFromJSObject(obj)

}

func (a Array) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator
	if obj, err = a.Call("entries"); err != nil {
		return iter, err
	}
	return iterator.NewFromJSObject(obj)
}

func (a Array) Every(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("every", jsfunc); err != nil {
		jsfunc.Release()
		return result, err
	}
	jsfunc.Release()
	if obj.Type() != js.TypeBoolean {
		return result, baseobject.ErrObjectNotBool
	}
	return obj.Bool(), nil
}

//Fill (value, begin, end)
func (a Array) Fill(i interface{}, opts ...int) error {
	var err error
	var arrayJS []interface{} = []interface{}{js.ValueOf(i)}
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	_, err = a.Call("fill", arrayJS...)
	return err
}

func (a Array) Filter(f func(interface{}) bool) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("filter", jsfunc); err != nil {
		jsfunc.Release()
		return newArr, err
	}
	jsfunc.Release()
	return NewFromJSObject(obj)
}

func (a Array) Find(f func(interface{}) bool) (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("find", jsfunc); err != nil {
		jsfunc.Release()
		return i, err
	}
	jsfunc.Release()
	if obj.Type() == js.TypeUndefined {
		return i, err
	}
	return baseobject.GoValue_(obj), nil
}

func (a Array) FindIndex(f func(interface{}) bool) (int, error) {
	var err error
	var obj js.Value
	var index int = -1
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("findIndex", jsfunc); err != nil {
		jsfunc.Release()
		return index, err
	}
	jsfunc.Release()
	if obj.Type() != js.TypeNumber {
		return index, nil
	}
	return obj.Int(), nil
}

func (a Array) Flat(opts ...int) (Array, error) {
	var err error
	var arrayJS []interface{}
	var obj js.Value
	var newArr Array
	if len(opts) < 2 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}
	if obj, err = a.Call("flat", arrayJS...); err != nil {
		return newArr, err
	}
	return NewFromJSObject(obj)
}

func (a Array) FlatMap(f func(interface{}, int) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]), args[1].Int())
		return b
	})
	if obj, err = a.Call("flatMap", jsfunc); err != nil {
		jsfunc.Release()
		return newArr, err
	}
	jsfunc.Release()
	return NewFromJSObject(obj)
}

func (a Array) ForEach(f func(interface{})) error {
	var err error
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(baseobject.GoValue_(args[0]))
		return nil
	})
	_, err = a.Call("forEach", jsfunc)
	jsfunc.Release()
	return err
}

func (a Array) Includes(i interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	var includecheck js.Value
	if objGo, ok := i.(baseobject.ObjectFrom); ok {
		includecheck = objGo.JSObject()
	} else {
		includecheck = js.ValueOf(i)
	}
	if obj, err = a.Call("includes", includecheck); err != nil {
		return result, err
	}
	if obj.Type() != js.TypeBoolean {
		return result, baseobject.ErrObjectNotBool
	}
	return obj.Bool(), nil
}

func (a Array) IndexOf(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1
	var indexCheck js.Value
	if objGo, ok := i.(baseobject.ObjectFrom); ok {
		indexCheck = objGo.JSObject()
	} else {
		indexCheck = js.ValueOf(i)
	}
	if obj, err = a.Call("indexOf", indexCheck); err != nil {
		return index, err
	}
	if obj.Type() != js.TypeNumber {
		return index, nil
	}
	return obj.Int(), nil
}

func IsArray(bobj baseobject.BaseObject) (bool, error) {
	var err error
	var result bool
	var ai, obj js.Value
	if ai = GetInterface(); ai.IsUndefined() {
		return result, ErrNotImplemented
	}
	if obj, err = baseobject.Call(ai, "isArray", bobj.JSObject()); err != nil {
		return result, err
	}
	if obj.Type() != js.TypeBoolean {
		return result, baseobject.ErrObjectNotBool
	}
	return obj.Bool(), nil
}

func (a Array) Join(separator string) (string, error) {
	var err error
	var result string
	var obj js.Value
	if obj, err = a.Call("join", js.ValueOf(separator)); err != nil {
		return result, err
	}
	if obj.Type() != js.TypeString {
		return result, baseobject.ErrObjectNotString
	}
	return obj.String(), nil
}

func (a Array) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator
	if obj, err = a.Call("keys"); err != nil {
		return iter, err
	}
	return iterator.NewFromJSObject(obj)
}

func (a Array) LastIndexOf(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1
	var indexCheck js.Value
	if objGo, ok := i.(baseobject.ObjectFrom); ok {
		indexCheck = objGo.JSObject()
	} else {
		indexCheck = js.ValueOf(i)
	}
	if obj, err = a.Call("lastIndexOf", indexCheck); err != nil {
		return index, err
	}
	if obj.Type() != js.TypeNumber {
		return index, nil
	}
	return obj.Int(), nil
}

func (a Array) Map(f func(interface{}) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("map", jsfunc); err != nil {
		jsfunc.Release()
		return newArr, err
	}
	jsfunc.Release()
	return NewFromJSObject(obj)
}

func (a Array) Pop() error {
	var err error
	_, err = a.Call("pop")
	return err
}

func (a Array) Push(i interface{}) (int, error) {
	var err error
	var obj js.Value
	var index int = -1
	var pushdata js.Value
	if objGo, ok := i.(baseobject.ObjectFrom); ok {
		pushdata = objGo.JSObject()
	} else {
		pushdata = js.ValueOf(i)
	}
	if obj, err = a.Call("push", pushdata); err != nil {
		return index, err
	}
	if obj.Type() != js.TypeNumber {
		return index, nil
	}
	return obj.Int(), nil
}

func (a Array) Reduce(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj js.Value
	var newValue interface{}
	var argCall []interface{}
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, js.ValueOf(args[i]))
		}
		b := f(baseobject.GoValue_(args[0]), baseobject.GoValue_(args[1]), arrayJS...)
		return js.ValueOf(b)
	})
	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj, err = a.Call("reduce", argCall...); err != nil {
		return newValue, err
	}
	jsfunc.Release()
	return baseobject.GoValue_(obj), nil
}

func (a Array) ReduceRight(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue ...interface{}) (interface{}, error) {
	var err error
	var obj js.Value
	var newValue interface{}
	var argCall []interface{}
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var arrayJS []interface{}
		for i := 2; i < len(args); i++ {
			arrayJS = append(arrayJS, js.ValueOf(args[i]))
		}
		b := f(baseobject.GoValue_(args[0]), baseobject.GoValue_(args[1]), arrayJS...)
		return js.ValueOf(b)
	})
	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj, err = a.Call("reduceRight", argCall...); err != nil {
		return newValue, err
	}
	jsfunc.Release()
	return baseobject.GoValue_(obj), nil
}

func (a Array) Reverse() error {
	var err error
	_, err = a.Call("reverse")
	return err
}

func (a Array) Shift() (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}
	if obj, err = a.Call("shift"); err != nil {
		return i, err
	}
	return baseobject.GoValue_(obj), nil
}

func (a Array) Slice(opts ...int) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{}
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	if obj, err = a.Call("slice", arrayJS...); err != nil {
		return newArr, err
	}
	return NewFromJSObject(obj)
}

func (a Array) Some(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue_(args[0]))
		return js.ValueOf(b)
	})
	if obj, err = a.Call("some", jsfunc); err != nil {
		jsfunc.Release()
		return result, err
	}
	jsfunc.Release()
	if obj.Type() != js.TypeBoolean {
		return result, baseobject.ErrObjectNotBool
	}
	return obj.Bool(), nil
}

func (a Array) Sort() error {
	var err error
	_, err = a.Call("sort")
	return err
}

func (a Array) Splice(begin, suppress int, values ...interface{}) error {
	var err error
	var arrayJS []interface{} = []interface{}{js.ValueOf(begin), js.ValueOf(suppress)}
	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	_, err = a.Call("splice", arrayJS...)
	return err
}

func (a Array) ToLocaleString() (string, error) {
	return a.GetAttributeString("toLocaleString")
}

func (a Array) Unshift(values ...interface{}) (int, error) {
	var err error
	var arrayJS []interface{}
	var obj js.Value
	var index int = -1
	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if obj, err = a.Call("unshift", arrayJS...); err != nil {
		return index, err
	}
	if obj.Type() != js.TypeNumber {
		return index, nil
	}
	return obj.Int(), nil
}

func (a Array) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator
	if obj, err = a.Call("values"); err != nil {
		return iter, err
	}
	return iterator.NewFromJSObject(obj)
}

func (a Array) SetValue(index int, i interface{}) error {
	a.JSObject().SetIndex(index, baseobject.GetJsValueOf(i))
	return nil
}

func (a Array) GetValue(index int) (interface{}, error) {
	obj := a.JSObject().Index(index)
	return baseobject.GoValue_(obj), nil
}
