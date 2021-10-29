package file

// https://developer.mozilla.org/fr/docs/Web/API/File
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/date"
)

var singleton sync.Once

var fileinterface js.Value

//GetInterface get the  JS interface File
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if fileinterface, err = baseobject.Get(js.Global(), "File"); err != nil {
			fileinterface = js.Undefined()
		}
		blob.GetInterface()
		baseobject.Register(fileinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return fileinterface
}

type File struct {
	blob.Blob
}

type FileFrom interface {
	File_() File
}

func (f File) File_() File {
	return f
}

func New(bits interface{}, name string, value ...map[string]interface{}) (File, error) {

	var f File
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if objGo, ok := bits.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(bits))
	}

	arrayJS = append(arrayJS, js.ValueOf(name))
	if len(value) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(value[0]))
	}

	if fi := GetInterface(); !fi.IsUndefined() {

		if obj, err = baseobject.New(fi, arrayJS...); err == nil {
			f.BaseObject = f.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented

	}
	return f, err
}

func NewFromJSObject(obj js.Value) (File, error) {
	var f File
	var err error
	if fi := GetInterface(); !fi.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fi) {
				f.BaseObject = f.SetObject(obj)

			} else {
				err = ErrNotAFile
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return f, err
}

func (f File) Name() (string, error) {

	return f.GetAttributeString("name")
}

func (f File) Type() (string, error) {
	return f.GetAttributeString("type")
}

func (f File) LastModified() (int64, error) {
	return f.GetAttributeInt64("lastModified")
}

func (f File) LastModifiedDate() (date.Date, error) {
	var obj js.Value
	var d date.Date
	var err error
	if obj, err = f.Get("lastModifiedDate"); err == nil {
		d, err = date.NewFromJSObject(obj)
	}
	return d, err
}
