package eventsource

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

// use sse echo test from heroku

var sseurl string = "https://sse-echo.herokuapp.com/events?delay=3&data=data%3A%20msg1%0A%0Adata%3A%20msg2%0A%0Aevent%3A%20close%0Adata%3A%20msgend%0A%0A"

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`sse=new EventSource("https://sse-echo.herokuapp.com")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if ws, err := New(sseurl); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object EventSource]", ws.ToString_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "sse"); testingutils.AssertErr(t, err) {
		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object EventSource]", event.ToString_())

		}
	}

}
