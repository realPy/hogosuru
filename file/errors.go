package file

import "errors"

var (
	//ErrNotAnBlob ErrNotAnBlob error
	ErrNotAFile = errors.New("Object is not a File")
)
