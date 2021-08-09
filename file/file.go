package file

// https://developer.mozilla.org/fr/docs/Web/API/File
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/blob"
)

var singleton sync.Once

var fileinterface js.Value

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if fileinterface, err = js.Global().GetWithErr("File"); err != nil {
			fileinterface = js.Null()
		}
	})
	baseobject.Register(fileinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return fileinterface
}

type File struct {
	blob.Blob
}

type FileFrom interface {
	File() File
}

func (f File) File() File {
	return f
}

func NewFromJSObject(obj js.Value) (File, error) {
	var f File

	if fi := GetInterface(); !fi.IsNull() {
		if obj.InstanceOf(fi) {
			f.BaseObject = f.SetObject(obj)
			return f, nil
		}
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
