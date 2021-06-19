package datatransfert

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/filelist"
)

var singleton sync.Once

var dtinterface js.Value

//Channel struct
type DataTransfer struct {
	baseobject.BaseObject
}

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if dtinterface, err = js.Global().GetWithErr("DataTransfer"); err != nil {
			dtinterface = js.Null()
		}
	})

	baseobject.Register(dtinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return dtinterface
}

//New Get a new channel broadcast
func New() (DataTransfer, error) {
	var dt DataTransfer

	if dti := GetInterface(); !dti.IsNull() {
		dt.BaseObject = dt.SetObject(dti.New())
		return dt, nil
	}
	return DataTransfer{}, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (DataTransfer, error) {
	var dt DataTransfer

	if dti := GetInterface(); !dti.IsNull() {
		if obj.InstanceOf(dti) {
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
