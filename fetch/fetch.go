package fetch

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/realPy/jswasm"
	"github.com/realPy/jswasm/arraybuffer"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
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
	object js.Value
}

//FetchResponse struct
type FetchResponse struct {
	object js.Value
	err    error
	status int
}

func NewFetchResponse(obj js.Value) (FetchResponse, error) {
	var response FetchResponse

	if object.String(obj) == "[object Response]" {

		response.object = obj
		return response, nil
	}
	return response, ErrNotAnFResp
}

func (fr FetchResponse) Status() int {
	if fr.status == 0 {
		fr.status = 456
		if statusObject, err := fr.object.GetWithErr("status"); err == nil {
			if statusObject.Type() == js.TypeNumber {
				fr.status = statusObject.Int()
			}
		}
	}
	return fr.status
}

func (fr FetchResponse) Text() (string, error) {

	var txtObject js.Value
	var err error
	if txtObject, err = fr.object.CallWithErr("text"); err == nil {
		jsTxt := <-jswasm.Await(txtObject)
		if len(jsTxt) > 0 {
			return jsTxt[0].String(), nil
		}

	}
	return "", err
}

func (fr FetchResponse) ArrayBuffer() ([]byte, error) {

	var buffer []byte
	var err error
	var arrayObject js.Value
	var ab arraybuffer.ArrayBuffer
	if arrayObject, err = fr.object.CallWithErr("arrayBuffer"); err == nil {
		binary := <-jswasm.Await(arrayObject)

		if len(binary) > 0 {

			if ab, err = arraybuffer.NewArrayBuffer(binary[0]); err == nil {

				return ab.Bytes()
			}
		}
	}
	return buffer, err
}

//NewFetch New fetch
func NewFetch(urlfetch *url.URL, method string, headers *map[string]interface{}, data *url.Values, handlerResponse func(FetchResponse)) (Fetch, error) {
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

		fetch.object = fetchi.objectInterface.Invoke(urlfetch.String(), arg)

		then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			var err error
			var fr FetchResponse
			if len(args) > 0 {
				rsp := args[0]
				if fr, err = NewFetchResponse(rsp); err != nil {
					fr.err = err
				}

			} else {
				fr.err = fmt.Errorf("fetch response must contains args")
			}
			handlerResponse(fr)
			return nil
		})

		fetch.object.Call("then", then)

		return fetch, nil
	}
	return fetch, ErrNotImplemented
}
