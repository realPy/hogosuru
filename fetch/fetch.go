package fetch

import (
	"net/url"
	"sync"

	"github.com/realPy/jswasm"
	"github.com/realPy/jswasm/js"
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

func handleHTTPBytesResult(rsp js.Value, httpHandler func(int, []byte)) {

	if statusObject, err := rsp.GetWithErr("status"); err == nil {
		if statusObject.Type() == js.TypeNumber {
			status := statusObject.Int()

			if arrayObject, err := rsp.CallWithErr("arrayBuffer"); err == nil {
				binary := <-jswasm.Await(arrayObject)

				if arrayConstructor, err := js.Global().GetWithErr("Uint8Array"); err == nil {
					dataJS := arrayConstructor.New(binary[0])

					binaryObject := binary[0]
					//	len := binary[0].Get("byteLength").Int()
					if byteLengthObject, err := binaryObject.GetWithErr("byteLength"); err == nil {
						if byteLengthObject.Type() == js.TypeNumber {
							len := byteLengthObject.Int()
							var data []byte = make([]byte, len)
							if _, err := js.CopyBytesToGoWithErr(data, dataJS); err == nil {
								httpHandler(status, data)
							} else {
								httpHandler(456, []byte(err.Error()))
							}

						} else {
							httpHandler(456, []byte("byteLength is not an number"))
						}

					} else {
						httpHandler(456, []byte(err.Error()))
					}

				} else {
					httpHandler(456, []byte("unable to allocate Uint8Array constructor"))
				}
			} else {
				httpHandler(456, []byte(err.Error()))
			}
		} else {
			httpHandler(456, []byte("status is not an number"))
		}

	} else {
		httpHandler(456, []byte(err.Error()))
	}

}
func handleHTTPTextResult(rsp js.Value, httpHandler func(int, string)) {

	if statusObject, err := rsp.GetWithErr("status"); err == nil {
		if statusObject.Type() == js.TypeNumber {
			status := statusObject.Int()
			if txtObject, err := rsp.CallWithErr("text"); err == nil {
				jsTxt := <-jswasm.Await(txtObject)
				httpHandler(status, jsTxt[0].String())

			} else {
				httpHandler(456, err.Error())
			}

		} else {
			httpHandler(456, "status is not an number")
		}
	} else {
		httpHandler(456, err.Error())
	}

}

func http(url *url.URL, arg js.Value, resultHandler func(js.Value, error)) {
	go func() {
		if fetchi := GetJSInterface(); fetchi != nil {
			ch := jswasm.Await(fetchi.objectInterface.Invoke(url.String(), arg))
			go func() {
				results := <-ch
				rsp := results[0]
				resultHandler(rsp, nil)
			}()
		} else {
			resultHandler(js.Value{}, ErrNotImplemented)

		}
	}()
}
func httpGetRequest(url *url.URL, resultHandler func(js.Value, error)) {
	http(url, js.ValueOf(nil), resultHandler)
}

func httpPost(url *url.URL, data *url.Values, resultHandler func(js.Value, error)) {
	arg := js.ValueOf(map[string]interface{}{"method": "POST", "headers": map[string]interface{}{"content-type": "application/x-www-form-urlencoded"}, "body": data.Encode()})
	http(url, arg, resultHandler)
}

//HTTPGetText get a url ressource with string response
func HTTPGetText(url *url.URL, httpHandler func(int, string)) {
	httpGetRequest(url, func(rsp js.Value, err error) {
		if err == nil {
			handleHTTPTextResult(rsp, httpHandler)
		} else {
			httpHandler(456, err.Error())
		}

	})
}

//HTTPGetBytes get a url ressource with bytes response
func HTTPGetBytes(url *url.URL, httpHandler func(int, []byte)) {
	httpGetRequest(url, func(rsp js.Value, err error) {
		if err == nil {
			handleHTTPBytesResult(rsp, httpHandler)
		} else {
			httpHandler(456, []byte(err.Error()))
		}

	})
}

//HTTPPostText post data to url with text response
func HTTPPostText(url *url.URL, data *url.Values, httpHandler func(int, string)) {

	httpPost(url, data, func(rsp js.Value, err error) {
		if err == nil {
			handleHTTPTextResult(rsp, httpHandler)
		} else {
			httpHandler(456, err.Error())
		}
	})
}

//HTTPPostBytes post data with bytes response
func HTTPPostBytes(url *url.URL, data *url.Values, httpHandler func(int, []byte)) {

	httpPost(url, data, func(rsp js.Value, err error) {
		if err == nil {
			handleHTTPBytesResult(rsp, httpHandler)
		} else {
			httpHandler(456, []byte(err.Error()))
		}
	})
}
