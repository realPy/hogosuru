package hogosuru

import (
	"github.com/realPy/hogosuru/hogosurudebug"
)

func AssertErr(err error) bool {

	return hogosurudebug.AssertErr(err)
}
