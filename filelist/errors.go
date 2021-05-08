package filelist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented FileList")
	ErrNotAnFileList  = errors.New("Object is not a FileList")
)
