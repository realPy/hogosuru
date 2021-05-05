package uint8array

import (
	"sync"

	"github.com/realPy/jswasm/js"
)

var singleton sync.Once

var uint8arrayinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var uint8arrayinstance JSInterface
		var err error
		if uint8arrayinstance.objectInterface, err = js.Global().GetWithErr("Uint8Array"); err == nil {
			uint8arrayinterface = &uint8arrayinstance
		}
	})

	return uint8arrayinterface
}

func (j *JSInterface) New(obj js.Value) js.Value {
	return j.objectInterface.New(obj)
}
