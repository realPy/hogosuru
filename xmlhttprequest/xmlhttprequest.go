package xmlhttprequest

import (
	"sync"

	"github.com/realPy/jswasm/js"
)

var singleton sync.Once

var xhrinterface *JSInterface

//JSInterface of XML HTTP Request
type JSInterface struct {
	xhrinterface js.Value
}

//XMLHTTPRequest XMLHTTPRequest struct
type XMLHTTPRequest struct {
	xhrobject js.Value
}

//GetJSInterface Get the JS XMLHTTPRequest Interface If nil browser doesn't implement it
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var xhrinstance JSInterface
		var err error
		if xhrinstance.xhrinterface, err = js.Global().GetWithErr("XMLHttpRequest"); err == nil {
			xhrinterface = &xhrinstance
		}
	})

	return xhrinterface
}

//NewXMLHTTPRequest Get an XML HTTP Request
func NewXMLHTTPRequest() (XMLHTTPRequest, error) {
	var request XMLHTTPRequest

	if xhri := GetJSInterface(); xhri != nil {

		request.xhrobject = xhri.xhrinterface.New()
		return request, nil

	}
	return request, ErrNotImplemented
}
