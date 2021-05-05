package fetch

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Fetch")
	//ErrNotAnFResp ErrNotAnFResp error
	ErrNotAnFResp = errors.New("The given value must be an fetch response")
)
