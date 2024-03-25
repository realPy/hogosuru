package usb

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented USB")
	ErrNotAUSB        = errors.New("Object is not a USB")
)
