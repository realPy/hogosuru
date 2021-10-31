package filelist

// https://developer.mozilla.org/fr/docs/Web/API/FileList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/file"
)

var singleton sync.Once

var filelistinterface js.Value

//FileList struct
type FileList struct {
	baseobject.BaseObject
}

type FileListFrom interface {
	FileList_() FileList
}

func (f FileList) FileList_() FileList {
	return f
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if filelistinterface, err = baseobject.Get(js.Global(), "FileList"); err != nil {
			filelistinterface = js.Undefined()
		}
		baseobject.Register(filelistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return filelistinterface
}

func NewFromJSObject(obj js.Value) (FileList, error) {
	var f FileList
	var err error
	if fli := GetInterface(); !fli.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fli) {
				f.BaseObject = f.SetObject(obj)

			} else {
				err = ErrNotAnFileList
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return f, err
}

func (f FileList) Item(index int) (file.File, error) {

	return file.NewFromJSObject(f.JSObject().Index(index))

}
func (f FileList) Length() (int, error) {

	return f.GetAttributeInt("length")

}
