package eventtarget

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`event=new EventTarget()
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if e, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object EventTarget]", e.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "event"); testingutils.AssertErr(t, err) {
		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object EventTarget]", event.ToString_())

		}
	}

}

func TestEvent(t *testing.T) {

	//var io chan bool = make(chan bool)
	var eventreceive bool = false

	if e, err := New(); testingutils.AssertErr(t, err) {

		if _, err := e.AddEventListener("test", func(e event.Event) {
			eventreceive = true
		}); testingutils.AssertErr(t, err) {

			ev, _ := event.New("test")
			e.DispatchEvent(ev)
			testingutils.AssertExpect(t, true, eventreceive)
		}

	}

}

func TestRemoveEventListener(t *testing.T) {

	if e, err := New(); testingutils.AssertErr(t, err) {

		if f, err := e.AddEventListener("test", func(e event.Event) {

		}); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, e.RemoveEventListener(f, "test"))

		}

	}

}
