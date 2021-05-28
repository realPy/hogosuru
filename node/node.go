package node

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
)

var singleton sync.Once

var nodeinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var nodeinstance JSInterface
		var err error
		if nodeinstance.objectInterface, err = js.Global().GetWithErr("Node"); err == nil {
			nodeinterface = &nodeinstance
		}
	})

	return nodeinterface
}

//we use here aenw method of chaining method

type Node struct {
	eventtarget.EventTarget
	//Error *error
}

/*
func (n Node) NotError() bool {

	if n.Error == nil || (*n.Error) == nil {
		return true
	}
	return false
}

func (n *Node) IsError() error {
	var err *error
	err = n.Error
	n.Error = nil
	return *err
}
*/

/*
func New() Node {

	var n Node
	if ni := GetJSInterface(); ni != nil {
		n.Object = n.SetObject(ni.objectInterface.New())
		return n
	}

	n.Error = &ErrNotImplemented
	return n
}*/

func NewFromJSObject(obj js.Value) (Node, error) {
	var n Node
	var err error

	if ni := GetJSInterface(); ni != nil {

		if obj.InstanceOf(ni.objectInterface) {
			n.BaseObject = n.SetObject(obj)

		} else {
			err = ErrNotANode
		}

	}

	return n, err
}

func (n Node) getAttributeNode(attribute string) (Node, error) {
	var nodeObject js.Value
	var newNode Node
	var err error

	if nodeObject, err = n.JSObject().GetWithErr(attribute); err == nil {

		if nodeObject.IsNull() {
			err = ErrNodeNoChilds

		} else {

			newNode, err = NewFromJSObject(nodeObject)

		}

	}

	return newNode, err
}

func (n Node) getAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var val string

	if obj, err = n.JSObject().GetWithErr(attribute); err == nil {

		val = obj.String()
	}
	return val, err
}

func (n Node) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var result bool

	if obj, err = n.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (n Node) getAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	if obj, err = n.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Int()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (n Node) BaseURI() (string, error) {

	return n.getAttributeString("baseURI")
}

func (n Node) FirstChild() (Node, error) {

	return n.getAttributeNode("firstChild")
}

func (n Node) IsConnected() (bool, error) {

	return n.getAttributeBool("isConnected")
}

func (n Node) LastChild() (Node, error) {
	return n.getAttributeNode("lastChild")
}

func (n Node) NextSibling() (Node, error) {
	return n.getAttributeNode("nextSibling")
}

func (n Node) NodeName() (string, error) {

	return n.getAttributeString("nodeName")

}

func (n Node) NodeType() (int, error) {
	return n.getAttributeInt("nodeType")
}

func (n Node) NodeValue() (Node, error) {
	return n.getAttributeNode("nodeValue")
}

func (n Node) SetNodeValue(nset Node) error {

	return n.JSObject().SetWithErr("nodeValue", nset.JSObject())
}

func (n Node) OwnerDocument() (Node, error) {
	return n.getAttributeNode("ownerDocument")
}

func (n Node) ParentNode() (Node, error) {
	return n.getAttributeNode("parentNode")

}

func (n Node) ParentElement() (Node, error) {
	return n.getAttributeNode("parentElement")
}

func (n Node) PreviousSibling() (Node, error) {

	return n.getAttributeNode("previousSibling")
}

func (n Node) TextContent() (string, error) {

	return n.getAttributeString("textContent")
}

func (n Node) SetTextContent(content string) error {

	return n.JSObject().SetWithErr("textContent", js.ValueOf(content))
}

func (n Node) AppendChild(add Node) error {

	_, err := n.JSObject().CallWithErr("appendChild", add.JSObject())
	return err
}

func (n Node) CloneNode(deep bool) (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj, err = n.JSObject().CallWithErr("cloneNode", js.ValueOf(deep)); err == nil {
		return NewFromJSObject(obj)
	}

	return newNode, err
}

func (n Node) CompareDocumentPosition(node Node) (int, error) {
	var err error
	var obj js.Value
	var result int

	if obj, err = n.JSObject().CallWithErr("compareDocumentPosition", node.JSObject()); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}
	return result, err

}

func (n Node) Contains(node Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = n.JSObject().CallWithErr("contains", node.JSObject()); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (n Node) GetRootNode() (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj, err = n.JSObject().CallWithErr("getRootNode"); err == nil {
		newNode, err = NewFromJSObject(obj)
	}
	return newNode, err
}

func (n Node) HasChildNodes() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.JSObject().CallWithErr("hasChildNodes"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n Node) InsertBefore(elem, before Node) (Node, error) {
	var err error

	_, err = n.JSObject().CallWithErr("insertBefore", elem, before)

	return elem, err

}

func (n *Node) IsDefaultNamespace() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.JSObject().CallWithErr("isDefaultNamespace"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n *Node) IsEqualNode() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.JSObject().CallWithErr("isEqualNode"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n *Node) IsSameNode() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.JSObject().CallWithErr("isSameNode"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n *Node) LookupPrefix() (string, error) {
	var err error
	var obj js.Value
	var result string

	if obj, err = n.JSObject().CallWithErr("lookupPrefix"); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		}
	}

	return result, err

}

func (n *Node) LookupNamespaceURI(prefix string) error {
	var err error
	_, err = n.JSObject().CallWithErr("lookupNamespaceURI", js.ValueOf(prefix))
	return err
}

func (n *Node) Normalize() error {
	var err error
	_, err = n.JSObject().CallWithErr("normalize")
	return err
}

func (n Node) RemoveChild(node Node) (Node, error) {
	var err error
	_, err = n.JSObject().CallWithErr("removeChild", node)
	return node, err

}

func (n Node) ReplaceChild(new, old Node) (Node, error) {
	var err error

	_, err = n.JSObject().CallWithErr("replaceChild", new, old)

	return old, err

}
