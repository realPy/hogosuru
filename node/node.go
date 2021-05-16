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

func (n *Node) BaseURI() string {
	var nodeObject js.Value
	var err error

	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("baseURI"); err == nil {
			return nodeObject.String()
		} else {
			n.Error = &err
		}
	}

	return ""
}

func (n Node) FirstChild() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

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

func (n *Node) IsConnected() bool {

	var err error
	var obj js.Value

	if obj, err = n.JSObject().GetWithErr("isConnected"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool()
		}
	} else {
		n.Error = &err
	}

	return false
}

func (n Node) LastChild() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("lastChild"); err == nil {

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

func (n Node) NextSibling() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("nextSibling"); err == nil {

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

func (n *Node) NodeValue() Node {
	var err error
	var obj js.Value
	var newNode Node

	if n.Error != nil {
		return *n
	}

	if obj, err = n.JSObject().GetWithErr("nodeValue"); err == nil {
		newNode = NewFromJSObject(obj)

	} else {
		newNode.Error = &err
	}

	return newNode
}

func (n *Node) SetNodeValue(nset Node) Node {

	if n.Error != nil {
		return *n
	}

	if err := n.JSObject().SetWithErr("nodeValue", nset.JSObject()); err != nil {
		n.Error = &err
	}

	return *n
}

func (n Node) OwnerDocument() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("ownerDocument"); err == nil {

			if nodeObject.IsNull() {
				err = ErrNodeNoParent

			} else {

				newNode = NewFromJSObject(nodeObject)

			}

		} else {
			newNode.Error = &err
		}

	}

	return newNode
}

func (n Node) ParentNode() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("parentNode"); err == nil {

			if nodeObject.IsNull() {
				err = ErrNodeNoParent

			} else {

				newNode = NewFromJSObject(nodeObject)

			}

		} else {
			newNode.Error = &err
		}

	}

	return newNode
}

func (n Node) ParentElement() Node {

	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("parentElement"); err == nil {

			if nodeObject.IsNull() {
				err = ErrNodeNoParentElement

			} else {

				newNode = NewFromJSObject(nodeObject)

			}

		} else {
			newNode.Error = &err
		}

	}

	return newNode

}

func (n Node) PreviousSibling() Node {
	var nodeObject js.Value
	var newNode Node
	var err error

	if n.Error != nil {
		return n
	}

	newNode.Error = n.Error
	if n.Error == nil {
		if nodeObject, err = n.JSObject().GetWithErr("previousSibling"); err == nil {

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

func (n *Node) TextContent() string {

	var err error
	var obj js.Value

	if n.Error == nil {
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
	if n.Error == nil {
		if err = n.JSObject().SetWithErr("textContent", js.ValueOf(content)); err != nil {

			n.Error = &err
		}
	}
	return *n
}

func (n *Node) AppendChild(add Node) {
	var err error

	if n.Error == nil {
		if _, err = n.JSObject().CallWithErr("appendChild", add.JSObject()); err != nil {
			n.Error = &err
		}

	}

}

func (n Node) CloneNode(deep bool) Node {
	var err error
	var obj js.Value
	var newNode Node

	if n.Error == nil {

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
	if n.Error == nil {
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
	if n.Error == nil {
		if obj, err = n.JSObject().CallWithErr("contains", node.JSObject()); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
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

	if n.Error == nil {

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
	if n.Error == nil {
		if obj, err = n.JSObject().CallWithErr("hasChildNodes"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
			}
		} else {

			n.Error = &err

		}
	}

	return false

}

func (n Node) InsertBefore(elem, before Node) Node {
	var err error

	if n.Error == nil {
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
	if n.Error == nil {
		if obj, err = n.JSObject().CallWithErr("isDefaultNamespace"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
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
	if n.Error == nil {
		if obj, err = n.JSObject().CallWithErr("isEqualNode"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
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
	if n.Error == nil {
		if obj, err = n.JSObject().CallWithErr("isSameNode"); err == nil {
			if obj.Type() == js.TypeBoolean {
				return obj.Bool()
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
	if n.Error == nil {
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
	if n.Error == nil {
		if _, err = n.JSObject().CallWithErr("lookupNamespaceURI", js.ValueOf(prefix)); err != nil {

			n.Error = &err

		}

	}

}

func (n *Node) Normalize() {
	var err error
	if n.Error == nil {
		if _, err = n.JSObject().CallWithErr("normalize"); err != nil {

			n.Error = &err

		}

	}

}

func (n Node) RemoveChild(node Node) Node {
	var err error

	if n.Error == nil {
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

	if n.Error == nil {
		if _, err = n.JSObject().CallWithErr("replaceChild", new, old); err == nil {
			return old
		} else {
			n.Error = &err
		}
	}

	return n

}
