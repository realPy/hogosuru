package datatransfert

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/filelist"
)

var singleton sync.Once

var dtinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Channel struct
type DataTransfer struct {
	baseobject.BaseObject
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var dtinstance JSInterface
		var err error
		if dtinstance.objectInterface, err = js.Global().GetWithErr("DataTransfer"); err == nil {
			dtinterface = &dtinstance
		}
	})

	return dtinterface
}

//New Get a new channel broadcast
func New() (DataTransfer, error) {
	var dt DataTransfer

	if dti := GetJSInterface(); dti != nil {
		dt.BaseObject = dt.SetObject(dti.objectInterface.New())
		return dt, nil
	}
	return DataTransfer{}, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (DataTransfer, error) {
	var dt DataTransfer

	if di := GetJSInterface(); di != nil {
		if obj.InstanceOf(di.objectInterface) {
			dt.BaseObject = dt.SetObject(obj)
			return dt, nil
		}
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
