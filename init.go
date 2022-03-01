package hogosuru

import (
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
)

func Init() {
	baseobject.SetSyscall()
	initinterface.Init()
}
