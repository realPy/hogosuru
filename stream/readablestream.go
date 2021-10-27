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
		if readablestreaminterface, err = baseobject.Get(js.Global(), "ReadableStream"); err != nil {
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
	var err error
	if rsi := GetInterface(); !rsi.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(rsi) {
				r.BaseObject = r.SetObject(obj)

			} else {
				err = ErrNotAReadableStream
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return r, err
}

func (r ReadableStream) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = r.Call("cancel"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (r ReadableStream) GetReader() (ReadableStreamDefaultReader, error) {
	var err error
	var obj js.Value

	if obj, err = r.Call("getReader"); err == nil {
		return NewReadableStreamDefaultReaderFromJSObject(obj)

	}
	return ReadableStreamDefaultReader{}, err

}
