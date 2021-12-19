package datatranferitemlist

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatransferitem"
	"github.com/realPy/hogosuru/file"
)

var singleton sync.Once

var datatransferitemlistinterface js.Value

//DataTransferItemList struct
type DataTransferItemList struct {
	baseobject.BaseObject
}

type DataTransferItemListFrom interface {
	DataTransferItemList_() DataTransferItemList
}

func (d DataTransferItemList) DataTransferItemList_() DataTransferItemList {
	return d
}

//GetJSInterface get the JS interface DataTransferItemList
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if datatransferitemlistinterface, err = baseobject.Get(js.Global(), "DataTransferItemList"); err != nil {
			datatransferitemlistinterface = js.Undefined()
		}
		baseobject.Register(datatransferitemlistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return datatransferitemlistinterface
}

func NewFromJSObject(obj js.Value) (DataTransferItemList, error) {
	var d DataTransferItemList
	var dli js.Value
	if dli = GetInterface(); dli.IsUndefined() {
		return d, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return d, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(dli) {
		return d, ErrNotADataTransferItemList
	}
	d.BaseObject = d.SetObject(obj)
	return d, nil
}

func (d DataTransferItemList) Length() (int, error) {
	return d.GetAttributeInt("length")
}

// doc said input can be file or string but string not work
func (d DataTransferItemList) Add(f file.File) error {
	var err error
	_, err = d.Call("add", f.JSObject())
	return err
}

func (d DataTransferItemList) Remove(index int) error {
	var err error
	_, err = d.Call("remove", js.ValueOf(index))
	return err
}

func (d DataTransferItemList) Clear() error {
	var err error
	_, err = d.Call("clear")
	return err
}

//this func doesn't work but exist in doc
func (d DataTransferItemList) DataTransferItem(index int) (datatransferitem.DataTransferItem, error) {
	var err error
	var obj js.Value
	var dt datatransferitem.DataTransferItem
	if obj, err = d.GetIndex(index); err != nil {
		return dt, err
	}
	return datatransferitem.NewFromJSObject(obj)

}
