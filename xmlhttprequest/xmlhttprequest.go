package xmlhttprequest

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
