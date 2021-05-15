package element

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var elementinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var elementinstance JSInterface
		var err error
		if elementinstance.objectInterface, err = js.Global().GetWithErr("Element"); err == nil {
			elementinterface = &elementinstance
		}
	})

	return elementinterface
}

type Element struct {
	node.Node
}

func New() Element {

	var e Element
	if ei := GetJSInterface(); ei != nil {
		e.Object = e.SetObject(ei.objectInterface.New())
		return e
	}

	e.Error = &ErrNotImplemented
	return e
}

func NewFromJSObject(obj js.Value) Element {
	var e Element

	if ei := GetJSInterface(); ei != nil {
		if obj.InstanceOf(ei.objectInterface) {
			e.Object = e.SetObject(obj)
			return e
		}

	}

	e.Error = &ErrNotAnElement
	return e
}
