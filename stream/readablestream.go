package stream

// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

var singleton sync.Once

var readablestreaminterface js.Value

//GetInterface get the JS interface ReadableStream
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if readablestreaminterface, err = baseobject.Get(js.Global(), "ReadableStream"); err != nil {
			readablestreaminterface = js.Undefined()
		}
		baseobject.Register(readablestreaminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return readablestreaminterface
}

type ReadableStream struct {
	baseobject.BaseObject
}

type ReadableStreamFrom interface {
	ReadableStream_() ReadableStream
}

func (r ReadableStream) ReadableStream_() ReadableStream {
	return r
}

func (r ReadableStream) Locked() (bool, error) {
	return r.GetAttributeBool("locked")
}

//New Create a new ReadableStream
func New() (ReadableStream, error) {
	var r ReadableStream
	var obj js.Value
	var err error
	if ri := GetInterface(); !ri.IsUndefined() {

		if obj, err = baseobject.New(ri); err == nil {
			r.BaseObject = r.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return r, err
}

func NewFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream
	var err error
	if rsi := GetInterface(); !rsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
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

func (r ReadableStream) Tee() ([]ReadableStream, error) {
	var err error
	var obj js.Value
	var ret []ReadableStream
	var a array.Array

	if obj, err = r.Call("tee"); err == nil {

		if a, err = array.NewFromJSObject(obj); err == nil {

			a.ForEach(func(i interface{}) {

				if r, ok := i.(ReadableStreamFrom); ok {
					ret = append(ret, r.ReadableStream_())
				}

			})
		}

	}
	return ret, err

}
