package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/response"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var fetchinterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if fetchinterface, err = baseobject.Get(js.Global(), "fetch"); err != nil {
			fetchinterface = js.Undefined()
		}

		response.GetInterface()
		baseobject.Register(fetchinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return fetchinterface
}

// Fetch struct
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
	arrayJS = append(arrayJS, urlfetch)
	for _, value := range opts {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		promisefetchobj := fetchi.Invoke(arrayJS...)
		f.BaseObject = f.SetObject(promisefetchobj)
	} else {
		err = ErrNotImplemented
	}
	return f, err

}

func NewFromJSObject(obj js.Value) (Fetch, error) {
	var h Fetch
	var err error
	if fetchi := GetInterface(); !fetchi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fetchi) {

				h.BaseObject = h.SetObject(obj)
			} else {
				err = ErrNotAFetch
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
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
