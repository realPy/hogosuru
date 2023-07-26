package object

import "errors"

var (
	ErrNotAnObject    = errors.New("The given value must be a Object")
	ErrNotImplemented = errors.New("Browser not implemented Object")
)
