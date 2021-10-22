// Copyright 202  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

// Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
// Its API is based on JavaScript semantics.
//
// This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a
// comprehensive API for users. It is exempt from the Go compatibility promise.
package js

import (
	"errors"
	"runtime"
)

func CopyBytesToGoWithErr(dst []byte, src Value) (int, error) {
	n, ok := copyBytesToGo(dst, src.ref)
	runtime.KeepAlive(src)
	if !ok {
		return 0, errors.New("syscall/js: CopyBytesToGo: expected src to be an Uint8Array or Uint8ClampedArray")
	}
	return n, nil
}

func CopyBytesToJSWithErr(dst Value, src []byte) (int, error) {
	n, ok := copyBytesToJS(dst.ref, src)
	runtime.KeepAlive(dst)
	if !ok {
		return 0, errors.New("syscall/js: CopyBytesToJS: expected dst to be an Uint8Array or Uint8ClampedArray")

	}
	return n, nil
}

func AsyncFuncOf(fn func(this Value, args []Value) interface{}) Func {
	funcsMu.Lock()
	id := nextFuncID
	nextFuncID++
	funcs[id] = fn
	funcsMu.Unlock()
	return Func{
		id:    id,
		Value: jsGo.Call("_makeAsyncFuncWrapper", id),
	}
}
