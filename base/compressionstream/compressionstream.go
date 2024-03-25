package compressionstream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/response"
	"github.com/realPy/hogosuru/base/stream"
)

func init() {

	initinterface.RegisterInterface(GetInterface)

}

var singleton sync.Once

var compressionstreaminterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if compressionstreaminterface, err = baseobject.Get(js.Global(), "CompressionStream"); err != nil {
			compressionstreaminterface = js.Undefined()
		}

		response.GetInterface()
		baseobject.Register(compressionstreaminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return compressionstreaminterface
}

// CompressionStream struct
type CompressionStream struct {
	baseobject.BaseObject
}

type CompressionStreamFrom interface {
	CompressionStream_() CompressionStream
}

func (c CompressionStream) CompressionStream() CompressionStream {
	return c
}

func New(format string) (CompressionStream, error) {
	var c CompressionStream
	var err error
	var obj js.Value

	if compressionstreami := GetInterface(); !compressionstreami.IsUndefined() {

		if obj, err = baseobject.New(compressionstreami, js.ValueOf(format)); err == nil {
			c.BaseObject = c.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return c, err

}

func NewFromJSObject(obj js.Value) (CompressionStream, error) {
	var c CompressionStream
	var err error
	if compressionstreami := GetInterface(); !compressionstreami.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(compressionstreami) {

				c.BaseObject = c.SetObject(obj)
			} else {
				err = ErrNotACompressionStream
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return c, err
}

func (c CompressionStream) Readable() (stream.ReadableStream, error) {
	var err error
	var obj js.Value

	if obj, err = c.Get("readable"); err == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return stream.ReadableStream{}, err
}

func (c CompressionStream) Writable() (stream.WritableStream, error) {
	var err error
	var obj js.Value

	if obj, err = c.Get("writable"); err == nil {
		return stream.NewWriteableStreamFromJSObject(obj)

	}
	return stream.WritableStream{}, err
}
