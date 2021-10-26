package jserror

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if e, err := New("throw error"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "Error: throw error", e.ToString_())

	}

	var message string = "message error"

	if e, err := New(message); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "Error: message error", e.ToString_())

		if str, err := e.Message(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, message, str)
		}
		message = "message error2"
		e.SetMessage(message)

		testingutils.AssertExpect(t, "Error: message error2", e.ToString_())

	}

	var customname string = "custom name"
	if e, err := New(message); testingutils.AssertErr(t, err) {
		e.SetName(customname)
		testingutils.AssertExpect(t, "custom name: message error2", e.ToString_())

		if str, err := e.Name(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, customname, str)
		}
	}

}
