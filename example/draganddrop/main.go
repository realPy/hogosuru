package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"syscall/js"

	"github.com/realPy/hogosuru/blob"
	datatransfert "github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/dragevent"
	"github.com/realPy/hogosuru/file"
	"github.com/realPy/hogosuru/filelist"
)

func md5File(f file.File) string {

	var buffersize int = 2 * 1024 * 1024
	stream := blob.NewBlobStream(f.Blob, buffersize)

	var data []byte = make([]byte, buffersize)
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
}

func dropHandler() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var err error
		var e dragevent.DragEvent
		var dt datatransfert.DataTransfer
		var files filelist.FileList
		var f file.File

		if e, err = dragevent.NewFromJSObject(args[0]); err == nil {
			e.PreventDefault()
			if dt, err = e.DataTransfer(); err == nil {
				if files, err = dt.Files(); err == nil {
					for i := 0; i < files.Length(); i++ {
						if f, err = files.Item(i); err == nil {
							md5sum := md5File(f)
							println(f.Name() + "  MD5: " + md5sum)

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
		}
		return nil
	})
}

func main() {

	js.Global().Set("dropHandler", dropHandler())

	js.Global().Set("dragOverHandler", dragOverHandler())

	ch := make(chan struct{})
	<-ch

}
