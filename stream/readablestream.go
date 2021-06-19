package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var readablestreaminterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if readablestreaminterface, err = js.Global().GetWithErr("ReadableStream"); err != nil {
			readablestreaminterface = js.Null()
		}
	})

	return readablestreaminterface
}

type ReadableStream struct {
	baseobject.BaseObject
}

func NewReadableStreamFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream

	if rsi := GetInterface(); !rsi.IsNull() {
		if obj.InstanceOf(rsi) {
			r.BaseObject = r.SetObject(obj)
			return r, nil

		}
	}

	return r, ErrNotAReadableStream
}

func (r ReadableStream) GetReader() (ReadableStreamDefaultReader, error) {
	var err error
	var obj js.Value

	if obj, err = r.JSObject().CallWithErr("getReader"); err == nil {
		return NewReadableStreamDefaultReaderFromJSObject(obj)

	}
	return ReadableStreamDefaultReader{}, err

}
