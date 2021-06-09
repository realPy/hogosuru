package fetch

// https://developer.mozilla.org/fr/docs/Web/API/Fetch_API

import (
	"fmt"
	"net/url"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
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
	baseobject.BaseObject
}

//NewFetch New fetch
func NewFetch(urlfetch *url.URL, method string, headers *map[string]interface{}, data *url.Values, handlerResponse func(jsresponse.Response, error)) (Fetch, error) {
	var fetch Fetch

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

		fetch.BaseObject = fetch.SetObject(fetchi.Invoke(urlfetch.String(), arg))

		then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			var err error
			var r jsresponse.Response
			if len(args) > 0 {
				rsp := args[0]
				r, err = jsresponse.NewFromJSObject(rsp)

			} else {
				err = fmt.Errorf("fetch response must contains args")
			}
			handlerResponse(r, err)
			return nil
		})

		fetch.JSObject().Call("then", then)

		return fetch, nil
	}
	return fetch, ErrNotImplemented
}
