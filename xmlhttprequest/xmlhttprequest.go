package xmlhttprequest

import (
	"sync"

	"github.com/realPy/jswasm/js"
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
