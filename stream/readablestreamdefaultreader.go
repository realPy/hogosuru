package stream

import (
	"io"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/uint8array"
)

var singletonReadableStreamDefault sync.Once

var readablestreamdefaultinterface js.Value

//GetReadStreamInterface
func GetReadStreamInterface() js.Value {

	singletonReadableStreamDefault.Do(func() {

		var err error
		if readablestreamdefaultinterface, err = js.Global().GetWithErr("ReadableStreamDefaultReader"); err != nil {
			readablestreamdefaultinterface = js.Undefined()
		}
	})

	return readablestreamdefaultinterface
}

type ReadableStreamDefaultReader struct {
	baseobject.BaseObject
}

type ReadableStreamDefaultReaderFrom interface {
	ReadableStreamDefaultReader_() ReadableStreamDefaultReader
}

func (r ReadableStreamDefaultReader) ReadableStreamDefaultReader_() ReadableStreamDefaultReader {
	return r
}

func NewReadableStreamDefaultReaderFromJSObject(obj js.Value) (ReadableStreamDefaultReader, error) {
	var r ReadableStreamDefaultReader

	if rsi := GetReadStreamInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			r.BaseObject = r.SetObject(obj)
			return r, nil

		}
	}

	return r, ErrNotAReadableStream
}

func (r ReadableStreamDefaultReader) Read(b []byte) (n int, err error) {

	var promiseread js.Value
	var p promise.Promise
	donechan := make(chan bool)
	err = nil

	if promiseread, err = r.JSObject().CallWithErr("read"); err == nil {

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			p.Then(func(i interface{}) *promise.Promise {
				if obj, ok := i.(baseobject.ObjectFrom); ok {
					if obj.JSObject().Get("done").Bool() == true {
						err = io.EOF
						donechan <- true
						return nil
					} else {
						var u8array uint8array.Uint8Array
						uint8arrayObject := obj.JSObject().Get("value")
						if u8array, err = uint8array.NewFromJSObject(uint8arrayObject); err == nil {
							n, err = u8array.CopyBytes(b)
						}

					}

				}

				donechan <- false
				return nil

			}, func(e error) {
				err = e
				donechan <- false
			})

		}
		<-donechan

	} else {
		err = io.ErrUnexpectedEOF
	}

	return
}

func (r ReadableStreamDefaultReader) asyncRead(preallocateBytes []byte, dataHandle func([]byte, error)) (n int, err error) {
	var promiseread js.Value
	var p promise.Promise
	err = nil

	if promiseread, err = r.JSObject().CallWithErr("read"); err == nil {

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			p.Then(func(i interface{}) *promise.Promise {
				var obj js.Value

				if b, ok := i.(baseobject.ObjectFrom); ok {
					obj = b.JSObject()
					if obj.Get("done").Bool() == true {
						err = io.EOF
						dataHandle(nil, err)
						return nil
					} else {
						var u8array uint8array.Uint8Array

						uint8arrayObject := obj.Get("value")

						if u8array, err = uint8array.NewFromJSObject(uint8arrayObject); err == nil {

							if _, err = u8array.CopyBytes(preallocateBytes); err == nil {
								dataHandle(preallocateBytes, err)
							}

						}

						p2, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
							_, err := r.asyncRead(preallocateBytes, dataHandle)
							return nil, err
						})
						return &p2

					}
				} else {
					dataHandle(nil, baseobject.ErrNotABaseObject)
				}

				return nil
			}, func(e error) {
				err = e
				dataHandle(nil, err)
			})

		}

	} else {
		err = io.ErrUnexpectedEOF
	}

	return

}

func (r ReadableStreamDefaultReader) AsyncRead(dataHandle func([]byte, error)) {
	//preallocate memory (dont loop with make its slow!)
	var data []byte = make([]byte, 2*1024*1024)
	r.asyncRead(data, dataHandle)

}
