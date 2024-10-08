package promise

import (
	"errors"
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/domexception"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/jserror"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var promiseinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if promiseinterface, err = baseobject.Get(js.Global(), "Promise"); err != nil {
			promiseinterface = js.Undefined()
		}
		baseobject.Register(promiseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return promiseinterface
}

// Promise struct
type Promise struct {
	baseobject.BaseObject
}

type PromiseFrom interface {
	Promise_() Promise
}

func (p Promise) Promise_() Promise {
	return p
}

func New(handler func(resolvefunc, errfunc js.Value) (interface{}, error)) (Promise, error) {

	var p Promise
	var err error
	var obj js.Value
	if pi := GetInterface(); !pi.IsUndefined() {
		fh := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if result, err := handler(args[0], args[1]); err == nil {
				if result != nil {
					args[0].Invoke(result)
				}

			} else {
				args[1].Invoke(err.Error())
			}

			return nil
		})

		if obj, err = baseobject.New(pi, fh); err == nil {
			p.BaseObject = p.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func SetTimeout(ms int) (Promise, error) {

	var p Promise
	var err error

	timeout := js.Global().Get("window").Get("setTimeout")

	p, err = New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		timeout.Invoke(resolvefunc, js.ValueOf(ms))

		return nil, nil
	})

	return p, err
}

func NewFromJSObject(obj js.Value) (Promise, error) {
	var p Promise
	var err error
	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(pi) {
				p.BaseObject = p.SetObject(obj)

			} else {
				err = ErrNotAPromise
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return p, err
}

func (p Promise) Then(resolve func(interface{}) *Promise, reject func(error)) (Promise, error) {

	var err error
	var obj interface{}
	var newp Promise
	resolveFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) > 0 {
			obj, err = baseobject.GoValue(args[0])
			if resolve != nil {
				if retp := resolve(obj); retp != nil {
					return retp.JSObject()
				}
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var errRejected error
		var exception domexception.DomException
		if exception, errRejected = domexception.NewFromJSObject(args[0]); errRejected == nil {
			message, _ := exception.Message()
			errRejected = errors.New(message)
		} else {

			var strerr string
			b, _ := baseobject.NewFromJSObject(args[0])
			if target, err := b.Get("target"); err == nil {
				t, _ := baseobject.NewFromJSObject(target)
				if targeterror, err := t.Get("error"); err == nil {
					if exception, errRejected = domexception.NewFromJSObject(targeterror); errRejected == nil {
						message, _ := exception.Message()
						errRejected = errors.New(message)
					}

					if reject != nil {
						reject(errRejected)
					}
					return nil

				}
			}

			if strerr, errRejected = baseobject.ToStringWithErr(args[0]); errRejected == nil {
				errRejected = errors.New(strerr)
			}

		}

		if reject != nil {
			reject(errRejected)
		}

		return nil
	})
	var newpromiseobj js.Value
	if newpromiseobj, err = p.Call("then", resolveFunc, rejectedFunc); err == nil {
		newp, err = NewFromJSObject(newpromiseobj)
	}
	return newp, err
}

func iterablePromises(method string, values ...interface{}) (Promise, error) {
	var err error
	var pr Promise
	var promiseobj js.Value
	var arr array.Array

	var arrayJS []interface{}
	if pi := GetInterface(); !pi.IsUndefined() {
		for _, value := range values {
			arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
		}
		if arr, err = array.New(arrayJS...); err == nil {

			if promiseobj, err = baseobject.Call(pi, method, arr.JSObject()); err == nil {
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

func (p Promise) Catch(reject func(error)) (Promise, error) {
	var err error
	var newp Promise
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
	var newpromiseobj js.Value
	if newpromiseobj, err = p.Call("catch", rejectedFunc); err == nil {
		newp, err = NewFromJSObject(newpromiseobj)
	}

	return newp, err
}

func (p Promise) Finally(f func()) error {
	var err error
	finallyFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	_, err = p.Call("finally", finallyFunc)
	return err
}

// avoid used it can deadlocks
func (p Promise) Await() (interface{}, error) {
	var obj interface{}
	var err error
	var ok bool

	ch := make(chan interface{})

	_, err = p.Then(func(i interface{}) *Promise {

		ch <- i
		return nil

	}, func(e error) {
		ch <- e
	})

	returnvalue := <-ch

	if err, ok = returnvalue.(error); !ok {

		obj = returnvalue
	}

	return obj, err
}

func Reject(reason error) (Promise, error) {
	var p Promise
	var obj js.Value
	var jserr jserror.JSError

	var err error
	if pi := GetInterface(); !pi.IsUndefined() {

		if jserr, err = jserror.New(reason.Error()); err == nil {
			if obj, err = baseobject.Call(pi, "reject", jserr.JSObject()); err == nil {

				p, err = NewFromJSObject(obj)
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func Resolve(result interface{}) (Promise, error) {
	var p Promise
	var obj js.Value
	var err error
	if pi := GetInterface(); !pi.IsUndefined() {
		if obj, err = baseobject.Call(pi, "resolve", baseobject.GetJsValueOf(result)); err == nil {
			p, err = NewFromJSObject(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return p, err
}
