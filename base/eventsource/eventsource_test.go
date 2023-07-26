//go:build localtest
// +build localtest

package eventsource

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/messageevent"
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

	if sse, err := New(sseurl); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object EventSource]", sse.ToString_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "sse"); testingutils.AssertErr(t, err) {
		if sse, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object EventSource]", sse.ToString_())

		}
	}

}

func TestOnError(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New("htt://error.com"); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object EventSource]", sse.ToString_())
		sse.SetOnError(func(e event.Event) {

			success <- false
		})

		select {
		case <-success:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No error receive")

		}

	}

}

func TestOnOpen(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New(sseurl); testingutils.AssertErr(t, err) {

		sse.SetOnOpen(func(e event.Event) {

			success <- true

		})

		select {
		case <-success:
		case <-time.After(time.Duration(10000) * time.Millisecond):
			t.Errorf("No open receive")

		}

	}

}

var str_expect []string = []string{"msg1", "msg2"}

func TestOnMessage(t *testing.T) {
	var success chan bool = make(chan bool)
	if sse, err := New(sseurl); testingutils.AssertErr(t, err) {
		var i int = 0
		sse.SetOnMessage(func(e messageevent.MessageEvent) {

			if str, err := e.Data(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, str, str_expect[i])
				i++
			}
			if i == 2 {
				success <- true
			}

		})

		select {
		case <-success:
		case <-time.After(time.Duration(10000) * time.Millisecond):
			t.Errorf("No open receive")

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "ReadyState", "resultattempt": 0},
	{"method": "Url", "resultattempt": "https://sse-echo.herokuapp.com/"},
	{"method": "WithCredentials", "resultattempt": false},
	{"method": "Close"},
}

func TestMethods(t *testing.T) {

	baseobject.Eval(`sse1=new EventSource("https://sse-echo.herokuapp.com")`)

	if obj, err := baseobject.Get(js.Global(), "sse1"); testingutils.AssertErr(t, err) {

		if anchor, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, anchor, result)
			}

		}

	}
}
