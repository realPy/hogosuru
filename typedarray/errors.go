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
	ErrNotAInt16Array                  = errors.New("Object is not a Int16Array")
	ErrNotImplementedInt16Array        = errors.New("Browser not implemented Int16Array")
	ErrNotAInt32Array                  = errors.New("Object is not a Int32Array")
	ErrNotImplementedInt32Array        = errors.New("Browser not implemented Int32Array")
	ErrNotAUint32Array                 = errors.New("Object is not a Uint32Array")
	ErrNotImplementedUint32Array       = errors.New("Browser not implemented Uint32Array")
	ErrNotAFloat32Array                = errors.New("Object is not a Float32Array")
	ErrNotImplementedFloat32Array      = errors.New("Browser not implemented Float32Array")
	ErrNotAFloat64Array                = errors.New("Object is not a Float64Array")
	ErrNotImplementedFloat64Array      = errors.New("Browser not implemented Float64Array")
)
