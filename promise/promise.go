package promise

import (
	"errors"
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
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
	return promiseinterface
}

//Promise struct
type Promise struct {
	baseobject.BaseObject
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

func (p Promise) Async(resolve func(js.Value) *Promise, reject func(error)) error {
	var err error
	resolveFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) > 0 {
			if p := resolve(args[0]); p != nil {
				return p.JSObject()
			}

		}

		return nil
	})

	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var err error
		err = errors.New(args[0].String())
		reject(err)
		return nil
	})

	_, err = p.JSObject().CallWithErr("then", resolveFunc, rejectedFunc)
	return err
}

func (p Promise) All(arrpromise array.Array) (Promise, error) {
	var err error
	var pr Promise
	var promiseobj js.Value

	if promiseobj, err = p.JSObject().CallWithErr("all", arrpromise.JSObject()); err == nil {
		pr, err = NewFromJSObject(promiseobj)

	}
	return pr, err
}

func (p Promise) Catch(reject func(error)) error {
	var err error
	rejectedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var err error
		err = errors.New(args[0].String())
		reject(err)
		return nil
	})
	_, err = p.JSObject().CallWithErr("catch", rejectedFunc)
	return err
}

func (p Promise) Await() (baseobject.BaseObject, error) {
	var obj baseobject.BaseObject
	var err error
	var ok bool

	ch := make(chan interface{})

	err = p.Async(func(v js.Value) *Promise {
		ch <- v
		return nil
	}, func(e error) {

		ch <- e
	})
	returnvalue := <-ch

	if err, ok = returnvalue.(error); !ok {
		if jsobj, ok := returnvalue.(js.Value); ok {

			obj, err = baseobject.NewFromJSObject(jsobj)

		} else {
			err = ErrResultPromiseError
		}

	}

	return obj, err
}
