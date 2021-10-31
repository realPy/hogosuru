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
	var obj js.Value
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = baseobject.New(ai, js.ValueOf(size)); err == nil {
			a.BaseObject = a.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func From(iterable interface{}, f ...func(interface{}) interface{}) (Array, error) {
	var a Array
	var err error
	var obj js.Value
	var opts []interface{}
	var jsfunc js.Func

	if ai := GetInterface(); !ai.IsUndefined() {

		if objGo, ok := iterable.(baseobject.ObjectFrom); ok {
			opts = append(opts, objGo.JSObject())

		} else {
			opts = append(opts, js.ValueOf(iterable))
		}

		if f != nil && len(f) == 1 {
			jsfunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				b := f[0](baseobject.GoValue(args[0]))
				return js.ValueOf(b)
			})
			opts = append(opts, jsfunc)

		}

		if obj, err = baseobject.Call(ai, "from", opts...); err == nil {
			a.BaseObject = a.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func Of(values ...interface{}) (Array, error) {

	var a Array
	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}
	if ai := GetInterface(); !ai.IsUndefined() {
		a.BaseObject = a.SetObject(ai.Call("of", arrayJS...))
		return a, nil
	}
	return a, ErrNotImplemented

}

func New(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}
	var obj js.Value
	var err error
	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = baseobject.New(ai, arrayJS...); err == nil {
			a.BaseObject = a.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err

}

func NewFromJSObject(obj js.Value) (Array, error) {
	var a Array
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				a.BaseObject = a.SetObject(obj)

			} else {
				err = ErrNotAnArray
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Array) Length() (int, error) {

	return a.GetAttributeInt("length")

}

func (a Array) Concat(a2 Array) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array

	if obj, err = a.Call("concat", a2.JSObject()); err == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) CopyWithin(cible int, opts ...int) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(cible))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj, err = a.Call("copyWithin", arrayJS...); err == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = a.Call("entries"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (a Array) Every(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("every", jsfunc); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

//Fill (value, begin, end)
func (a Array) Fill(i interface{}, opts ...int) error {
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(i))

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
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("filter", jsfunc); err == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Find(f func(interface{}) bool) (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("find", jsfunc); err == nil {
		if obj.Type() != js.TypeUndefined {
			i = baseobject.GoValue(obj)
		}

	}
	jsfunc.Release()
	return i, err
}

func (a Array) FindIndex(f func(interface{}) bool) (int, error) {
	var err error
	var obj js.Value
	var index int = -1

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("findIndex", jsfunc); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}
	jsfunc.Release()
	return index, err
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

	if obj, err = a.Call("flat", arrayJS...); err == nil {
		newArr, err = NewFromJSObject(obj)
	}
	return newArr, err
}

func (a Array) FlatMap(f func(interface{}, int) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]), args[1].Int())
		return b
	})

	if obj, err = a.Call("flatMap", jsfunc); err == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) ForEach(f func(interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(baseobject.GoValue(args[0]))
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

	if obj, err = a.Call("includes", includecheck); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
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

	if obj, err = a.Call("indexOf", indexCheck); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}

	return index, err
}

func IsArray(bobj baseobject.BaseObject) (bool, error) {

	var err error
	var result bool
	var obj js.Value

	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = baseobject.Call(ai, "isArray", bobj.JSObject()); err == nil {
			if obj.Type() == js.TypeBoolean {
				result = obj.Bool()
			} else {
				err = baseobject.ErrObjectNotBool
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return result, err
}

func (a Array) Join(separator string) (string, error) {
	var err error
	var result string
	var obj js.Value

	if obj, err = a.Call("join", js.ValueOf(separator)); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return result, err
}

func (a Array) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = a.Call("keys"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
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

	if obj, err = a.Call("lastIndexOf", indexCheck); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}

	return index, err
}

func (a Array) Map(f func(interface{}) interface{}) (Array, error) {
	var err error
	var obj js.Value
	var newArr Array

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("map", jsfunc); err == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
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

	if obj, err = a.Call("push", pushdata); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}

	return index, err

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

		b := f(baseobject.GoValue(args[0]), baseobject.GoValue(args[1]), arrayJS...)
		return js.ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj, err = a.Call("reduce", argCall...); err == nil {
		newValue = baseobject.GoValue(obj)
	}
	jsfunc.Release()
	return newValue, err
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

		b := f(baseobject.GoValue(args[0]), baseobject.GoValue(args[1]), arrayJS...)
		return js.ValueOf(b)
	})

	argCall = append(argCall, jsfunc)
	if len(initialValue) > 0 {
		argCall = append(argCall, js.ValueOf(initialValue[0]))
	}
	if obj, err = a.Call("reduceRight", argCall...); err == nil {
		newValue = baseobject.GoValue(obj)
	}
	jsfunc.Release()
	return newValue, err
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
	if obj, err = a.Call("shift"); err == nil {
		i = baseobject.GoValue(obj)
	}
	return i, err
}

func (a Array) Slice(opts ...int) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array
	var arrayJS []interface{}
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj, err = a.Call("slice", arrayJS...); err == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Some(f func(interface{}) bool) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := f(baseobject.GoValue(args[0]))
		return js.ValueOf(b)
	})

	if obj, err = a.Call("some", jsfunc); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

func (a Array) Sort() error {
	var err error

	_, err = a.Call("sort")

	return err
}

func (a Array) Splice(begin, suppress int, values ...interface{}) error {

	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(begin))
	arrayJS = append(arrayJS, js.ValueOf(suppress))

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

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
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}
	if obj, err = a.Call("unshift", arrayJS...); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}
	return index, err
}

func (a Array) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = a.Call("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (a Array) SetValue(index int, i interface{}) error {

	var obj interface{}
	if objGo, ok := i.(baseobject.ObjectFrom); ok {

		obj = objGo.JSObject()
	} else {
		obj = i
	}

	a.JSObject().SetIndex(index, obj)
	return nil
}

func (a Array) GetValue(index int) (interface{}, error) {

	obj := a.JSObject().Index(index)
	return baseobject.GoValue(obj), nil
}
