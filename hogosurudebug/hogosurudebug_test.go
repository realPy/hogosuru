package hogosurudebug

import (
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
