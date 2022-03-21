Hogosuru [![Unit Test](https://github.com/realPy/hogosuru/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/realPy/hogosuru/actions/workflows/ci.yml)
=========


Hogosuru is a framework to easily write a complete single page application in Go or just export some functionnality directly accessible in javascript (functionnality that need speed for example).  

## Documentation

A detailed documentation with code examples can be found here : https://hogosuru.v-ip.fr/

## How it work?
He use an addon syscall/js (that catch error) and implement a large part of the major features of javascript directly accessible in go.  

For a list of functionnality , see the array of compatibility below or check the name of directory that match the API MDN implementation: https://developer.mozilla.org/fr/docs/Web/API  
   
Each object is herited from a baseobject.BaseObject. The baseobject keep a true JS object reference.  
   
All structure available can interact with this object and use the internal functionnality  of each object.  
If an error occur , this error is handle and return by each function.  
   
## How to use

Just import the lib in your project (hogosuru no need use an extended rewrite of base syscall)

```

GOOS=js GOARCH=wasm go get github.com/realPy/hogosuru
```
The project work with Go and Tinygo compiler.
Use Go compiler for developpement (faster) and tinygo for production


## Use the special "autodiscover" of hogosuru

The autodiscover is a special function. 

When a function receives a data which is global object (exemple with the attribute data of an event) or which is not an expected data, it is possible to ask hogosuru to guess the content of the object and to create the corresponding Go object.

This is not magic, and it only works if the type of object has already been seen before (hogosuru keeps in memory the type of objects and their constructor).  
It would be possible to register all the objects known to hogosuru but that would amount to integrating the whole implementation in your binary. 
When you use an object (by use it or cast it) , the constructor is loaded automatically.
If the type is not known, autodiscover will return an object of type BaseObject



## Chaining capabilities

Some structure like nodelist or document can be used with chaining capabilities (function that return one object that will be used immediatly without check error).It's a shortcut that ignores errors.
The functions which allow this functionality are by convention with the same prototype, with the same name but which end with "_". It is a convention created for hogosuru. All the functions available in chaining are contained in the "chaining.go" file at the root of each component



## How to help

This is a young project and there are a lot of work to do  
All help is welcome. If you are interested by this project, please contact me


## Implemented API/Object status 

    

|  API/Object |  Implemented Support |  MDN URL |
|-------------|:--------------------:|----------|
| AbortController | Full | https://developer.mozilla.org/en-US/docs/Web/API/AbortController | 
| AbortSignal | Full | https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal | 
| AnimationEvent | Full | https://developer.mozilla.org/fr/docs/Web/API/AnimationEvent | 
| Array | Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Array | 
| Arraybuffer |  Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer | 
| Attr | Full | https://developer.mozilla.org/fr/docs/Web/API/Attr |
| Blob | Full |  https://developer.mozilla.org/fr/docs/Web/API/Blob |
| Broadcast Channel |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel  | 
| Clipboard |  Full | https://developer.mozilla.org/en-US/docs/Web/API/Clipboard |
| ClipboardItem | Full | https://developer.mozilla.org/en-US/docs/Web/API/ClipboardItem |
| ClipboardEvent | Full | https://developer.mozilla.org/en-US/docs/Web/API/ClipboardEvent | 
| Console |  Full |  https://developer.mozilla.org/fr/docs/Web/API/Console  | 
| CSSRule |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/CSSRule | 
| CSSStyleDeclaration |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration | 
| CustomEvent |  Full |  https://developer.mozilla.org/fr/docs/Web/API/CustomEvent |
| DataTransfer | Full | https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer |
| DataTransferItem | Partial | https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItem |
| DataTransferItemList | Full | https://developer.mozilla.org/en-US/docs/Web/API/DataTransferItemList |
| Date| Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Date | 
| DedicatedWorkerGlobalScope | Full | https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope |
| Document | Mostly  | https://developer.mozilla.org/fr/docs/Web/API/Document | 
| DragEvent |  Full |  https://developer.mozilla.org/en-US/docs/Web/API/DragEvent |
| Element | Full | https://developer.mozilla.org/fr/docs/Web/API/Element | 
| Event | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Event |
| EventTarget | Full | https://developer.mozilla.org/fr/docs/Web/API/EventTarget/EventTarget | 
| Fetch | Full | https://developer.mozilla.org/fr/docs/Web/API/Fetch_API |
| File | Full | https://developer.mozilla.org/fr/docs/Web/API/File |
| FileList | Full | https://developer.mozilla.org/fr/docs/Web/API/FileList |
| FormData | Full | https://developer.mozilla.org/fr/docs/Web/API/FormData |
| Headers | Full | https://developer.mozilla.org/en-US/docs/Web/API/Headers |
| History | Full | https://developer.mozilla.org/fr/docs/Web/API/History |
| HTMLAnchorElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement|
| HTMLAreaElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement|
| HTMLBaseElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLBaseElement|
| HTMLBodyElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLBodyElement|
| HTMLBRElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLBRElement|
| HTMLButtonElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement|
| HTMLCollection| Full | https://developer.mozilla.org/fr/docs/Web/API/HTMLCollection |
| HTMLDataElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLDataElement|
| HTMLDataListElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLDataListElement|
| HTMLDetailsElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLDetailsElement|
| HTMLDivElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLDivElement|
| HTMLDListElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLDListElement|
| HTMLElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement|
| HTMLEmbbedElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLEmbbedElement|
| HTMLFieldSetElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLFieldSetElement|
| HTMLFormElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement|
| HTMLHeadElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement|
| HTMLHeadingElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadingElement|
| HTMLHRElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLHRElement|
| HTMLIFrameElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement|
| HTMLImageElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement|
| HTMLInputElement| Full| https://developer.mozilla.org/fr/docs/Web/API/HTMLInputElement |
| HTMLLabelElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement|
| HTMLLegendElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLLegendElement|
| HTMLLIElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLLIElement|
| HTMLLinkElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement|
| HTMLMapElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLMapElement|
| HTMLMetaElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLMetaElement|
| HTMLMeterElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLMeterElement|
| HTMLIFrameElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement|
| HTMLOptionsCollection| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLOptionsCollection|
| HTMLParagraphElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLParagraphElement|
| HTMLQuoteElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLQuoteElement|
| HTMLScriptElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLScriptElement|
| HTMLSelectElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement|
| HTMLSourceElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLSourceElement|
| HTMLSpanElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLSpanElement|
| HTMLStyleElement | Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLStyleElement |
| HTMLCaptionElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLCaptionElement|
| HTMLTableCaptionElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableCaptionElement|
| HTMLTableCellElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableCellElement|
| HTMLTableColElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableColElement|
| HTMLTableElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement|
| HTMLTableRowElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableRowElement|
| HTMLTableSectionElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableSectionElement|
| HTMLTemplateElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTemplateElement|
| HTMLTextAreaElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTextAreaElement|
| HTMLTimeElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTimeElement|
| HTMLTitleElement| Full| https://developer.mozilla.org/en-US/docs/Web/API/HTMLTitleElement|
| Indexeddb | Full | https://developer.mozilla.org/fr/docs/Web/API/IndexedDB_API |
| Iterator | - | - |
| JSON | Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON |
| Location | Full | https://developer.mozilla.org/fr/docs/Web/API/window/location |
| MessageEvent | Full | https://developer.mozilla.org/fr/docs/Web/API/MessageEvent |
| NamedNodeMap | Full | https://developer.mozilla.org/fr/docs/Web/API/NamedNodeMap |
| Navigator | Partial | https://developer.mozilla.org/en-US/docs/Web/API/Navigator | 
| NavigatorPreloadManager | Full | https://developer.mozilla.org/en-US/docs/Web/API/NavigationPreloadManager |
| Node | Full | https://developer.mozilla.org/en-US/docs/Web/API/Node |
| NodeList | Considerated at Full (Partial implemented but no more need )| https://developer.mozilla.org/fr/docs/Web/API/NodeList | 
| Object | Partial| https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Object | 
| Map | Full | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Map |
| Permissions | Full | https://developer.mozilla.org/en-US/docs/Web/API/Permissions_API | 
| PermissionStatus | Full | https://developer.mozilla.org/en-US/docs/Web/API/PermissionStatus |
| ProgressEvent | Full | https://developer.mozilla.org/en-US/docs/Web/API/ProgressEvent |
| Promise | Full | https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise |
| PushManager | Full | https://developer.mozilla.org/en-US/docs/Web/API/PushManager |
| ReadableStream | Full | https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream |
| Response | Full | https://developer.mozilla.org/fr/docs/Web/API/Response |
| ServiceWorker | Full | https://developer.mozilla.org/en-US/docs/Web/API/ServiceWorker|
| ServiceWorkerContainer | Full | https://developer.mozilla.org/en-US/docs/Web/API/ServiceWorkerContainer|
| ServiceWorkerRegistration | Full | https://developer.mozilla.org/en-US/docs/Web/API/ServiceWorkerRegistration |
| Storage | Full | https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage |
| Stream | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/Streams_API |
| StyleSheet | Full | https://developer.mozilla.org/en-US/docs/Web/API/StyleSheet |
| Uint8array | Partial implemented | https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/Uint8Array 
| UrlSearchParams | Full | https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams |
| Url | Full | https://developer.mozilla.org/en-US/docs/Web/API/URL | 
| ValidityState | Full | https://developer.mozilla.org/en-US/docs/Web/API/ValidityState |
| Websocket | Full | https://developer.mozilla.org/fr/docs/Web/API/WebSocket |
| Webassembly | Partial |https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/WebAssembly|
| Window | Partial |https://developer.mozilla.org/en-US/docs/Web/API/Window |
| Worker | Full | https://developer.mozilla.org/en-US/docs/Web/API/Worker |
| WorkerGlobalScope| Partial | https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope |
| XMLHttpRequest | Partial implemented | https://developer.mozilla.org/fr/docs/Web/API/XMLHttpRequest/XMLHttpRequest |




## Custom components

### A toaster components for hogosuru

https://github.com/realPy/hogosurutoaster

### A pagination components for hogosuru with templating

https://github.com/realPy/hogosurupagination


### A datatable components for hogosuru with templating

https://github.com/realPy/hogosurudatatable

## How to help

All help is welcome. If you are interested by this project, please contact me.  

