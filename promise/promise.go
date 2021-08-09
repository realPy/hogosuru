package promise

import (
	"errors"
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domexception"
)

var singleton sync.Once

var promiseinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if promiseinterface, err = js.Global().GetWithErr("Promise"); err != nil {
			promiseinterface = js.Null()
		}
	})

	baseobject.Register(promiseinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})
	return promiseinterface
}

//Promise struct
type Promise struct {
	baseobject.BaseObject
}

type PromiseFrom interface {
	Promise() Promise
}

func (p Promise) Promise() Promise {
	return p
}

func New(handler func(Promise) (interface{}, error)) (Promise, error) {

	var p Promise
	var err error

	if pi := GetInterface(); !pi.IsNull() {
		fh := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if result, err := handler(p); err == nil {
				args[0].Invoke(result)
			} else {
				args[1].Invoke(err.Error())
			}

			return nil
		})

		p.BaseObject = p.SetObject(pi.New(fh))

	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func NewFromJSObject(obj js.Value) (Promise, error) {
	var p Promise
	var err error
	if pi := GetInterface(); !pi.IsNull() {
		if obj.InstanceOf(pi) {
			p.BaseObject = p.SetObject(obj)

		} else {
			err = ErrNotAPromise
		}
	} else {
		err = ErrNotImplemented
	}
	return p, err
}

func (p Promise) Async(resolve func(baseobject.BaseObject) *Promise, reject func(error)) error {
	var err error
	var obj baseobject.BaseObject
	resolveFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) > 0 {
			obj, err = baseobject.NewFromJSObject(args[0])

			if p := resolve(obj); p != nil {
				return p.JSObject()
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var err error
		var exception domexception.DomException
		if exception, err = domexception.NewFromJSObject(args[0]); err == nil {
			message, _ := exception.Message()
			err = errors.New(message)
		} else {
			err = errors.New(args[0].String())
		}

		if reject != nil {
			reject(err)
		}

		return nil
	})

	_, err = p.JSObject().CallWithErr("then", resolveFunc, rejectedFunc)
	return err
}

func (p Promise) Then(resolve func(interface{}) *Promise, reject func(error)) error {

	var err error
	var obj interface{}
	resolveFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) > 0 {
			obj, err = baseobject.Discover(args[0])
			if resolve != nil {
				if retp := resolve(obj); retp != nil {
					return retp.JSObject()
				}
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var err error
		var exception domexception.DomException
		if exception, err = domexception.NewFromJSObject(args[0]); err == nil {
			message, _ := exception.Message()
			err = errors.New(message)
		} else {
			err = errors.New(args[0].String())
		}

		if reject != nil {
			reject(err)
		}

		return nil
	})

	_, err = p.JSObject().CallWithErr("then", resolveFunc, rejectedFunc)
	return err
}

func iterablePromises(method string, values ...interface{}) (Promise, error) {
	var err error
	var pr Promise
	var promiseobj js.Value
	var arr array.Array

	var arrayJS []interface{}
	if pi := GetInterface(); !pi.IsNull() {
		for _, value := range values {
			if objGo, ok := value.(baseobject.ObjectFrom); ok {
				arrayJS = append(arrayJS, objGo.JSObject())
			} else {
				arrayJS = append(arrayJS, js.ValueOf(value))
			}

		}
		if arr, err = array.New(arrayJS...); err == nil {
			if promiseobj, err = pi.CallWithErr(method, arr.JSObject()); err == nil {
				pr, err = NewFromJSObject(promiseobj)
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return pr, err
}

func All(values ...interface{}) (Promise, error) {
	return iterablePromises("all", values...)

}
func AllSettled(values ...interface{}) (Promise, error) {
	return iterablePromises("allSettled", values...)
}

func Any(values ...interface{}) (Promise, error) {
	return iterablePromises("any", values...)
}

func Race(values ...interface{}) (Promise, error) {
	return iterablePromises("race", values...)
}

func (p Promise) Catch(reject func(error)) error {
	var err error
	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var err error
		var exception domexception.DomException
		if exception, err = domexception.NewFromJSObject(args[0]); err == nil {
			message, _ := exception.Message()
			err = errors.New(message)
		} else {
			err = errors.New(args[0].String())
		}

		if reject != nil {
			reject(err)
		}
		return nil
	})
	_, err = p.JSObject().CallWithErr("catch", rejectedFunc)
	return err
}

func (p Promise) Finally(f func()) error {
	var err error
	finallyFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	_, err = p.JSObject().CallWithErr("finally", finallyFunc)
	return err
}

func (p Promise) Await() (baseobject.BaseObject, error) {
	var obj baseobject.BaseObject
	var err error
	var ok bool

	ch := make(chan interface{})

	err = p.Async(func(obj baseobject.BaseObject) *Promise {
		ch <- obj
		return nil
	}, func(e error) {

		ch <- e
	})
	returnvalue := <-ch

	if err, ok = returnvalue.(error); !ok {
		if obj, ok = returnvalue.(baseobject.BaseObject); !ok {
			err = ErrResultPromiseError

		}
	}

	return obj, err
}
