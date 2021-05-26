package stream

import (
	"fmt"
	"io"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/uint8array"
)

type ReadableStreamDefaultReader struct {
	baseobject.BaseObject
}

func NewReadableStreamDefaultReaderFromJSObject(obj js.Value) (ReadableStreamDefaultReader, error) {
	var r ReadableStreamDefaultReader
	if baseobject.String(obj) == "[object ReadableStreamDefaultReader]" {
		r.BaseObject = r.SetObject(obj)
		return r, nil
	}

	return r, ErrNotAReadableStream
}

func (r ReadableStreamDefaultReader) read() {

	var err error
	var promiseread js.Value

	if promiseread, err = r.JSObject().CallWithErr("read"); err == nil {
		var then js.Func

		then = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			js.Global().Set("debug", args[0])

			if args[0].Get("done").Bool() == true {
				return nil
			} else {
				fmt.Printf("ici %d\n", args[0].Get("value").Length())
			}
			//fmt.Printf("%s %s\n", baseobject.String(args[0]), baseobject.String(args[1]))

			r.read()
			return nil
		})
		promiseread.Call("then", then)

	}

}

func (r ReadableStreamDefaultReader) Read(b []byte) (n int, err error) {

	var promiseread js.Value
	donechan := make(chan bool)
	err = nil

	if promiseread, err = r.JSObject().CallWithErr("read"); err == nil {
		var then js.Func

		then = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if args[0].Get("done").Bool() == true {
				donechan <- true
				err = io.EOF
				return nil
			} else {
				var u8array uint8array.Uint8Array
				uint8arrayObject := args[0].Get("value")
				if u8array, err = uint8array.NewFromJSObject(uint8arrayObject); err == nil {
					n, err = u8array.CopyBytes(b)
				}

			}
			donechan <- false
			return nil
		})

		promiseread.Call("then", then)
		<-donechan

	} else {
		err = io.ErrUnexpectedEOF
	}

	return
}
