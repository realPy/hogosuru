package hogosuru

import (
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {
	baseobject.SetSyscall()

}

func Init() {

	initinterface.Init()
}
