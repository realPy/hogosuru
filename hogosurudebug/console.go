package hogosurudebug

import (
	"syscall/js"

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
