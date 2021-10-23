package uint8array

import "errors"

var (
	//ErrNotAUint8Array ErrNotAnBlob error
	ErrNotAUint8Array = errors.New("Object is not a Uint8Array")
	ErrNotImplemented = errors.New("Browser not implemented Uint8Array")
)
