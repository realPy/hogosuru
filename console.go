package hogosuru

import (
	"sync"

	"github.com/realPy/hogosuru/console"
)

var singletonConsole sync.Once
var globalconsole console.Console

func Console() console.Console {

	singletonConsole.Do(func() {
		var err error
		if globalconsole, err = console.New(); err != nil {
			panic(err)
		}

	})
	return globalconsole
}

func AssertErr(err error) bool {
	if err != nil {
		Console().Assert(true, err.Error())
	}

	return err == nil
}
