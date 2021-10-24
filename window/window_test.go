package window

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestWindow(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if d, err := w.Document(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLDocument]", d.ToString_())
		}

	}

}

func TestHistory(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if h, err := w.History(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object History]", h.ToString_())
		}

	}

}
