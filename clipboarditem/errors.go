package clipboarditem

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented ClipboardItem")
	ErrNotAClipboard  = errors.New("Object is not a ClipboardItem")
)
