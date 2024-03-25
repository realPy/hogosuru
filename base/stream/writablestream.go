package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/promise"
)

var singletonw sync.Once

var writablestreaminterface js.Value

// GetWInterface get the JS interface WritableStream.
func GetWInterface() js.Value {

	singletonw.Do(func() {

		var err error
		if writablestreaminterface, err = baseobject.Get(js.Global(), "WritableStream"); err != nil {
			writablestreaminterface = js.Undefined()
		}
		baseobject.Register(writablestreaminterface, func(v js.Value) (interface{}, error) {
			return NewWriteableStreamFromJSObject(v)
		})
	})

	return writablestreaminterface
}

type WritableStream struct {
	baseobject.BaseObject
}

type WritableStreamFrom interface {
	WritableStream_() WritableStream
}

func (r WritableStream) WritableStream_() WritableStream {
	return r
}

// NewWritableStream Create a new NewWritableStream
func NewWritableStream() (WritableStream, error) {
	var w WritableStream
	var obj js.Value
	var err error
	if wi := GetWInterface(); !wi.IsUndefined() {

		if obj, err = baseobject.New(wi); err == nil {
			w.BaseObject = w.SetObject(obj)
		}

	} else {
		err = ErrNotImplementedWritableStream
	}
	return w, err
}

func NewWriteableStreamFromJSObject(obj js.Value) (WritableStream, error) {
	var w WritableStream
	var err error
	if wsi := GetWInterface(); !wsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(wsi) {
				w.BaseObject = w.SetObject(obj)

			} else {
				err = ErrNotAWritableStream
			}
		}
	} else {
		err = ErrNotImplementedWritableStream
	}

	return w, err
}

func (w WritableStream) Locked() (bool, error) {
	return w.GetAttributeBool("locked")
}

func (w WritableStream) Abort(reason string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("abort", js.ValueOf(reason)); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStream) Close() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("close"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStream) GetWriter() (WritableStreamDefaultWriter, error) {
	var err error
	var obj js.Value

	if obj, err = w.Call("getWriter"); err == nil {
		return NewWritableStreamDefaultWriterFromJSObject(obj)

	}
	return WritableStreamDefaultWriter{}, err
}
