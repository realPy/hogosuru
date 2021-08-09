package htmlquoteelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLQuoteElement")
	ErrNotAnHTMLQuoteElement = errors.New("Object is not an HTMLQuoteElement")
)
