package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var readablestreaminterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var readablestreaminstance JSInterface
		var err error
		if readablestreaminstance.objectInterface, err = js.Global().GetWithErr("ReadableStream"); err == nil {
			readablestreaminterface = &readablestreaminstance
		}
	})

	return readablestreaminterface
}

type ReadableStream struct {
	baseobject.BaseObject
}

func NewReadableStreamFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream

	if rsi := GetJSInterface(); rsi != nil {
		if obj.InstanceOf(rsi.objectInterface) {
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
