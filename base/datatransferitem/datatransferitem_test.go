package datatransferitem

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

// doc say that can accept add(string) but dont work return (Failed to execute 'add' on 'DataTransferItemList': parameter 1 is not of type 'File'.)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	item1=dt.items[0]`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj, err = baseobject.Get(js.Global(), "item1"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DataTransferItem]", d.ToString_())

		}
	}

}

func TestKind(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj, err = baseobject.Get(js.Global(), "item1"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if kind, err := d.Kind(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "file", kind)
			}

		}
	}

}

func TestType(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj, err = baseobject.Get(js.Global(), "item1"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if typed, err := d.Type(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "image/png", typed)
			}

		}
	}

}
func TestGetAsFile(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItem

	if obj, err = baseobject.Get(js.Global(), "item1"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if file, err := d.GetAsFile(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object File]", file.ToString_())
			}

		}
	}

}
