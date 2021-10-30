package hogosurudebug

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/console"
)

/***********************************************
For debug
************************************************/

func InstallConsoleHandler(typehandler string, handler func(string)) error {
	var err error
	var c console.Console

	if c, err = console.New(); err == nil {
		var f js.Value

		if f, err = c.Get(typehandler); err == nil {
			var defaultType js.Value

			if defaultType, err = baseobject.Call(f, "bind", c.JSObject()); err == nil {

				if err = c.Set("default"+typehandler, defaultType); err == nil {

					handlerFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

						defaultType.Call("apply", c.JSObject(), js.ValueOf([]interface{}{args[0]}))
						handler(args[0].String())
						return nil
					})

					c.Set(typehandler, handlerFunc)

				}

			}
		}
	}

	return err
}

var singletonConsole sync.Once
var globalconsole console.Console
var AssertErr func(err error) bool = assertErr_wstrace

func Console() console.Console {

	singletonConsole.Do(func() {
		var err error
		if globalconsole, err = console.New(); err != nil {
			panic(err)
		}

	})
	return globalconsole
}

func EnableDebug() {
	baseobject.SetConsoleDebug(Console())
	AssertErr = assertErr_strace
}

func assertErr_wstrace(err error) bool {
	if err != nil {
		Console().Assert(err == nil, err.Error())
	}

	return err == nil
}

func assertErr_strace(err error) bool {

	if err != nil {
		_, file, line, _ := runtime.Caller(2)
		strerr := fmt.Sprintf("%s:%d >> %s", filepath.Base(file), line, err.Error())
		Console().Assert(err == nil, strerr)
	}

	return err == nil
}

func AssertDebug(err error) bool {
	if err != nil {
		Console().BaseObject.Debug(err.Error())
	}

	return err == nil
}
