package file

// https://developer.mozilla.org/fr/docs/Web/API/File
import (
	"syscall/js"

	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/object"
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

func (f File) Name() string {
	var err error
	var obj js.Value

	if obj, err = f.JSObject().GetWithErr("name"); err == nil {

		return obj.String()
	}
	return ""
}

func (f File) Type() string {
	var err error
	var obj js.Value

	if obj, err = f.JSObject().GetWithErr("type"); err == nil {

		return obj.String()
	}
	return ""
}

func (f File) LastModified() string {
	var err error
	var obj js.Value

	if obj, err = f.JSObject().GetWithErr("lastModified"); err == nil {

		return obj.String()
	}
	return ""
}
