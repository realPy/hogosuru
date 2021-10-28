package datatransfer

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatranferitemlist"
	"github.com/realPy/hogosuru/filelist"
)

var singleton sync.Once

var dtinterface js.Value

//DataTransfer struct
type DataTransfer struct {
	baseobject.BaseObject
}

type DataTransferFrom interface {
	DataTransfer_() DataTransfer
}

func (d DataTransfer) DataTransfer_() DataTransfer {
	return d
}

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if dtinterface, err = baseobject.Get(js.Global(), "DataTransfer"); err != nil {
			dtinterface = js.Undefined()
		}
		baseobject.Register(dtinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return dtinterface
}

//New Get a new channel broadcast
func New() (DataTransfer, error) {
	var dt DataTransfer

	if dti := GetInterface(); !dti.IsUndefined() {
		dt.BaseObject = dt.SetObject(dti.New())
		return dt, nil
	}
	return DataTransfer{}, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (DataTransfer, error) {
	var dt DataTransfer
	var err error
	if dti := GetInterface(); !dti.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dti) {
				dt.BaseObject = dt.SetObject(obj)

			} else {
				err = ErrNotADataTransfer
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return dt, err
}

func (dt DataTransfer) Files() (filelist.FileList, error) {

	var err error
	var obj js.Value
	var f filelist.FileList

	if obj, err = dt.Get("files"); err == nil {

		f, err = filelist.NewFromJSObject(obj)
	}
	return f, err
}
func (dt DataTransfer) SetFiles(files filelist.FileList) error {

	return dt.SetAttribute("files", files)
}

func (dt DataTransfer) Items() (datatranferitemlist.DataTransferItemList, error) {

	var err error
	var obj js.Value
	var items datatranferitemlist.DataTransferItemList

	if obj, err = dt.Get("items"); err == nil {

		items, err = datatranferitemlist.NewFromJSObject(obj)
	}
	return items, err

}

func (dt DataTransfer) SetItems(list datatranferitemlist.DataTransferItemList) error {

	return dt.SetAttribute("items", list)
}
func (dt DataTransfer) Types() (array.Array, error) {

	var err error
	var obj js.Value
	var types array.Array

	if obj, err = dt.Get("types"); err == nil {

		types, err = array.NewFromJSObject(obj)
	}
	return types, err
}

func (dt DataTransfer) DropEffect() (string, error) {
	return dt.GetAttributeString("dropEffect")
}

func (dt DataTransfer) SetDropEffect(value string) error {
	return dt.SetAttributeString("dropEffect", value)
}

func (dt DataTransfer) EffectAllowed() (string, error) {
	return dt.GetAttributeString("effectAllowed")
}

func (dt DataTransfer) SetEffectAllowed(value string) error {
	return dt.SetAttributeString("effectAllowed", value)
}
