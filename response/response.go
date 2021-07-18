package response

// https://developer.mozilla.org/fr/docs/Web/API/Response

import (
	"errors"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/uint8array"
)

var (
	ErrNotAnFResp = errors.New("The given value must be an fetch response")
)

var singleton sync.Once

var responseinterface js.Value

//FetchResponse struct
type Response struct {
	baseobject.BaseObject
}

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if responseinterface, err = js.Global().GetWithErr("Response"); err != nil {
			responseinterface = js.Null()
		}
	})
	baseobject.Register(responseinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})
	return responseinterface
}

//New Create a response
func New() (Response, error) {
	var r Response

	if ri := GetInterface(); !ri.IsNull() {
		r.BaseObject = r.SetObject(ri.New())
		return r, nil
	}
	return r, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Response, error) {
	var response Response

	if ri := GetInterface(); !ri.IsNull() {
		if obj.InstanceOf(ri) {
			response.BaseObject = response.SetObject(obj)
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
		} else {
			err = baseobject.ErrObjectNotBool
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
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return false, err
}

func (r Response) Status() (int, error) {
	var code int
	var err error
	if code, err = r.GetAttributeInt("status"); err != nil {
		code = 456
	}
	return code, err
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

	var promiseObject js.Value
	var p promise.Promise
	var jsTxt baseobject.BaseObject
	var err error
	if promiseObject, err = r.JSObject().CallWithErr("text"); err == nil {
		if p, err = promise.NewFromJSObject(promiseObject); err == nil {

			if jsTxt, err = p.Await(); err == nil {

				return jsTxt.JSObject().String(), nil

			}
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
		} else {
			err = baseobject.ErrObjectNotBool
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
	var promiseObject js.Value
	var p promise.Promise
	var binary baseobject.BaseObject

	if promiseObject, err = r.JSObject().CallWithErr("arrayBuffer"); err == nil {
		if p, err = promise.NewFromJSObject(promiseObject); err == nil {

			if binary, err = p.Await(); err == nil {

				ab, err = arraybuffer.NewFromJSObject(binary.JSObject())
			}

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
