package promise

import (
	"sync"
	"syscall/js"

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

func New(handler func(Promise) interface{}) (Promise, error) {

	var p Promise
	var err error

	if pi := GetInterface(); !pi.IsNull() {
		fh := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			return handler(p)
		})

		p.BaseObject = p.SetObject(pi.New(fh))

	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func (p Promise) Await(awaitable js.Value) chan []js.Value {
	ch := make(chan []js.Value)
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args
		return nil
	})
	awaitable.Call("then", cb)
	return ch
}
