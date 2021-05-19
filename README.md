Hogosuru
=========
Hogosuru is a part of a personel project and export to the community and provide a framework to easly provide write front end in go.
He use a rewrite syscall/js (that catch error) and wrap Object and Interface provide by browser 

Hogosuru provide implementation for the WEB API  http,event,fetch,json,xmlhttprequest,blob,broadcastchannel,indexeddb, localstorage ... with a goal to provide all JS API access https://developer.mozilla.org/fr/docs/Web/API and write full web app in GO


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
docker run --rm -it -w /go/src/hogosuru -v $PWD:/go/src/hogosuru tinygo/tinygo bash 
cp hogosuru/hogosuru.go /usr/local/go/src/syscall/js/
tinygo build  -o ./example/wasm_main/wasm.wasm --no-debug -target wasm example/wasm_main/main.go
tinygo build  -o ./example/wasm_main/dragandrop.wasm --no-debug -target wasm example/wasm_main/dragandrop/main.go
```

Run server to the the result in js developer console
```
go run example/server/main.go
```

Dont forget to get to use the wasm_exec.js provide by tinygo.


## How to help

This is a young project and there are a lot of work to do  
All help is welcome. If you are interested by this project, please contact me


## Implemented API/Object status

    

|  API/Object |  Implemented Support |  MDN URL |
|-------------|:--------------------:|----------|
| Arraybuffer |  Partial | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer | 
| Attr | Full | https://developer.mozilla.org/fr/docs/Web/API/Attr |
| Blob | Full + stream additional support|  https://developer.mozilla.org/fr/docs/Web/API/Blob |
| Broadcast Channel |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel  | 
| CustomEvent |  Full |  https://developer.mozilla.org/fr/docs/Web/API/CustomEvent |
| DataTransfer | Partial implemented | https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer |
| Document | Mostly  | https://developer.mozilla.org/fr/docs/Web/API/Document | 
| DragEvent |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/DragEvent |
| Element | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Element | 
| Event | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Event |
| EventTarget | Full | https://developer.mozilla.org/fr/docs/Web/API/EventTarget/EventTarget | 
| Fetch | Partial implemented  | https://developer.mozilla.org/fr/docs/Web/API/Fetch_API |
| File | Full | https://developer.mozilla.org/fr/docs/Web/API/File |
| FileList | Full | https://developer.mozilla.org/fr/docs/Web/API/FileList |
| FormData | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/FormData |
| HTMLInputElement| Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement |
| HTMLCollection| Full | https://developer.mozilla.org/fr/docs/Web/API/HTMLCollection |
| Indexedddb | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/IndexedDB_API |
| JSON | Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON |
| MessageEvent | Full | https://developer.mozilla.org/fr/docs/Web/API/MessageEvent |
| Node | Full | https://developer.mozilla.org/en-US/docs/Web/API/Node |
| NodeList | Considerated at Full (Partial implemented but no more need )| https://developer.mozilla.org/fr/docs/Web/API/NodeList |
| Response | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Response |
| Storage | Full | https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage |
| Stream | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Streams_API |
| uint8array | Partial implemented | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Uint8Array |
| WebSocket | Full | https://developer.mozilla.org/fr/docs/Web/API/WebSocket |
| XMLHttpRequest | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/XMLHttpRequest/XMLHttpRequest |







### Local MD5 sum of file with drag and drop

You can test a local example of local hash  
Go to https://realpy.github.io/hogosuru/example/static/draganddrop.html

Open you console developer. Drag some file and see result :)  




