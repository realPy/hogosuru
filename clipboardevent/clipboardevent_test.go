package clipboardevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`file1 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
	dt=new DataTransfer()
	dt.items.add(file1)
	event=new ClipboardEvent(dt)`)
	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "event"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ClipboardEvent]", d.ToString_())

		}
	}

}

func TestNew(t *testing.T) {

	if dt, err := datatransfer.New(); testingutils.AssertErr(t, err) {
		if d, err := New(dt); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ClipboardEvent]", d.ToString_())

		}
	}

}

/*
func TestClipboardData(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "event"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if data, err := d.ClipboardData(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object ClipboardEvent]", data.ToString_())
			}

		}
	}
}
*/
