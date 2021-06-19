// Copyright 202  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func (v Value) SetWithErr(p string, x interface{}) error {

	if vType := v.Type(); !vType.isObject() {
		return errors.New("Unable to set value" + p)
	}
	xv := ValueOf(x)
	valueSet(v.ref, p, xv.ref)
	runtime.KeepAlive(v)
	runtime.KeepAlive(xv)
	return nil
}

func (v Value) CallWithErr(m string, args ...interface{}) (Value, error) {
	argVals, argRefs := makeArgs(args)
	res, ok := valueCall(v.ref, m, argRefs)
	runtime.KeepAlive(v)
	runtime.KeepAlive(argVals)
	if !ok {
		if vType := v.Type(); !vType.isObject() { // check here to avoid overhead in success case
			return Value{}, errors.New("Value.Call on invalid object")
		}
		if propType := v.Get(m).Type(); propType != TypeFunction {
			return Value{}, errors.New("syscall/js: Value.Call: property " + m + " is not a function, got " + propType.String())

		}

		return Value{}, errors.New(makeValue(res).Get("message").String())

	}
	return makeValue(res), nil
}

func (v Value) GetWithErr(p string) (Value, error) {
	if vType := v.Type(); !vType.isObject() {
		return Value{}, errors.New("Unable to get value" + p)
	}
	r := makeValue(valueGet(v.ref, p))
	runtime.KeepAlive(v)
	return r, nil
}

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
