Hogosuru
=========
Hogosuru is a part of a personel project and export to the community and provide a framework to easly provide write front end in go.
He use a rewrite syscall/js (that catch error) and wrap Object and Interface provide by browser 

Hogosuru provide functions replacements for http,event,fetch,json,xmlhttprequest,blob,broadcastchannel,indexeddb, localstorage ... with a goal to provide all JS API access https://developer.mozilla.org/fr/docs/Web/API and write full web app in GO


## How to use
Just import the module with go module  
go get github.com/realPy/hogosuru  

The project is written to work with tinygo but compatible with the go official compiler


## Try example

Start project in tinygo docker
```
docker run -it -v $PWD:/go/src/hogosuru tinygo/tinygo bash
```
Build  

```
docker run --rm -it -w /go/src/hogosuru -v $PWD:/go/src/hogosuru tinygo/tinygo tinygo build  -o ./example/wasm_main/wasm.wasm --no-debug -target wasm example/wasm_main/main.go

```

Run server to the the result in js developer console
```
go run example/server/main.go
```

Dont forget to get to use the wasm_exec.js provide by tinygo.


## How to help

This is a young project and there are a lot of work to do  
All help is welcome. If you are interested by this project, please contact me
