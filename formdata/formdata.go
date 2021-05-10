package formdata

import (
	"sync"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var formadatainterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//FormData struct
type FormData struct {
	object.Object
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
		formdata.Object = formdata.SetObject(fci.objectInterface.New())

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
