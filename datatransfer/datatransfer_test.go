package datatransfer

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

	if d, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object DataTransfer]", d.ToString_())

	}
}
