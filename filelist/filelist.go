package filelist

// https://developer.mozilla.org/fr/docs/Web/API/FileList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/file"
)

var singleton sync.Once

var filelistinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//FileList struct
type FileList struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var filelistinstance JSInterface
		var err error
		if filelistinstance.objectInterface, err = js.Global().GetWithErr("FileList"); err == nil {
			filelistinterface = &filelistinstance
		}
	})

	return filelistinterface
}

func NewFromJSObject(obj js.Value) (FileList, error) {
	var f FileList

	if fli := GetJSInterface(); fli != nil {
		if obj.InstanceOf(fli.objectInterface) {
			f.BaseObject = f.SetObject(obj)
			return f, nil
		}
	}
	return f, ErrNotAnFileList
}

func (f FileList) Item(index int) (file.File, error) {

	return file.NewFromJSObject(f.JSObject().Index(index))

}
