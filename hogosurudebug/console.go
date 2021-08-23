package hogosurudebug

import (
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

		if f, err = c.JSObject().GetWithErr(typehandler); err == nil {
			var defaultType js.Value
			if defaultType, err = f.CallWithErr("bind", c.JSObject()); err == nil {

				if err = c.JSObject().SetWithErr("default"+typehandler, defaultType); err == nil {

					handlerFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

						defaultType.Call("apply", c.JSObject(), js.ValueOf([]interface{}{args[0]}))
						handler(args[0].String())
						return nil
					})

					c.JSObject().SetWithErr(typehandler, handlerFunc)

				}

			}
		}
	}

	return err
}

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

func EnableDebug() {
	baseobject.SetConsoleDebug(Console())
}

func AssertErr(err error) bool {
	if err != nil {
		Console().Assert(err == nil, err.Error())
	}

	return err == nil
}

func AssertDebug(err error) bool {
	if err != nil {
		Console().BaseObject.Debug(err.Error())
	}

	return err == nil
}
