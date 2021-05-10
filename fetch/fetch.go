package fetch

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
	jsresponse "github.com/realPy/hogosuru/response"
)

var singleton sync.Once

var fetchinterface *JSInterface

//JSInterface of  fetch
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var fetchinstance JSInterface
		var err error
		if fetchinstance.objectInterface, err = js.Global().GetWithErr("fetch"); err == nil {
			fetchinterface = &fetchinstance
		}
	})

	return fetchinterface
}

//Fetch struct
type Fetch struct {
	object.Object
}

//NewFetch New fetch
func NewFetch(urlfetch *url.URL, method string, headers *map[string]interface{}, data *url.Values, handlerResponse func(jsresponse.Response)) (Fetch, error) {
	var fetch Fetch

	if fetchi := GetJSInterface(); fetchi != nil {
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

		fetch.Object = fetch.SetObject(fetchi.objectInterface.Invoke(urlfetch.String(), arg))

		then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			var err error
			var r jsresponse.Response
			if len(args) > 0 {
				rsp := args[0]
				if r, err = jsresponse.NewFromJSObject(rsp); err != nil {
					r.Err = err
				}

			} else {
				r.Err = fmt.Errorf("fetch response must contains args")
			}
			handlerResponse(r)
			return nil
		})

		fetch.JSObject().Call("then", then)

		return fetch, nil
	}
	return fetch, ErrNotImplemented
}
