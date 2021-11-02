package node

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
)

var singleton sync.Once

var nodeinterface js.Value

//GetInterface
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if nodeinterface, err = baseobject.Get(js.Global(), "Node"); err != nil {
			nodeinterface = js.Undefined()
		}
		baseobject.Register(nodeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return nodeinterface
}

//we use here aenw method of chaining method

type Node struct {
	eventtarget.EventTarget
}

type NodeFrom interface {
	Node_() Node
}

func (n Node) Node_() Node {
	return n
}

func NewFromJSObject(obj js.Value) (Node, error) {
	var n Node
	var err error

	if ni := GetInterface(); !ni.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ni) {
				n.BaseObject = n.SetObject(obj)

			} else {
				err = ErrNotANode
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return n, err
}

func (n Node) getAttributeNode(attribute string) (Node, error) {
	var nodeObject js.Value
	var newNode Node
	var err error

	if nodeObject, err = n.Get(attribute); err == nil {

		if nodeObject.IsUndefined() {
			err = ErrNodeNoChilds

		} else {

			newNode, err = NewFromJSObject(nodeObject)

		}

	}

	return newNode, err
}

func (n Node) BaseURI() (string, error) {

	return n.GetAttributeString("baseURI")
}

func (n Node) FirstChild() (Node, error) {

	return n.getAttributeNode("firstChild")
}

func (n Node) IsConnected() (bool, error) {

	return n.GetAttributeBool("isConnected")
}

func (n Node) LastChild() (Node, error) {
	return n.getAttributeNode("lastChild")
}

func (n Node) NextSibling() (Node, error) {
	return n.getAttributeNode("nextSibling")
}

func (n Node) NodeName() (string, error) {

	return n.GetAttributeString("nodeName")

}

func (n Node) NodeType() (int, error) {
	return n.GetAttributeInt("nodeType")
}

func (n Node) NodeValue() (Node, error) {
	return n.getAttributeNode("nodeValue")
}

func (n Node) SetNodeValue(nset Node) error {

	return n.Set("nodeValue", nset.JSObject())
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

	return n.GetAttributeString("textContent")
}

func (n Node) SetTextContent(content string) error {

	return n.Set("textContent", js.ValueOf(content))
}

func (n Node) AppendChild(add Node) error {

	_, err := n.Call("appendChild", add.JSObject())
	return err
}

func (n Node) CloneNode(deep bool) (Node, error) {
	var err error
	var obj js.Value
	var newNode Node

	if obj, err = n.Call("cloneNode", js.ValueOf(deep)); err == nil {
		return NewFromJSObject(obj)
	}

	return newNode, err
}

func (n Node) CompareDocumentPosition(node Node) (int, error) {
	var err error
	var obj js.Value
	var result int

	if obj, err = n.Call("compareDocumentPosition", node.JSObject()); err == nil {
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
	if obj, err = n.Call("contains", node.JSObject()); err == nil {
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

	if obj, err = n.Call("getRootNode"); err == nil {
		newNode, err = NewFromJSObject(obj)
	}
	return newNode, err
}

func (n Node) HasChildNodes() (bool, error) {
	return n.CallBool("hasChildNodes")
}

func (n Node) InsertBefore(elem, before Node) (Node, error) {
	var err error

	_, err = n.Call("insertBefore", elem.JSObject(), before.JSObject())

	return elem, err

}

func (n *Node) IsDefaultNamespace() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.Call("isDefaultNamespace"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n *Node) IsEqualNode(n1 Node) (bool, error) {

	var err error
	var obj js.Value
	var result bool

	if obj, err = n.Call("isEqualNode", n1.JSObject()); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (n *Node) IsSameNode(n1 Node) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = n.Call("isSameNode", n1.JSObject()); err == nil {
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

	if obj, err = n.Call("lookupPrefix"); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		}
	}

	return result, err

}

func (n *Node) LookupNamespaceURI(prefix string) error {
	var err error
	_, err = n.Call("lookupNamespaceURI", js.ValueOf(prefix))
	return err
}

func (n *Node) Normalize() error {
	var err error
	_, err = n.Call("normalize")
	return err
}

func (n Node) RemoveChild(node Node) (Node, error) {
	var err error
	_, err = n.Call("removeChild", node.JSObject())
	return node, err

}

func (n Node) ReplaceChild(new, old Node) (Node, error) {
	var err error

	_, err = n.Call("replaceChild", new.JSObject(), old.JSObject())

	return old, err

}
