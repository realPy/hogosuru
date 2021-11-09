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

//GetInterface get the JS interface
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
	var obj js.Value
	var err error
	if ri := GetInterface(); !ri.IsUndefined() {

		if obj, err = baseobject.New(ri); err == nil {
			r.BaseObject = r.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return r, err
}

func Error() (Response, error) {

	var response Response
	var err error
	var obj js.Value

	if ri := GetInterface(); !ri.IsUndefined() {

		if obj, err = baseobject.Call(ri, "error"); err == nil {
			response.BaseObject = response.SetObject(obj)

		} else {
			err = ErrNotAnFResp
		}

	} else {
		err = ErrNotImplemented
	}

	return response, err

}

func NewFromJSObject(obj js.Value) (Response, error) {
	var response Response
	var err error
	if ri := GetInterface(); !ri.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ri) {
				response.BaseObject = response.SetObject(obj)

			} else {
				err = ErrNotAnFResp
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return response, err
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
	return r.GetAttributeBool("redirected")
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

func (r Response) BodyUsed() (bool, error) {

	return r.GetAttributeBool("bodyUsed")
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

/* not exist on chrome
func (r Response) UseFinalURL() (bool, error) {

	return r.GetAttributeBool("useFinalURL")
}

func (r Response) SetUseFinalURL(b bool) {

	r.JSObject().Set("useFinalURL", js.ValueOf(b))
}*/

func (r Response) ArrayBuffer() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject, err = r.Call("arrayBuffer"); err == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err

}

func (r Response) Blob() (promise.Promise, error) {

	var promiseObject js.Value
	var p promise.Promise
	var err error
	if promiseObject, err = r.Call("blob"); err == nil {
		p, err = promise.NewFromJSObject(promiseObject)
	}
	return p, err

}

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

func (r Response) Clone() (Response, error) {

	var cloneObject js.Value
	var clone Response
	var err error
	if cloneObject, err = r.Call("clone"); err == nil {
		clone, err = NewFromJSObject(cloneObject)
	}
	return clone, err
}
