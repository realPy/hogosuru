package blob

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/arraybuffer"
	"github.com/realPy/hogosuru/base/jserror"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/typedarray"
)

type BlobStream struct {
	Blob
	cur        int64
	buffersize int64
	size       int64
}

func NewBlobStream(blob Blob) BlobStream {
	var b BlobStream
	b.Blob = blob
	return b
}

func (b *BlobStream) newRead(data []byte, dataHandle func([]byte, int)) *promise.Promise {
	var pp *promise.Promise

	var bytesneed int64
	var done bool

	if (b.cur + b.buffersize) > b.size {
		bytesneed = b.size - b.cur
		done = true

	} else {
		bytesneed = b.buffersize
	}

	if blob, err := b.Slice(b.cur, b.cur+bytesneed); err == nil {

		if p, err := blob.ArrayBuffer(); err == nil {

			p1, _ := p.Then(func(i interface{}) *promise.Promise {
				var rawdata typedarray.Uint8Array

				if rawdata, err = typedarray.NewUint8Array(i.(arraybuffer.ArrayBuffer)); err == nil {
					var n int
					if n, err = rawdata.CopyBytes(data); err == nil {

						dataHandle(data, n)
					} else {
						rej, _ := promise.Reject(err)
						return &rej
					}

					b.cur = b.cur + int64(n)

					if done == false {
						return b.newRead(data, dataHandle)
					} else {
						return nil
					}
				}

				return nil
			}, nil)

			pp = &p1

		}

	}
	return pp
}

func (b *BlobStream) AsyncRead(data []byte, dataHandle func([]byte, int)) (promise.Promise, error) {

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		var p *promise.Promise

		b.cur = 0
		b.buffersize = int64(len(data))
		b.size, _ = b.Blob.Size()

		p = b.newRead(data, dataHandle)

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

func (b *BlobStream) Write(p []byte) (n int, err error) {

	var arraybuf arraybuffer.ArrayBuffer
	var array8buf typedarray.Uint8Array

	if arraybuf, err = arraybuffer.New(len(p)); err == nil {

		if array8buf, err = typedarray.NewUint8Array(arraybuf); err == nil {
			currentBlob := b.Blob
			array8buf.CopyFromBytes(p)
			b.Blob, err = New(b.Blob, array8buf)
			currentBlob.Close()
			n = len(p)
		}
	}

	return
}

func (b BlobStream) GetBlob() Blob {
	return b.Blob
}
