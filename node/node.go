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
	Error *error
}

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

func NewFromJSObject(obj js.Value) Node {
	var n Node

	if ni := GetJSInterface(); ni != nil {

		if obj.InstanceOf(ni.objectInterface) {
			n.BaseObject = n.SetObject(obj)

		} else {
			n.Error = &ErrNotANode
		}

	}

	return n
}

func (n Node) getAttributeNode(attribute string) Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.NotError() {
		if nodeObject, err = n.JSObject().GetWithErr(attribute); err == nil {

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

func (n *Node) BaseURI() string {
	var nodeObject js.Value
	var err error

	if n.NotError() {
		if nodeObject, err = n.JSObject().GetWithErr("baseURI"); err == nil {
			return nodeObject.String()
		} else {
			n.Error = &err
		}
	}

	return ""
}

func (n Node) FirstChild() Node {

	return n.getAttributeNode("firstChild")
}

func (n *Node) IsConnected() bool {

	var err error
	var obj js.Value

	if obj, err = n.JSObject().GetWithErr("isConnected"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		} else {
			n.Error = &baseobject.ErrObjectNotBool
		}
	} else {
		n.Error = &err
	}

	return false
}

func (n Node) LastChild() Node {
	return n.getAttributeNode("lastChild")
}

func (n Node) NextSibling() Node {
	return n.getAttributeNode("nextSibling")
}

func (n *Node) NodeName() string {

	var err error
	var obj js.Value

	if obj, err = n.JSObject().GetWithErr("nodeName"); err == nil {

		return obj.String()
	} else {
		n.Error = &err
	}
	return ""
}

func (n Node) NodeType() int {
	var err error
	var obj js.Value

	if obj, err = n.JSObject().GetWithErr("nodeType"); err == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return 0
}

func (n Node) NodeValue() Node {
	return n.getAttributeNode("nodeValue")
}

func (n *Node) SetNodeValue(nset Node) Node {

	if !n.NotError() {
		return *n
	}

	if err := n.JSObject().SetWithErr("nodeValue", nset.JSObject()); err != nil {
		n.Error = &err
	}

	return *n
}

func (n Node) OwnerDocument() Node {
	return n.getAttributeNode("ownerDocument")
}

func (n Node) ParentNode() Node {
	return n.getAttributeNode("parentNode")

}

func (n Node) ParentElement() Node {
	return n.getAttributeNode("parentElement")
}

func (n Node) PreviousSibling() Node {

	return n.getAttributeNode("previousSibling")
}

func (n *Node) TextContent() string {

	var err error
	var obj js.Value

	if n.NotError() {
		if obj, err = n.JSObject().GetWithErr("textContent"); err == nil {

			return obj.String()
		} else {
			n.Error = &err
		}
	}
	return ""
}

func (n *Node) SetTextContent(content string) Node {

	var err error
	if n.NotError() {
		if err = n.JSObject().SetWithErr("textContent", js.ValueOf(content)); err != nil {

			n.Error = &err
		}
	}
	return *n
}

func (n *Node) AppendChild(add Node) {
	var err error

	if n.NotError() {

		if _, err = n.JSObject().CallWithErr("appendChild", add.JSObject()); err != nil {
			n.Error = &err
		}

	}

}

func (n Node) CloneNode(deep bool) Node {
	var err error
	var obj js.Value
	var newNode Node

	if n.NotError() {

		if obj, err = n.JSObject().CallWithErr("cloneNode", js.ValueOf(deep)); err == nil {
			return NewFromJSObject(obj)
		}
		newNode.Error = &err
	}

	return newNode
}

func (n *Node) CompareDocumentPosition(node Node) int {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("compareDocumentPosition", node.JSObject()); err == nil {
			if obj.Type() == js.TypeNumber {
				return obj.Int()
			}
		} else {

			n.Error = &err

		}
	}

	return 0

}

func (n *Node) Contains(node Node) bool {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("contains", node.JSObject()); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				n.Error = &baseobject.ErrObjectNotBool
			}
		} else {

			n.Error = &err

		}
	}

	return false
}

func (n Node) GetRootNode() Node {
	var err error
	var obj js.Value
	var newNode Node

	if n.NotError() {

		if obj, err = n.JSObject().CallWithErr("getRootNode"); err == nil {
			return NewFromJSObject(obj)
		}
		newNode.Error = &err
	}
	return newNode
}

func (n *Node) HasChildNodes() bool {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("hasChildNodes"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				n.Error = &baseobject.ErrObjectNotBool
			}
		} else {

			n.Error = &err

		}
	}

	return false

}

func (n Node) InsertBefore(elem, before Node) Node {
	var err error

	if n.NotError() {
		if _, err = n.JSObject().CallWithErr("insertBefore", elem, before); err == nil {
			return elem
		} else {
			n.Error = &err
		}
	}

	return n

}

func (n *Node) IsDefaultNamespace() bool {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("isDefaultNamespace"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				n.Error = &baseobject.ErrObjectNotBool
			}
		} else {

			n.Error = &err

		}
	}
	return false

}

func (n *Node) IsEqualNode() bool {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("isEqualNode"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				n.Error = &baseobject.ErrObjectNotBool
			}
		} else {

			n.Error = &err

		}
	}
	return false

}

func (n *Node) IsSameNode() bool {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("isSameNode"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			} else {
				n.Error = &baseobject.ErrObjectNotBool
			}
		} else {

			n.Error = &err

		}
	}
	return false

}

func (n *Node) LookupPrefix() string {
	var err error
	var obj js.Value
	if n.NotError() {
		if obj, err = n.JSObject().CallWithErr("lookupPrefix"); err == nil {
			if obj.Type() == js.TypeString {
				return obj.String()
			}
		} else {

			n.Error = &err

		}
	}
	return ""

}

func (n *Node) LookupNamespaceURI(prefix string) {
	var err error
	if n.NotError() {
		if _, err = n.JSObject().CallWithErr("lookupNamespaceURI", js.ValueOf(prefix)); err != nil {

			n.Error = &err

		}

	}

}

func (n *Node) Normalize() {
	var err error
	if n.NotError() {
		if _, err = n.JSObject().CallWithErr("normalize"); err != nil {

			n.Error = &err

		}

	}

}

func (n Node) RemoveChild(node Node) Node {
	var err error

	if n.NotError() {
		if _, err = n.JSObject().CallWithErr("removeChild", node); err == nil {
			return node
		} else {
			n.Error = &err
		}
	}

	return n

}

func (n Node) ReplaceChild(new, old Node) Node {
	var err error

	if n.NotError() {
		if _, err = n.JSObject().CallWithErr("replaceChild", new, old); err == nil {
			return old
		} else {
			n.Error = &err
		}
	}

	return n

}
