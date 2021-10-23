package xmlhttprequest

// https://developer.mozilla.org/fr/docs/Web/API/XMLHttpRequest/XMLHttpRequest
/*

TODO: Document Class




*/

import (
	"net/url"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/formdata"
	"github.com/realPy/hogosuru/progressevent"
)

var singleton sync.Once

var xhrinterface js.Value

//XMLHTTPRequest XMLHTTPRequest struct
type XMLHTTPRequest struct {
	baseobject.BaseObject
}

type XMLHTTPRequestFrom interface {
	XMLHTTPRequest_() XMLHTTPRequest
}

func (x XMLHTTPRequest) XMLHTTPRequest_() XMLHTTPRequest {
	return x
}

//GetInterface Get the JS XMLHTTPRequest Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if xhrinterface, err = js.Global().GetWithErr("XMLHttpRequest"); err != nil {
			xhrinterface = js.Undefined()
		}
		baseobject.Register(xhrinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return xhrinterface
}

func NewFromJSObject(obj js.Value) (XMLHTTPRequest, error) {
	var x XMLHTTPRequest
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			x.BaseObject = x.SetObject(obj)

		}
	} else {
		err = ErrNotAXMLHTTPRequest
	}

	return x, err
}

//New Get an XML HTTP Request
func New() (XMLHTTPRequest, error) {
	var request XMLHTTPRequest

	if xhri := GetInterface(); !xhri.IsUndefined() {

		request.BaseObject = request.SetObject(xhri.New())
		return request, nil

	}
	return request, ErrNotImplemented
}

func (x XMLHTTPRequest) Open(method string, url *url.URL) error {
	var err error
	_, err = x.JSObject().CallWithErr("open", js.ValueOf(method), js.ValueOf(url.String()))
	return err
}

func (x XMLHTTPRequest) SetRequestHeader(header string, value string) error {
	var err error
	_, err = x.JSObject().CallWithErr("setRequestHeader", js.ValueOf(header), js.ValueOf(value))
	return err
}

func (x XMLHTTPRequest) Send() error {
	var err error
	_, err = x.JSObject().CallWithErr("send")
	return err
}

func (x XMLHTTPRequest) SendForm(f formdata.FormData) error {
	var err error
	_, err = x.JSObject().CallWithErr("send", f.JSObject())
	return err
}

func (x XMLHTTPRequest) Abort() error {
	var err error
	_, err = x.JSObject().CallWithErr("abort")
	return err
}

func (x XMLHTTPRequest) setHandler(jshandlername string, handler func(XMLHTTPRequest)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(x)

		return nil
	})

	x.JSObject().Set(jshandlername, jsfunc)
}

//SetOnload Set OnLoad
func (x XMLHTTPRequest) SetOnload(handler func(XMLHTTPRequest)) {
	x.setHandler("onload", handler)
}

//SetOnAbort Set SetOnAbort
func (x XMLHTTPRequest) SetOnAbort(handler func(XMLHTTPRequest)) {
	x.setHandler("onabort", handler)
}

//SetOnError Set SetOnError
func (x XMLHTTPRequest) SetOnError(handler func(XMLHTTPRequest)) {
	x.setHandler("onerror", handler)
}

//SetOnReadyStateChange Set SetOnReadyStateChange
func (x XMLHTTPRequest) SetOnReadyStateChange(handler func(XMLHTTPRequest)) {
	x.setHandler("onreadystatechange", handler)
}

//SetOnProgress Set  OnProgress
func (x XMLHTTPRequest) SetOnProgress(handler func(XMLHTTPRequest, progressevent.ProgressEvent)) {
	onprogress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
			handler(x, pe)
		} else {
			println("erreur " + err.Error())
		}

		return nil
	})

	x.JSObject().Set("onprogress", onprogress)

}

func (x XMLHTTPRequest) ReadyState() (int, error) {

	return x.GetAttributeInt("readyState")
}

func (x XMLHTTPRequest) ResponseText() (string, error) {

	return x.GetAttributeString("responseText")
}

//GetResponseHeader https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (x XMLHTTPRequest) GetResponseHeader(header string) (string, error) {
	var responseHeader js.Value
	var err error
	if responseHeader, err = x.JSObject().CallWithErr("getResponseHeader", js.ValueOf(header)); err == nil {

		if responseHeader.Type() == js.TypeString {
			return responseHeader.String(), nil
		} else {
			return "", baseobject.ErrObjectNotString
		}

	}
	return "", err
}

//Response
func (x XMLHTTPRequest) Response() (js.Value, error) {
	return x.JSObject().GetWithErr("response")
}

func (x XMLHTTPRequest) SetResponseType(typeResponse string) {

	x.JSObject().Set("responseType", js.ValueOf(typeResponse))

}

func (x XMLHTTPRequest) SetWithCredentials(withcredentials bool) {

	x.JSObject().Set("withCredentials", js.ValueOf(withcredentials))

}

func (x XMLHTTPRequest) ResponseURL() (string, error) {

	return x.GetAttributeString("responseURL")
}

func (x XMLHTTPRequest) ResponseXML() (js.Value, error) {
	var responseXML js.Value
	var err error
	if responseXML, err = x.JSObject().GetWithErr("responseXML"); err == nil {
		//return a document object : TO DO IMPLEMENTATION
		return responseXML, nil

	}
	return js.Value{}, err
}

func (x XMLHTTPRequest) Status() (int, error) {
	var readystate js.Value
	var err error
	if readystate, err = x.JSObject().GetWithErr("status"); err == nil {
		if readystate.Type() == js.TypeNumber {
			return readystate.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}

	}
	return 0, err
}

func (x XMLHTTPRequest) StatusText() (string, error) {
	var responseUrl js.Value
	var err error
	if responseUrl, err = x.JSObject().GetWithErr("statusText"); err == nil {

		if responseUrl.Type() == js.TypeString {
			return responseUrl.String(), nil
		} else {
			return "", baseobject.ErrObjectNotString
		}

	}
	return "", err
}

func (x XMLHTTPRequest) uploadSetHandler(jshandlername string, handler func(XMLHTTPRequest)) {
	var uploadAbstractObject js.Value
	var err error

	if uploadAbstractObject, err = x.JSObject().GetWithErr("upload"); err == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			handler(x)

			return nil
		})

		uploadAbstractObject.Set(jshandlername, jsfunc)
	}

}

//UploadSetOnloadstart
func (x XMLHTTPRequest) UploadSetOnloadstart(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadstart", handler)

}

//UploadSetOnabort
func (x XMLHTTPRequest) UploadSetOnabort(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onabort", handler)

}

//UploadSetOnerror
func (x XMLHTTPRequest) UploadSetOnerror(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onerror", handler)

}

//UploadSetOnload
func (x XMLHTTPRequest) UploadSetOnload(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onload", handler)

}

//UploadSetOntimeout
func (x XMLHTTPRequest) UploadSetOntimeout(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("ontimeout", handler)

}

//UploadSetOnloadend
func (x XMLHTTPRequest) UploadSetOnloadend(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadend", handler)

}

//UploadSetOnprogress
func (x XMLHTTPRequest) UploadSetOnprogress(handler func(XMLHTTPRequest, progressevent.ProgressEvent)) {

	var uploadAbstractObject js.Value
	var err error

	if uploadAbstractObject, err = x.JSObject().GetWithErr("upload"); err == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
				handler(x, pe)
			}

			return nil
		})

		uploadAbstractObject.Set("onprogress", jsfunc)
	}

}
