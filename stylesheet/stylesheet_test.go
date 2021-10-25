package stylesheet

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
