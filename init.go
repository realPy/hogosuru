package hogosuru

import (
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {
	baseobject.SetSyscall()
	initinterface.Init()
}
