package file

import (
	"github.com/realPy/jswasm/js"

	"github.com/realPy/jswasm/blob"
	"github.com/realPy/jswasm/object"
)

type File struct {
	blob.Blob
}

func NewFromJSObject(obj js.Value) (File, error) {
	var f File
	if object.String(obj) == "[object File]" {
		f.Object = f.SetObject(obj)

		return f, nil
	}

	return f, ErrNotAFile
}
