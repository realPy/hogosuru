package htmltitleelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLTitleElement")
	ErrNotAnHTMLTitleElement = errors.New("Object is not an HTMLTitleElement")
)
