package stream

import (
	"io"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
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

func (r ReadableStreamDefaultReader) Read(b []byte) (n int, err error) {

	var promiseread js.Value
	var p promise.Promise
	donechan := make(chan bool)
	err = nil

	if promiseread, err = r.JSObject().CallWithErr("read"); err == nil {

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			p.Async(func(obj baseobject.BaseObject) *promise.Promise {

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

			p.Async(func(obj baseobject.BaseObject) *promise.Promise {

				if obj.JSObject().Get("done").Bool() == true {
					err = io.EOF
					dataHandle(nil, err)
					return nil
				} else {
					var u8array uint8array.Uint8Array

					uint8arrayObject := obj.JSObject().Get("value")

					if u8array, err = uint8array.NewFromJSObject(uint8arrayObject); err == nil {

						if _, err = u8array.CopyBytes(preallocateBytes); err == nil {
							dataHandle(preallocateBytes, err)
						}

					}

					p2, _ := promise.New(func(p promise.Promise) (interface{}, error) {
						_, err := r.asyncRead(preallocateBytes, dataHandle)
						return nil, err
					})
					return &p2

				}

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
