package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

var singleton sync.Once

var fetchinterface js.Value

//GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if fetchinterface, err = js.Global().GetWithErr("fetch"); err != nil {
			fetchinterface = js.Undefined()
		}

		response.GetInterface()
		baseobject.Register(fetchinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return fetchinterface
}

//Fetch struct
type Fetch struct {
	promise.Promise
}

type FetchFrom interface {
	Fetch_() Fetch
}

func (f Fetch) Fetch_() Fetch {
	return f
}

func New(urlfetch string, opts ...interface{}) (Fetch, error) {

	var arrayJS []interface{}
	var f Fetch
	var err error

	for _, value := range opts {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}

	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		promisefetchobj := fetchi.Invoke(urlfetch, arrayJS)
		f.BaseObject = f.SetObject(promisefetchobj)
	} else {
		err = ErrNotImplemented
	}
	return f, err

}

func NewFromJSObject(obj js.Value) (Fetch, error) {
	var h Fetch

	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		if obj.InstanceOf(fetchi) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAFetch
}

func (f Fetch) Then(resolve func(response.Response) *promise.Promise, reject func(error)) (promise.Promise, error) {

	return f.Promise.Then(func(obj interface{}) *promise.Promise {
		var resp interface{}
		var err error

		if bo, ok := obj.(baseobject.ObjectFrom); ok {
			if resp, err = baseobject.Discover(bo.JSObject()); err == nil {

				if r, ok := resp.(response.ResponseFrom); ok {
					return resolve(r.Response_())
				}
			} else {
				if reject != nil {
					reject(err)
				}

			}
		}

		return nil

	}, func(e error) {
		if reject != nil {
			reject(e)
		}

	})

}
