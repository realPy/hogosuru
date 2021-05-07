package document

import (
	"sync"

	"github.com/realPy/jswasm/js"
)

var singleton sync.Once

var docinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var docinstance JSInterface
		var err error
		if docinstance.objectInterface, err = js.Global().GetWithErr("document"); err == nil {
			docinterface = &docinstance
		}
	})

	return docinterface
}

//Root Get the root obj document
func Root() js.Value {

	return GetJSInterface().objectInterface
}

func QuerySelector(selector string) (js.Value, error) {

	return GetJSInterface().objectInterface.CallWithErr("querySelector", js.ValueOf(selector))
}
