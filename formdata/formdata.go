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
		if formadatainterface, err = js.Global().GetWithErr("FormData"); err != nil {
			formadatainterface = js.Null()
		}

	})

	return formadatainterface
}

func New() (FormData, error) {

	var formdata FormData

	if fci := GetInterface(); !fci.IsNull() {
		formdata.BaseObject = formdata.SetObject(fci.New())

		return formdata, nil
	}
	return formdata, ErrNotImplemented
}

func (f FormData) AppendString(key string, value string) error {
	var err error
	_, err = f.JSObject().CallWithErr("append", js.ValueOf(key), js.ValueOf(value))

	return err

}

func (f FormData) AppendJSObject(key string, object js.Value) error {
	var err error
	_, err = f.JSObject().CallWithErr("append", js.ValueOf(key), object)
	return err

}
