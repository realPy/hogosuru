package element

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented Node")
	ErrNotAnElement     = errors.New("Object is not an Element")
	ErrElementNoChilds  = errors.New("Element has no childs")
	ErrAttributeEmpty   = errors.New("Attribute is empty")
	ErrInsertAdjacent   = errors.New("Insert Adjacent failed")
	ErrElementNotFound  = errors.New("Element not Found")
	ErrElementsNotFound = errors.New("Elements not Found")
	ErrSendUnknownType  = errors.New("Unknown type send data provide to send method")
)
