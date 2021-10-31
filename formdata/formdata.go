package formdata

// https://developer.mozilla.org/fr/docs/Web/API/FormData

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var formadatainterface js.Value

//FormData struct
type FormData struct {
	baseobject.BaseObject
}

type FormDataFrom interface {
	FormData_() FormData
}

func (f FormData) FormData_() FormData {
	return f
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if formadatainterface, err = baseobject.Get(js.Global(), "FormData"); err != nil {
			formadatainterface = js.Undefined()
		}

	})

	return formadatainterface
}

func New() (FormData, error) {

	var formdata FormData
	var obj js.Value
	var err error
	if fci := GetInterface(); !fci.IsUndefined() {

		if obj, err = baseobject.New(fci); err == nil {
			formdata.BaseObject = formdata.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return formdata, err
}

func (f FormData) AppendString(key string, value string) error {
	var err error
	_, err = f.Call("append", js.ValueOf(key), js.ValueOf(value))

	return err

}

func (f FormData) AppendJSObject(key string, object js.Value) error {
	var err error
	_, err = f.Call("append", js.ValueOf(key), object)
	return err

}
