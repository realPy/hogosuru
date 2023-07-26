package filelist

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	file2 = new File(['(⌐□_□)'], 'chucknorris2.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	dt.items.add(file2)
	inputfiles=dt.files`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj, err = baseobject.Get(js.Global(), "inputfiles"); testingutils.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object FileList]", f.ToString_())

		}
	}
}

func TestLength(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj, err = baseobject.Get(js.Global(), "inputfiles"); testingutils.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if l, err := f.Length(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, 2, l)
			}

		}
	}
}

func TestItem(t *testing.T) {
	var err error
	var obj js.Value
	var f FileList

	if obj, err = baseobject.Get(js.Global(), "inputfiles"); testingutils.AssertErr(t, err) {

		if f, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if file, err := f.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object File]", file.ToString_())
			}

		}
	}
}
