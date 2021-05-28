package nodelist

// https://developer.mozilla.org/fr/docs/Web/API/NodeList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var nodelistinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//NodeList struct
type NodeList struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var nodelistinstance JSInterface
		var err error
		if nodelistinstance.objectInterface, err = js.Global().GetWithErr("NodeList"); err == nil {
			nodelistinterface = &nodelistinstance
		}
	})

	return nodelistinterface
}

func NewFromJSObject(obj js.Value) (NodeList, error) {
	var n NodeList

	if nli := GetJSInterface(); nli != nil {
		if obj.InstanceOf(nli.objectInterface) {
			n.BaseObject = n.SetObject(obj)
			return n, nil
		}
	}
	return n, ErrNotAnNodeList
}

func (n NodeList) Item(index int) (node.Node, error) {

	return node.NewFromJSObject(n.JSObject().Index(index))
}