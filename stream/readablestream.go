package stream

import (
	"github.com/realPy/hogosuru/js"

	"github.com/realPy/hogosuru/object"
)

type ReadableStream struct {
	object.Object
}

func NewReadableStreamFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream
	if object.String(obj) == "[object ReadableStream]" {
		r.Object = r.SetObject(obj)
		return r, nil
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
