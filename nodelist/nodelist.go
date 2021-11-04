package nodelist

// https://developer.mozilla.org/fr/docs/Web/API/NodeList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var nodelistinterface js.Value

//NodeList struct
type NodeList struct {
	baseobject.BaseObject
}

type NodeListFrom interface {
	NodeList_() NodeList
}

func (n NodeList) NodeList_() NodeList {
	return n
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if nodelistinterface, err = baseobject.Get(js.Global(), "NodeList"); err != nil {
			nodelistinterface = js.Undefined()
		}
		baseobject.Register(nodelistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return nodelistinterface
}

func NewFromJSObject(obj js.Value) (NodeList, error) {
	var n NodeList
	var err error
	if nli := GetInterface(); !nli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(nli) {
				n.BaseObject = n.SetObject(obj)

			} else {
				err = ErrNotAnNodeList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return n, err
}

func (n NodeList) Item(index int) (node.Node, error) {

	var err error
	var nd node.Node

	obj := n.JSObject().Index(index)

	if !obj.IsUndefined() {
		nd, err = node.NewFromJSObject(obj)
	}

	return nd, err
}
