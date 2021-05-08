package eventtarget

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var singleton sync.Once

var eventtargetinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var eventtargetinstance JSInterface
		var err error
		if eventtargetinstance.objectInterface, err = js.Global().GetWithErr("EventTarget"); err == nil {
			eventtargetinterface = &eventtargetinstance
		}
	})

	return eventtargetinterface
}

type EventTarget struct {
	object.Object
}

func New() (EventTarget, error) {

	var e EventTarget
	if eti := GetJSInterface(); eti != nil {
		e.Object = e.SetObject(eti.objectInterface.New())
		return e, nil
	}
	return e, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (EventTarget, error) {
	var e EventTarget

	if object.String(obj) == "[object EventTarget]" {
		e.Object = e.SetObject(obj)
		return e, nil
	}

	return e, ErrNotAnEventTarget
}
