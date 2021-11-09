package indexeddb

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`iddb=window.indexedDB
	`)
	m.Run()
}
