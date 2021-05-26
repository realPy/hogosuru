package formdata

// https://developer.mozilla.org/fr/docs/Web/API/FormData

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var formadatainterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//FormData struct
type FormData struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var formadatainstance JSInterface
		var err error
		if formadatainstance.objectInterface, err = js.Global().GetWithErr("FormData"); err == nil {
			formadatainterface = &formadatainstance
		}
	})

	return formadatainterface
}

func New() (FormData, error) {

	var formdata FormData

	if fci := GetJSInterface(); fci != nil {
		formdata.BaseObject = formdata.SetObject(fci.objectInterface.New())

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
