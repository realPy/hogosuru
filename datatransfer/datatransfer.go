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

//GetJSInterface get the JS interface
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
	var dti, obj js.Value
	var err error
	if dti = GetInterface(); dti.IsUndefined() {
		return dt, ErrNotImplemented
	}
	if obj, err = baseobject.New(dti); err != nil {
		return dt, err
	}
	dt.BaseObject = dt.SetObject(obj)
	return dt, nil
}

func NewFromJSObject(obj js.Value) (DataTransfer, error) {
	var dt DataTransfer
	var dti js.Value
	if dti = GetInterface(); dti.IsUndefined() {
		return dt, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return dt, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(dti) {
		return dt, ErrNotADataTransfer
	}
	dt.BaseObject = dt.SetObject(obj)
	return dt, nil
}

func (dt DataTransfer) Files() (filelist.FileList, error) {
	var err error
	var obj js.Value
	var f filelist.FileList
	if obj, err = dt.Get("files"); err == nil {
		return filelist.NewFromJSObject(obj)
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
		return datatranferitemlist.NewFromJSObject(obj)
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
		return array.NewFromJSObject(obj)
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
