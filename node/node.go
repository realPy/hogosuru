package node

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/eventtarget"
)

var singleton sync.Once

var nodeinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var nodeinstance JSInterface
		var err error
		if nodeinstance.objectInterface, err = js.Global().GetWithErr("EventTarget"); err == nil {
			nodeinterface = &nodeinstance
		}
	})

	return nodeinterface
}

//we use here aenw method of chaining method

type Node struct {
	eventtarget.EventTarget
	Error *error
}

func New() Node {

	var n Node
	if ni := GetJSInterface(); ni != nil {
		n.Object = n.SetObject(ni.objectInterface.New())
		return n
	}

	n.Error = &ErrNotImplemented
	return n
}

func NewFromJSObject(obj js.Value) Node {
	var n Node

	if ni := GetJSInterface(); ni != nil {
		if obj.InstanceOf(ni.objectInterface) {
			n.Object = n.SetObject(obj)
			return n
		}

	}

	n.Error = &ErrNotANode
	return n
}

func (n Node) BaseURI() string {
	var nodeObject js.Value
	var err error

	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("baseURI"); err == nil {
			return nodeObject.String()
		}
	}

	return ""
}

func (n Node) FirstChild() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("firstChild"); err == nil {

			if nodeObject.IsNull() {
				err = ErrNodeNoChilds

			} else {

				newNode = NewFromJSObject(nodeObject)

			}

		} else {
			newNode.Error = &err
		}

	}

	return newNode
}
