package datatransferitem

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/file"
)

var singleton sync.Once

var datatransferiteminterface js.Value

//DataTransferItem struct
type DataTransferItem struct {
	baseobject.BaseObject
}

type DataTransferItemFrom interface {
	DataTransferItem_() DataTransferItem
}

func (d DataTransferItem) DataTransferItem_() DataTransferItem {
	return d
}

//GetJSInterface get the JS interface DataTransferItemList
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if datatransferiteminterface, err = baseobject.Get(js.Global(), "DataTransferItem"); err != nil {
			datatransferiteminterface = js.Undefined()
		}
		baseobject.Register(datatransferiteminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return datatransferiteminterface
}

func NewFromJSObject(obj js.Value) (DataTransferItem, error) {
	var d DataTransferItem
	var err error
	if dti := GetInterface(); !dti.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if obj.InstanceOf(dti) {
				d.BaseObject = d.SetObject(obj)

			} else {
				err = ErrNotADataTransferItem
			}
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DataTransferItem) Kind() (string, error) {
	return d.GetAttributeString("kind")
}

func (d DataTransferItem) Type() (string, error) {
	return d.GetAttributeString("type")
}

func (d DataTransferItem) GetAsFile() (file.File, error) {

	var err error
	var obj js.Value
	var f file.File

	if obj, err = d.Call("getAsFile"); err == nil {

		f, err = file.NewFromJSObject(obj)
	}
	return f, err

}
