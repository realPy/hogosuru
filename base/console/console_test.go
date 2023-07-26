package console

import (
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

var expectMethods []string = []string{"assert", "clear", "count", "countReset", "debug",
	"dir", "dirxml", "error", "group", "groupCollapsed", "groupEnd", "info", "log", "time", "timeEnd", "timeLog",
	"trace", "warn"}

// infortunately it's not possible to acces js console , we just verify
func TestNew(t *testing.T) {

	if c, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object console]", c.ToString_())

		testingutils.ImplementedExpect(t, c.BaseObject, expectMethods)
	}
}
