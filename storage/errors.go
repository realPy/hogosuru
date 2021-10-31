package storage

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented Storage")
	ErrNotAStorage    = errors.New("Object is not a Storage object")
)
