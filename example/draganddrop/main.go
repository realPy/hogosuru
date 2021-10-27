package main

import (
	"crypto/sha256"
	"encoding/hex"
	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/dragevent"
	"github.com/realPy/hogosuru/file"
	"github.com/realPy/hogosuru/filelist"
	"github.com/realPy/hogosuru/promise"
)

/*
func md5File(f file.File) string {

	var buffersize int64 = 2 * 1024 * 1024
	stream := blob.NewBlobStream(f.Blob)

	var n int
	var err error
	hashmd5 := md5.New()

	for {
		n, err = stream.Read(data)

		hashmd5.Write(data[:n])
		if err != nil {
			break
		}
	}
	if err == io.EOF {
		return hex.EncodeToString(hashmd5.Sum(nil))
	}

	return ""
}*/

func sha256FileStream(f file.File) string {

	var sha256result string
	if stream, err := f.Stream(); err == nil {

		if read, err := stream.GetReader(); err == nil {

			hashsha256 := sha256.New()

			p, _ := read.AsyncRead(2*1024*1024, func(b []byte, i int) {

				hashsha256.Write(b[:i])
			})
			p.Then(func(i interface{}) *promise.Promise {
				sha256result = hex.EncodeToString(hashsha256.Sum(nil))
				println(f.Name_() + "  SHA256 Stream: " + sha256result)
				return nil
			}, nil)

		} else {
			println(err.Error())
		}
	}
	return ""
}

func sha256File(f file.File) {
	var sha256result string

	//allocate memory in handler is not recommend
	var buffer []byte = make([]byte, 128*1024)
	hashsha256 := sha256.New()

	stream := blob.NewBlobStream(f.Blob)

	p, _ := stream.AsyncRead(buffer, func(b []byte, i int) {
		hashsha256.Write(b[:i])
	})

	p.Then(func(i interface{}) *promise.Promise {

		sha256result = hex.EncodeToString(hashsha256.Sum(nil))
		println(f.Name_() + "  SHA256 Blob: " + sha256result)
		return nil
	}, nil)
}

func dropHandler() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var err error
		var e dragevent.DragEvent
		var dt datatransfer.DataTransfer
		var files filelist.FileList
		var f file.File

		if e, err = dragevent.NewFromJSObject(args[0]); err == nil {
			e.PreventDefault()
			if dt, err = e.DataTransfer(); err == nil {
				if files, err = dt.Files(); err == nil {
					if l, err := files.Length(); err == nil {
						for i := 0; i < l; i++ {
							if f, err = files.Item(i); err == nil {

								//md5sum := md5File(f)
								//println(f.Name() + "  MD5: " + md5sum)
								//sha256sum := sha256File(f)
								//println(f.Name() + "  SHA256: " + sha256sum)
								//sha256FileStream(f)
								sha256File(f)

							}
						}
					}

				}
			}
		}
		if err != nil {
			println(err.Error())
		}

		return nil
	})
}

func dragOverHandler() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := dragevent.NewFromJSObject(args[0]); err == nil {
			e.PreventDefault()
		} else {
			println("erreur", err.Error())
		}
		return nil
	})
}

func main() {
	hogosuru.Init()
	js.Global().Set("dropHandler", dropHandler())

	js.Global().Set("dragOverHandler", dragOverHandler())

	ch := make(chan struct{})
	<-ch

}
