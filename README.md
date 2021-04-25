# jswasm

## Synopsis
jsawsm is a part of a personel project and export to the community and provide a framework to easly provide wabapi and interface access to go

jswasm is a lib replacement for syscall/js and help to use syscall with handle error (not panic) 

The standard syscall/js doesn't provide a way to handle error and panics when error occurs.  
Yes ! , Panic can be handle via recovery but with compiler like tinygo, the recovery can not be handle. https://tinygo.org/lang-support/#a-note-on-the-recover-builtin

Some other package like net/http , json is not available (https://tinygo.org/lang-support/stdlib/) but javascript can already make lots of things with API if error can be handle!   
For example with standard syscall/js, use of JSON API and parse function with invalid json Panic and stopped your current webassembly function. I think that an invalid json is not a good reason to stopped a job, we just want handle the error.

Jswasm provide functions replacements for http, event , json , but with a goal to provide all JS API access https://developer.mozilla.org/fr/docs/Web/API 

## How to use
Just import the module with import "github.com/realPy/jswasm/js" to use the standard syscall/js  
Use the import "github.com/realPy/jswasm" to import http, json and event function  



## try example

Start project in tinygo docker
```
docker run -it -v $PWD:/go/src/jswasm tinygo/tinygo bash
```
Build  

```
cd /go/src/jswasm
tinygo build  -o ./example/wasm_main/wasm.wasm --no-debug -target wasm example/wasm_main/main.go
```

Run server to the the result in js developer console
```
go run example/server/main.go
```

Dont forget to get to use the wasm_exec.js provide by tinygo.
