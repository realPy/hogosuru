package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

var singleton sync.Once

var readablestreaminterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if readablestreaminterface, err = js.Global().GetWithErr("ReadableStream"); err != nil {
			readablestreaminterface = js.Undefined()
		}
	})

	return readablestreaminterface
}

type ReadableStream struct {
	baseobject.BaseObject
}

type ReadableStreameFrom interface {
	ReadableStream_() ReadableStream
}

func (r ReadableStream) ReadableStream_() ReadableStream {
	return r
}

func (r ReadableStream) Locked() (bool, error) {
	return r.GetAttributeBool("locked")
}

func NewFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream

	if rsi := GetInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			r.BaseObject = r.SetObject(obj)
			return r, nil

		}
	}

	return r, ErrNotAReadableStream
}

func (r ReadableStream) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = r.JSObject().CallWithErr("cancel"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (r ReadableStream) GetReader() (ReadableStreamDefaultReader, error) {
	var err error
	var obj js.Value

	if obj, err = r.JSObject().CallWithErr("getReader"); err == nil {
		return NewReadableStreamDefaultReaderFromJSObject(obj)

	}
	return ReadableStreamDefaultReader{}, err

}
