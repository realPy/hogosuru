package file

// https://developer.mozilla.org/fr/docs/Web/API/File
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/blob"
)

var singleton sync.Once

var fileinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var fileinstance JSInterface
		var err error
		if fileinstance.objectInterface, err = js.Global().GetWithErr("File"); err == nil {
			fileinterface = &fileinstance
		}
	})

	return fileinterface
}

type File struct {
	blob.Blob
}

func NewFromJSObject(obj js.Value) (File, error) {
	var f File

	if fi := GetJSInterface(); fi != nil {
		if obj.InstanceOf(fi.objectInterface) {
			f.Object = f.SetObject(obj)
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
