package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"net/url"
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
			fetchinterface = js.Null()
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
	Fetch() Fetch
}

func (f Fetch) Fetch() Fetch {
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

	if fetchi := GetInterface(); !fetchi.IsNull() {
		promisefetchobj := fetchi.Invoke(urlfetch, arrayJS)
		f.BaseObject = f.SetObject(promisefetchobj)
	} else {
		err = ErrNotImplemented
	}
	return f, err

}

func NewFromJSObject(obj js.Value) (Fetch, error) {
	var h Fetch

	if fetchi := GetInterface(); !fetchi.IsNull() {
		if obj.InstanceOf(fetchi) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAFetch
}

func (f Fetch) Then(resolve func(response.Response) *promise.Promise, reject func(error)) error {

	return f.Promise.Then(func(obj interface{}) *promise.Promise {
		var resp interface{}
		var err error

		if bo, ok := obj.(baseobject.ObjectFrom); ok {
			if resp, err = baseobject.Discover(bo.JSObject()); err == nil {

				if r, ok := resp.(response.ResponseFrom); ok {
					return resolve(r.Response())
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

//deprecated for backward compatibilities

func NewFetch(urlfetch string, method string, headers *map[string]interface{}, data *url.Values, handlerResponse func(response.Response, error)) (Fetch, error) {

	var fetch Fetch
	var err error
	var p promise.Promise
	if fetchi := GetInterface(); !fetchi.IsNull() {
		var goarg map[string]interface{} = make(map[string]interface{})

		goarg["method"] = method
		if headers != nil {
			goarg["headers"] = *headers
		}
		if data != nil {
			goarg["body"] = data.Encode()
		}

		if headers == nil {
			headers = &map[string]interface{}{}

		}
		if data == nil {
			data = &url.Values{}
		}

		arg := js.ValueOf(goarg)

		promisefetchobj := fetchi.Invoke(urlfetch, arg)
		if p, err = promise.NewFromJSObject(promisefetchobj); err == nil {

			if handlerResponse != nil {
				p.Async(func(obj baseobject.BaseObject) *promise.Promise {

					var r response.Response
					r, err = response.NewFromJSObject(obj.JSObject())
					handlerResponse(r, err)

					return nil
				}, func(e error) {
					handlerResponse(response.Response{}, err)
				})
			}

			fetch.BaseObject = fetch.SetObject(p.JSObject())
		}

	} else {
		err = ErrNotImplemented
	}

	fetch.Debug("❗❗Use of fetch.NewFetch is deprecated❗❗")
	return fetch, err
}
