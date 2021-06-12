package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"net/url"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/promise"
	jsresponse "github.com/realPy/hogosuru/response"
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
	})

	return fetchinterface
}

//Fetch struct
type Fetch struct {
	promise.Promise
}

func NewFetch(urlfetch *url.URL, method string, headers *map[string]interface{}, data *url.Values, handlerResponse func(jsresponse.Response, error)) (Fetch, error) {
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

		promisefetchobj := fetchi.Invoke(urlfetch.String(), arg)
		if p, err = promise.NewFromJSObject(promisefetchobj); err == nil {

			p.Async(func(v js.Value) *promise.Promise {

				var r jsresponse.Response
				r, err = jsresponse.NewFromJSObject(v)
				handlerResponse(r, err)

				return nil
			}, func(e error) {
				handlerResponse(jsresponse.Response{}, err)
			})
			fetch.BaseObject = fetch.SetObject(p.JSObject())
		}

	} else {
		err = ErrNotImplemented
	}
	return fetch, err
}
