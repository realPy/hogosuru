package response

import (
	"errors"
	"sync"

	"github.com/realPy/jswasm"
	"github.com/realPy/jswasm/arraybuffer"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
	"github.com/realPy/jswasm/uint8array"
)

var (
	ErrNotAnFResp = errors.New("The given value must be an fetch response")
)

var singleton sync.Once

var responseinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//FetchResponse struct
type Response struct {
	object.Object
	Err    error
	status int
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var responseinstance JSInterface
		var err error
		if responseinstance.objectInterface, err = js.Global().GetWithErr("Response"); err == nil {
			responseinterface = &responseinstance
		}
	})

	return responseinterface
}

//New Create a newJSEvent
func New() (Response, error) {
	var r Response

	if ri := GetJSInterface(); ri != nil {
		r.Object = r.SetObject(ri.objectInterface.New())
		return r, nil
	}
	return r, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Response, error) {
	var response Response

	if object.String(obj) == "[object Response]" {

		response.Object = response.SetObject(obj)
		return response, nil
	}
	return response, ErrNotAnFResp
}

func (r Response) Status() int {
	if r.status == 0 {
		r.status = 456
		if statusObject, err := r.JSObject().GetWithErr("status"); err == nil {
			if statusObject.Type() == js.TypeNumber {
				r.status = statusObject.Int()
			}
		}
	}
	return r.status
}

func (r Response) Text() (string, error) {

	var txtObject js.Value
	var err error
	if txtObject, err = r.JSObject().CallWithErr("text"); err == nil {
		jsTxt := <-jswasm.Await(txtObject)
		if len(jsTxt) > 0 {
			return jsTxt[0].String(), nil
		}

	}
	return "", err
}

func (r Response) ArrayBuffer() (arraybuffer.ArrayBuffer, error) {

	var ab arraybuffer.ArrayBuffer
	var err error
	var arrayObject js.Value
	if arrayObject, err = r.JSObject().CallWithErr("arrayBuffer"); err == nil {
		binary := <-jswasm.Await(arrayObject)

		if len(binary) > 0 {

			ab, err = arraybuffer.NewFromJSObject(binary[0])
		}
	}
	return ab, err

}

func (r Response) ArrayBufferBytes() ([]byte, error) {

	var buffer []byte
	var ab arraybuffer.ArrayBuffer
	var arr8buf uint8array.Uint8Array

	var err error

	if ab, err = r.ArrayBuffer(); err == nil {
		if arr8buf, err = uint8array.NewFromArrayBuffer(ab); err == nil {
			buffer, err = arr8buf.Bytes()
		}
	}

	return buffer, err
}
