GO=tinygo

.PHONY: hello array json draganddrop fetch websocket node xmlhttprequest map date promise indexeddb history broadcastchannel console webassembly keyobservable
all: hello array json draganddrop fetch websocket node xmlhttprequest map date promise indexeddb history broadcastchannel console webassembly loading keyobservable


hello:
	$(GO) build  -o ./example/static/hello.wasm --no-debug -target wasm example/hello/main.go

array:
	$(GO) build  -o ./example/static/array.wasm --no-debug -target wasm example/array/main.go
map:
	$(GO) build  -o ./example/static/map.wasm --no-debug -target wasm example/map/main.go
date:
	$(GO) build  -o ./example/static/date.wasm --no-debug -target wasm example/date/main.go
promise:
	$(GO) build  -o ./example/static/promise.wasm --no-debug -target wasm example/promise/main.go
indexeddb:
	$(GO) build  -o ./example/static/indexeddb.wasm --no-debug -target wasm example/indexeddb/main.go
	
node:
	$(GO) build  -o ./example/static/node.wasm --no-debug -target wasm example/node/main.go

json:
	$(GO) build  -o ./example/static/json.wasm --no-debug -target wasm example/json/main.go

fetch:
	$(GO) build  -o ./example/static/fetch.wasm --no-debug -target wasm example/fetch/main.go

draganddrop:
	$(GO) build  -o ./example/static/draganddrop.wasm --no-debug -target wasm example/draganddrop/main.go

xmlhttprequest:
	$(GO) build  -o ./example/static/xmlhttprequest.wasm --no-debug -target wasm example/xmlhttprequest/main.go

websocket:
	$(GO) build  -o ./example/static/websocket.wasm --no-debug -target wasm example/websocket/main.go

history:
	$(GO) build  -o ./example/static/history.wasm --no-debug -target wasm example/history/main.go
routing:
	$(GO) build  -o ./example/static/routing.wasm --no-debug -target wasm example/routing/main.go
broadcastchannel:
	$(GO) build  -o ./example/static/broadcastchannel.wasm --no-debug -target wasm example/broadcastchannel/main.go
console:
	$(GO) build  -o ./example/static/console.wasm --no-debug -target wasm example/console/main.go
webassembly:
	$(GO) build  -o ./example/static/webassembly.wasm --no-debug -target wasm example/webassembly/main.go
loading:
	$(GO) build  -o ./example/static/loading.wasm --no-debug -target wasm example/loading/main.go
keyobservable:
	$(GO) build  -o ./example/static/keyobservable.wasm --no-debug -target wasm example/keyobservable/main.go