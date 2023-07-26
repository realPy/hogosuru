package worker

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/messageevent"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`myWorker = new Worker('assets_test/script_test.js');`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "myWorker"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Worker", nav.ConstructName_())

		}
	}

}

func TestNew(t *testing.T) {
	var wchan chan string = make(chan string)

	if w, err := New("assets_test/script_test.js"); testingutils.AssertErr(t, err) {

		w.OnMessage(func(m messageevent.MessageEvent) {

			if d, err := m.Data(); testingutils.AssertErr(t, err) {

				if message, ok := d.(string); ok {
					wchan <- message
				}

			}

		})

		select {
		case message := <-wchan:
			testingutils.AssertExpect(t, message, "installok")
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("ServiceWorker request timeout")

		}

		w.PostMessage("test")

		select {
		case message := <-wchan:
			testingutils.AssertExpect(t, message, "testok")
		case <-time.After(time.Duration(20000) * time.Millisecond):
			t.Errorf("ServiceWorker request timeout")

		}

		testingutils.AssertErr(t, w.Terminate())

	}

}
