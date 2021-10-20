package blob

import (
	"io"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/typedarray"
)

type BlobStream struct {
	totalsize  int
	buffersize int
	cur        int
	Blob
}

func NewBlobStream(blob Blob, buffersize int) BlobStream {
	var b BlobStream

	b.Blob = blob
	b.cur = 0
	b.buffersize = buffersize
	b.totalsize, _ = blob.Size()

	return b
}

func (b *BlobStream) Read(buffer []byte) (n int, err error) {

	var blob Blob
	var arr arraybuffer.ArrayBuffer
	var rawdata typedarray.Uint8Array
	var bytesneed int
	var done bool = false

	if (b.cur + b.buffersize) > b.totalsize {
		bytesneed = b.totalsize - b.cur
		done = true

	} else {
		bytesneed = b.buffersize
	}

	if blob, err = b.Blob.Slice(b.cur, b.cur+bytesneed); err == nil {

		if arr, err = blob.ArrayBuffer(); err == nil {

			if rawdata, err = typedarray.NewUint8Array(arr); err == nil {

				n, err = rawdata.CopyBytes(buffer)

				b.cur = b.cur + n
				if done {
					err = io.EOF
				}
			}

		}

	}

	return
}

func (b *BlobStream) Write(p []byte) (n int, err error) {

	var arraybuf arraybuffer.ArrayBuffer
	var array8buf typedarray.Uint8Array

	if arraybuf, err = arraybuffer.New(len(p)); err == nil {

		if array8buf, err = typedarray.NewUint8Array(arraybuf); err == nil {
			currentBlob := b.Blob
			array8buf.CopyFromBytes(p)
			b.Blob, err = b.Blob.Append(array8buf.BaseObject)
			currentBlob.Close()
			n = len(p)
		}
	}

	return
}

func (b BlobStream) GetBlob() Blob {
	return b.Blob
}
