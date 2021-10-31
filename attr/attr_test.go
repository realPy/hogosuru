package attr

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

//Attr can't be contructed with New. Test will be done in document
