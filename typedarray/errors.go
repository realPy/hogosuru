package typedarray

import "errors"

var (
	//ErrNotAUint8Array ErrNotAnBlob error
	ErrNotAUint8Array                  = errors.New("Object is not a Uint8Array")
	ErrNotImplementedUint8Array        = errors.New("Browser not implemented Uint8Array")
	ErrNotAUint8ClampedArray           = errors.New("Object is not a Uint8ClampedArray")
	ErrNotImplementedUint8ClampedArray = errors.New("Browser not implemented Uint8ClampedArray")
	ErrNotAInt8Array                   = errors.New("Object is not a Int8Array")
	ErrNotImplementedInt8Array         = errors.New("Browser not implemented Int8Array")
	ErrNotAUint16Array                 = errors.New("Object is not a Uint16Array")
	ErrNotImplementedUint16Array       = errors.New("Browser not implemented Uint16Array")
)
