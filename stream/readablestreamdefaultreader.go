package stream

import (
	"io"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/jserror"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/typedarray"
)

var singletonReadableStreamDefault sync.Once

var readablestreamdefaultinterface js.Value

//GetReadStreamInterface
func GetReadStreamInterface() js.Value {

	singletonReadableStreamDefault.Do(func() {

		var err error
		if readablestreamdefaultinterface, err = baseobject.Get(js.Global(), "ReadableStreamDefaultReader"); err != nil {
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

	if promiseread, err = r.Call("read"); err == nil {

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			p.Then(func(i interface{}) *promise.Promise {
				if obj, ok := i.(baseobject.ObjectFrom); ok {
					if obj.JSObject().Get("done").Bool() == true {
						err = io.EOF
						donechan <- true
						return nil
					} else {

						var u8array typedarray.Uint8Array

						uint8arrayObject := obj.JSObject().Get("value")
						if u8array, err = typedarray.NewUint8Array(uint8arrayObject); err == nil {
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

func (r ReadableStreamDefaultReader) newRead(data []byte, dataHandle func([]byte, int)) *promise.Promise {
	var pp *promise.Promise
	var err error
	var promiseread js.Value

	if promiseread, err = r.Call("read"); err == nil {
		var p promise.Promise

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			newpromise, _ := p.Then(func(i interface{}) *promise.Promise {
				var obj js.Value
				if b, ok := i.(baseobject.ObjectFrom); ok {
					obj = b.JSObject()
					var done bool = false
					if obj.Get("done").Bool() == true {
						done = true
					}

					var u8array typedarray.Uint8Array
					var n int
					uint8arrayObject := obj.Get("value")

					if u8array, err = typedarray.NewUint8Array(uint8arrayObject); err == nil {

						if n, err = u8array.CopyBytes(data); err == nil {
							dataHandle(data, n)
						} else {
							rej, _ := promise.Reject(err)
							return &rej
						}

					}

					if done == false {
						return r.newRead(data, dataHandle)
					} else {
						return nil
					}

				}

				return nil
			}, nil)
			pp = &newpromise

		}
	}
	return pp
}
func (r ReadableStreamDefaultReader) AsyncRead(buffersize int, dataHandle func([]byte, int)) (promise.Promise, error) {

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		var data []byte = make([]byte, buffersize)
		var p *promise.Promise

		p = r.newRead(data, dataHandle)

		p.Then(func(i interface{}) *promise.Promise {
			resolvefunc.Invoke(nil)
			return nil
		}, func(e error) {
			if errjs, err := jserror.New(e); err == nil {
				errfunc.Invoke(errjs.JSObject())
			}
		})

		return nil, nil

	})

}
