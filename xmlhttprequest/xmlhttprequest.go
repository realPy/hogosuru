package xmlhttprequest

/*

TODO: Document Class




*/

import (
	"net/url"
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
	"github.com/realPy/jswasm/object/event/progressevent"
)

var singleton sync.Once

var xhrinterface *JSInterface

//JSInterface of XML HTTP Request
type JSInterface struct {
	objectInterface js.Value
}

//XMLHTTPRequest XMLHTTPRequest struct
type XMLHTTPRequest struct {
	object js.Value
}

//GetJSInterface Get the JS XMLHTTPRequest Interface If nil browser doesn't implement it
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var xhrinstance JSInterface
		var err error
		if xhrinstance.objectInterface, err = js.Global().GetWithErr("XMLHttpRequest"); err == nil {
			xhrinterface = &xhrinstance
		}
	})

	return xhrinterface
}

//NewXMLHTTPRequest Get an XML HTTP Request
func NewXMLHTTPRequest() (XMLHTTPRequest, error) {
	var request XMLHTTPRequest

	if xhri := GetJSInterface(); xhri != nil {

		request.object = xhri.objectInterface.New()
		return request, nil

	}
	return request, ErrNotImplemented
}

func (x XMLHTTPRequest) Open(method string, url *url.URL) error {
	var err error
	_, err = x.object.CallWithErr("open", js.ValueOf(method), js.ValueOf(url.String()))
	return err
}

func (x XMLHTTPRequest) SetRequestHeader(header string, value string) error {
	var err error
	_, err = x.object.CallWithErr("setRequestHeader", js.ValueOf(header), js.ValueOf(value))
	return err
}

func (x XMLHTTPRequest) Send() error {
	var err error
	_, err = x.object.CallWithErr("send")
	return err
}

func (x XMLHTTPRequest) Abort() error {
	var err error
	_, err = x.object.CallWithErr("abort")
	return err
}

//SetOnload Set OnLoad
func (x XMLHTTPRequest) SetOnload(handler func(XMLHTTPRequest)) {
	onload := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(x)

		return nil
	})

	x.object.Set("onload", onload)

}

//SetOnAbort Set SetOnAbort
func (x XMLHTTPRequest) SetOnAbort(handler func(XMLHTTPRequest)) {
	onabort := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(x)

		return nil
	})
	x.object.Set("onabort", onabort)

}

//SetOnError Set SetOnError
func (x XMLHTTPRequest) SetOnError(handler func(XMLHTTPRequest)) {
	onerror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(x)

		return nil
	})
	x.object.Set("onerror", onerror)

}

//SetOnReadyStateChange Set SetOnReadyStateChange
func (x XMLHTTPRequest) SetOnReadyStateChange(handler func(XMLHTTPRequest)) {
	onreadystatechange := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(x)

		return nil
	})
	x.object.Set("onreadystatechange", onreadystatechange)

}

//SetOnProgress Set  OnProgress
func (x XMLHTTPRequest) SetOnProgress(handler func(XMLHTTPRequest, object.GOMap)) {
	onprogress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if gomap, err := progressevent.NewProgressEvent(args[0]); err == nil {
			handler(x, gomap)
		}

		return nil
	})

	x.object.Set("onprogress", onprogress)

}

func (x XMLHTTPRequest) ReadyState() (int, error) {
	var readystate js.Value
	var err error
	if readystate, err = x.object.GetWithErr("readyState"); err == nil {
		if readystate.Type() == js.TypeNumber {
			return readystate.Int(), nil
		} else {
			return 0, object.ErrObjectNotNumber
		}

	}
	return 0, err
}

func (x XMLHTTPRequest) ResponseText() (string, error) {
	var responseTexte js.Value
	var err error
	if responseTexte, err = x.object.GetWithErr("responseText"); err == nil {

		if responseTexte.Type() == js.TypeString {
			return responseTexte.String(), nil
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}

//GetResponseHeader https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (x XMLHTTPRequest) GetResponseHeader(header string) (string, error) {
	var responseHeader js.Value
	var err error
	if responseHeader, err = x.object.CallWithErr("getResponseHeader", js.ValueOf(header)); err == nil {

		if responseHeader.Type() == js.TypeString {
			return responseHeader.String(), nil
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}

//Response
func (x XMLHTTPRequest) Response() (js.Value, error) {
	return x.object.GetWithErr("response")
}

func (x XMLHTTPRequest) SetResponseType(typeResponse string) {

	x.object.Set("responseType", js.ValueOf(typeResponse))

}

func (x XMLHTTPRequest) SetWithCredentials(withcredentials bool) {

	x.object.Set("withCredentials", js.ValueOf(withcredentials))

}

func (x XMLHTTPRequest) ResponseURL() (string, error) {
	var responseUrl js.Value
	var err error
	if responseUrl, err = x.object.GetWithErr("responseURL"); err == nil {

		if responseUrl.Type() == js.TypeString {
			return responseUrl.String(), nil
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}

func (x XMLHTTPRequest) ResponseXML() (js.Value, error) {
	var responseXML js.Value
	var err error
	if responseXML, err = x.object.GetWithErr("responseXML"); err == nil {
		//return a document object : TO DO IMPLEMENTATION
		return responseXML, nil

	}
	return js.Value{}, err
}

func (x XMLHTTPRequest) Status() (int, error) {
	var readystate js.Value
	var err error
	if readystate, err = x.object.GetWithErr("status"); err == nil {
		if readystate.Type() == js.TypeNumber {
			return readystate.Int(), nil
		} else {
			return 0, object.ErrObjectNotNumber
		}

	}
	return 0, err
}

func (x XMLHTTPRequest) StatusText() (string, error) {
	var responseUrl js.Value
	var err error
	if responseUrl, err = x.object.GetWithErr("statusText"); err == nil {

		if responseUrl.Type() == js.TypeString {
			return responseUrl.String(), nil
		} else {
			return "", object.ErrObjectNotString
		}

	}
	return "", err
}
