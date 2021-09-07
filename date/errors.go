package date

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented Date")
	ErrNotADate       = errors.New("The given value must be date")
)
