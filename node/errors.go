package node

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented Node")
	ErrNotANode            = errors.New("Object is not a Node")
	ErrNodeNoChilds        = errors.New("Node has no childs")
	ErrNodeNoParent        = errors.New("Node has no parent")
	ErrNodeNoParentElement = errors.New("Node has no parent element")
)
