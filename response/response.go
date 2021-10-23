package response

// https://developer.mozilla.org/fr/docs/Web/API/Response

import (
	"errors"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/headers"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/stream"
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

type ResponseFrom interface {
	Response_() Response
}

func (r Response) Response_() Response {
	return r
}

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if responseinterface, err = baseobject.Get(js.Global(), "Response"); err != nil {
			responseinterface = js.Undefined()
		}
		baseobject.Register(responseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		arraybuffer.GetInterface()
	})

	return responseinterface
}

//New Create a response
func New() (Response, error) {
	var r Response

	if ri := GetInterface(); !ri.IsUndefined() {
		r.BaseObject = r.SetObject(ri.New())
		return r, nil
	}
	return r, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Response, error) {
	var response Response

	if ri := GetInterface(); !ri.IsUndefined() {
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

	if obj, err = r.Get("ok"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), nil
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return false, err
}

func (r Response) Redirected() (bool, error) {
	return r.CallBool("redirected")
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

	if obj, err = r.Get("statusText"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (r Response) Type() (string, error) {

	var err error
	var obj js.Value

	if obj, err = r.Get("type"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (r Response) Url() (string, error) {

	var err error
	var obj js.Value

	if obj, err = r.Get("url"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

//deprecated (never use Await)
func (r Response) _Text() (string, error) {

	var promiseObject js.Value
	var p promise.Promise
	var jsTxtObj interface{}
	var err error
	if promiseObject, err = r.Call("text"); err == nil {
		if p, err = promise.NewFromJSObject(promiseObject); err == nil {

			if jsTxtObj, err = p.Await(); err == nil {

				if jsTxt, ok := jsTxtObj.(baseobject.ObjectFrom); ok {
					return jsTxt.JSObject().String(), nil
				} else {
					err = baseobject.ErrNotABaseObject
				}

			}
		}

	}
	return "", err
}

func (r Response) Text() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject, err = r.Call("text"); err == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err
}

func (r Response) Json() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject, err = r.Call("json"); err == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err
}

func (r Response) UseFinalURL() (bool, error) {

	return r.CallBool("useFinalURL")
}

func (r Response) SetUseFinalURL(b bool) {

	r.JSObject().Set("useFinalURL", js.ValueOf(b))
}

func (r Response) ArrayBuffer_() (arraybuffer.ArrayBuffer, error) {

	var ab arraybuffer.ArrayBuffer
	var err error
	var promiseObject js.Value
	var p promise.Promise
	var binaryObj interface{}

	if promiseObject, err = r.Call("arrayBuffer"); err == nil {
		if p, err = promise.NewFromJSObject(promiseObject); err == nil {

			if binaryObj, err = p.Await(); err == nil {
				if binary, ok := binaryObj.(arraybuffer.ArrayBufferFrom); ok {
					ab = binary.ArrayBuffer_()
				} else {
					err = baseobject.ErrNotABaseObject
				}

			}

		}

	}
	return ab, err

}
func (r Response) ArrayBuffer() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject, err = r.Call("arrayBuffer"); err == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err

}

/*
func (r Response) ArrayBufferBytes() ([]byte, error) {

	var buffer []byte
	var ab arraybuffer.ArrayBuffer
	var arr8buf typedarray.Uint8Array

	var err error

	if ab, err = r.ArrayBuffer(); err == nil {
		if arr8buf, err = typedarray.NewUint8Array(ab); err == nil {
			buffer, err = arr8buf.Bytes()
		}
	}

	return buffer, err
}
*/
func (r Response) Headers() (headers.Headers, error) {
	var obj js.Value
	var err error
	var h headers.Headers
	if obj, err = r.Get("headers"); err == nil {
		h, err = headers.NewFromJSObject(obj)

	}
	return h, err
}

func (r Response) Body() (stream.ReadableStream, error) {
	var obj js.Value
	var err error
	var s stream.ReadableStream
	if obj, err = r.Get("body"); err == nil {
		s, err = stream.NewFromJSObject(obj)

	}
	return s, err
}
