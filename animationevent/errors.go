package animationevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented AnimationEvent")
	ErrNotAnAnimationEvent = errors.New("Object is not an AnimationEvent")
)
