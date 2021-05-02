package object

import (
	"errors"
)

var (
	//ErrNotAnObject ErrNotAnObject error
	ErrNotAnObject = errors.New("The given value must be an object")
	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnMEv = errors.New("The given value must be an Message Event")
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Object")
)
