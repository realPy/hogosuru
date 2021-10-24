package domexception

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestException(t *testing.T) {

	if e, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "Error", e.ToString_())
	}

	var message string = "message error"

	if e, err := New(message); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "Error: message error", e.ToString_())

		if str, err := e.Message(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, message, str)
		}
	}

	var customname string = "custom name"
	if e, err := New(message, customname); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "custom name: message error", e.ToString_())

		if str, err := e.Name(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, customname, str)
		}
	}

}
