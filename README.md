# jswasm

## Synopsis

jswasm is a lib replacement for syscall/js and help to use syscall with handle error (not panic)  
The standard syscall/js doesn't provide a way to handle error and panics when error occurs.  
Panic can be handle via recovery but with framework compiler like tinygo, the recovery can not be handle.  
Some other package like net/http , json is not available but javascript can already make lots of things if error can be handle  use JSON object and parse function with invalid json panic...)  
Jswasm provide function replacement for http, event , json ( and more to come )  

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
go run 
```

Dont forget to get to use the wasm_exec.js provide by tinygo.
