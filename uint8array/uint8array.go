package uint8array

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var singleton sync.Once

var uint8arrayinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Uint8Array struct
type Uint8Array struct {
	object.Object
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

/*
func (j *JSInterface) New(obj js.Value) js.Value {
	return j.objectInterface.New(obj)
}
*/

func NewFromArrayObjectBuffer(obj js.Value) (Uint8Array, error) {

	uint8arrayObject := GetJSInterface().objectInterface.New(obj)
	return NewFromJSObject(uint8arrayObject)

}

func NewFromJSObject(obj js.Value) (Uint8Array, error) {
	var u Uint8Array

	u.Object = u.SetObject(obj)
	return u, nil

}
