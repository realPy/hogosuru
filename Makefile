GO=tinygo

.PHONY: array json draganddrop fetch websocket node xmlhttprequest
all: array json draganddrop fetch websocket node xmlhttprequest



array:
	$(GO) build  -o ./example/static/array.wasm --no-debug -target wasm example/array/main.go
map:
	$(GO) build  -o ./example/static/map.wasm --no-debug -target wasm example/map/main.go
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