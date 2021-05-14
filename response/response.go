package response

// https://developer.mozilla.org/fr/docs/Web/API/Response

import (
	"errors"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/uint8array"
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

//New Create a response
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

	if ri := GetJSInterface(); ri != nil {
		if obj.InstanceOf(ri.objectInterface) {
			response.Object = response.SetObject(obj)
			return response, nil
		}
	}

	return response, ErrNotAnFResp
}

func (r Response) Ok() (bool, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("ok"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), nil
		}
	}

	return false, err
}

func (r Response) Redirected() (bool, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("redirected"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), nil
		}
	}
	return false, err
}

func (r Response) Status() (int, error) {
	var err error
	var obj js.Value
	if obj, err = r.JSObject().GetWithErr("status"); err == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int(), nil
		}
	}

	return 456, err
}

func (r Response) StatusText() (string, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("statusText"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (r Response) Type() (string, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("type"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (r Response) Url() (string, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("url"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (r Response) Text() (string, error) {

	var txtObject js.Value
	var err error
	if txtObject, err = r.JSObject().CallWithErr("text"); err == nil {
		jsTxt := <-hogosuru.Await(txtObject)
		if len(jsTxt) > 0 {
			return jsTxt[0].String(), nil
		}

	}
	return "", err
}

func (r Response) UseFinalURL() (bool, error) {

	var err error
	var obj js.Value

	if obj, err = r.JSObject().GetWithErr("useFinalURL"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), nil
		}
	}
	return false, err
}

func (r Response) SetUseFinalURL(b bool) {

	r.JSObject().Set("useFinalURL", js.ValueOf(b))
}

func (r Response) ArrayBuffer() (arraybuffer.ArrayBuffer, error) {

	var ab arraybuffer.ArrayBuffer
	var err error
	var arrayObject js.Value
	if arrayObject, err = r.JSObject().CallWithErr("arrayBuffer"); err == nil {
		binary := <-hogosuru.Await(arrayObject)

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
