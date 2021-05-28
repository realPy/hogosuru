package array

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

var singleton sync.Once

var arrayinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var arrayinstance JSInterface
		var err error
		if arrayinstance.objectInterface, err = js.Global().GetWithErr("Array"); err == nil {
			arrayinterface = &arrayinstance
		}
	})

	return arrayinterface
}

//Array struct
type Array struct {
	baseobject.BaseObject
}

func NewEmpty(size int) (Array, error) {

	var a Array

	if ai := GetJSInterface(); ai != nil {

		a.BaseObject = a.SetObject(ai.objectInterface.New(js.ValueOf(size)))
		return a, nil
	}
	return a, ErrNotImplemented
}

func From(iterable interface{}, f ...func(interface{}) interface{}) (Array, error) {
	var a Array
	var err error
	var obj js.Value
	var opts []interface{}
	var jsfunc js.Func

	if ai := GetJSInterface(); ai != nil {

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
		if obj, err = ai.objectInterface.CallWithErr("from", opts...); err == nil {
			a.BaseObject = a.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return a, err
}

func Of(values ...interface{}) (Array, error) {
	return New(values...)
}

func New(values ...interface{}) (Array, error) {
	var a Array
	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}
	if ai := GetJSInterface(); ai != nil {
		a.BaseObject = a.SetObject(ai.objectInterface.New(arrayJS...))
		return a, nil
	}
	return a, ErrNotImplemented

}

func NewFromJSObject(obj js.Value) (Array, error) {
	var a Array
	var err error
	if ai := GetJSInterface(); ai != nil {
		if obj.InstanceOf(ai.objectInterface) {
			a.BaseObject = a.SetObject(obj)
			return a, nil
		} else {
			err = ErrNotAnArray
		}
	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Array) Length() (int, error) {

	var LengthObject js.Value
	var err error
	if LengthObject, err = a.JSObject().GetWithErr("length"); err == nil {
		if LengthObject.Type() == js.TypeNumber {
			return LengthObject.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}
	}
	return 0, err
}

func (a Array) Concat(a2 Array) (Array, error) {

	var err error
	var obj js.Value
	var newArr Array

	if obj, err = a.JSObject().CallWithErr("concat", a2.JSObject()); err == nil {

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

	if obj, err = a.JSObject().CallWithErr("copyWithin", arrayJS...); err == nil {

		newArr, err = NewFromJSObject(obj)
	}

	return newArr, err

}

func (a Array) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = a.JSObject().CallWithErr("entries"); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("every", jsfunc); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	jsfunc.Release()
	return result, err
}

func (a Array) Fill(i interface{}, opts ...int) error {
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(i))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = a.JSObject().CallWithErr("fill", arrayJS...)
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

	if obj, err = a.JSObject().CallWithErr("filter", jsfunc); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("find", jsfunc); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("findIndex", jsfunc); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("flat", arrayJS...); err == nil {
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
		return js.ValueOf(b)
	})

	if obj, err = a.JSObject().CallWithErr("flatMap", jsfunc); err == nil {
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

	_, err = a.JSObject().CallWithErr("forEach", jsfunc)

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

	if obj, err = a.JSObject().CallWithErr("includes", includecheck); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("indexOf", indexCheck); err == nil {
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

	if ai := GetJSInterface(); ai != nil {

		if obj, err = ai.objectInterface.CallWithErr("isArray", bobj.JSObject()); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("join", js.ValueOf(separator)); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("keys"); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("lastIndexOf", indexCheck); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("map", jsfunc); err == nil {
		newArr, err = NewFromJSObject(obj)
	}
	jsfunc.Release()
	return newArr, err
}

func (a Array) Pop() error {
	var err error

	_, err = a.JSObject().CallWithErr("pop")

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

	if obj, err = a.JSObject().CallWithErr("lastIndexOf", pushdata); err == nil {
		if obj.Type() == js.TypeNumber {
			index = obj.Int()
		}
	}

	return index, err

}

func (a Array) Reduce(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue interface{}) (interface{}, error) {
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
	if initialValue != nil {
		argCall = append(argCall, js.ValueOf(initialValue))
	}
	if obj, err = a.JSObject().CallWithErr("reduce", argCall...); err == nil {
		newValue = baseobject.GoValue(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) ReduceRight(f func(accumulateur interface{}, value interface{}, opts ...interface{}) interface{}, initialValue interface{}) (interface{}, error) {
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
	if initialValue != nil {
		argCall = append(argCall, js.ValueOf(initialValue))
	}
	if obj, err = a.JSObject().CallWithErr("reduceRight", argCall...); err == nil {
		newValue = baseobject.GoValue(obj)
	}
	jsfunc.Release()
	return newValue, err
}

func (a Array) Reverse() error {
	var err error

	_, err = a.JSObject().CallWithErr("reverse")

	return err
}

func (a Array) Shift() (interface{}, error) {
	var err error
	var obj js.Value
	var i interface{}
	if obj, err = a.JSObject().CallWithErr("shift"); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("slice", arrayJS...); err == nil {

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

	if obj, err = a.JSObject().CallWithErr("some", jsfunc); err == nil {
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

	_, err = a.JSObject().CallWithErr("sort")

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
	_, err = a.JSObject().CallWithErr("splice", arrayJS...)

	return err
}

func (a Array) ToLocaleString() (string, error) {
	var err error
	var result string
	var obj js.Value

	if obj, err = a.JSObject().CallWithErr("toLocaleString"); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return result, err
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
	if obj, err = a.JSObject().CallWithErr("unshift", arrayJS...); err == nil {
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

	if obj, err = a.JSObject().CallWithErr("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}