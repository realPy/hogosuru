package datatransfer

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)`)
	m.Run()

}

func TestNew(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object DataTransfer]", d.ToString_())

	}
}

func TestFiles(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "dt"); testingutils.AssertErr(t, err) {

		if dt, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if files, err := dt.Files(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object FileList]", files.ToString_())
			}
		}
	}
}

func TestItems(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "dt"); testingutils.AssertErr(t, err) {

		if dt, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if items, err := dt.Items(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object DataTransferItemList]", items.ToString_())
			}
		}
	}
}

func TestTypes(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "dt"); testingutils.AssertErr(t, err) {

		if dt, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if types, err := dt.Types(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "Files", types.ToString_())
			}
		}
	}
}
