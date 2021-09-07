package objectmap

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented Mapy")
	ErrNotAMap        = errors.New("The given value must be a Map")
)
