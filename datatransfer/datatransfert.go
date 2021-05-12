package datatransfert

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/filelist"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var dtinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Channel struct
type DataTransfer struct {
	object.Object
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var dtinstance JSInterface
		var err error
		if dtinstance.objectInterface, err = js.Global().GetWithErr("DataTransfer"); err == nil {
			dtinterface = &dtinstance
			if object.String(dtinstance.objectInterface) == "" {
				dtinterface = nil
			}
		}
	})

	return dtinterface
}

//New Get a new channel broadcast
func New() (DataTransfer, error) {
	var dt DataTransfer

	if dti := GetJSInterface(); dti != nil {
		dt.Object = dt.SetObject(dti.objectInterface.New())
		return dt, nil
	}
	return DataTransfer{}, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (DataTransfer, error) {
	var dt DataTransfer

	if object.String(obj) == "[object DataTransfer]" {
		dt.Object = dt.SetObject(obj)
		return dt, nil
	}

	return dt, ErrNotADataTransfert
}

func (dt DataTransfer) Files() (filelist.FileList, error) {

	var err error
	var obj js.Value

	if obj, err = dt.JSObject().GetWithErr("files"); err == nil {

		return filelist.NewFromJSObject(obj)
	}
	return filelist.FileList{}, err

}
