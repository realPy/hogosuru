package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"

	"syscall/js"

	"github.com/realPy/hogosuru/json"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/broadcastchannel"
	"github.com/realPy/hogosuru/customevent"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/formdata"
	"github.com/realPy/hogosuru/htmlinputelement"
	"github.com/realPy/hogosuru/indexeddb"
	"github.com/realPy/hogosuru/indexeddb/idbdatabase"
	"github.com/realPy/hogosuru/messageevent"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/response"
	"github.com/realPy/hogosuru/storage"
	"github.com/realPy/hogosuru/websocket"

	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/xmlhttprequest"
)

func TestBlob() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("Click Blob")

		/*b, _ := */
		u, _ := arraybuffer.New(0)

		b, _ := blob.NewWithArrayBuffer(u)

		s, _ := b.Text()

		var buffersize int = 2 * 1024 * 1024
		stream2 := blob.NewBlobStream(b, buffersize)
		stream2.Write([]byte("pouet\000popo"))
		c := stream2.Blob
		c.Export("debug2")
		s, _ = c.Text()
		println("****" + s)

		doc := document.New()

		files, _ := doc.QuerySelector("[name=file]")
		if h, err := htmlinputelement.NewFromJSObject(files); err == nil {
			if file, err := h.Files(); err == nil {

				if firstfile, err := file.Item(0); err == nil {

					if stream, err := firstfile.Stream(); err == nil {
						if read, err := stream.GetReader(); err == nil {
							var data []byte = make([]byte, 2*1024*1024)
							var n int
							var err error
							hashmd5 := md5.New()
							hashsha256 := sha256.New()
							for {
								n, err = read.Read(data)
								hashmd5.Write(data[:n])
								hashsha256.Write(data[:n])
								if err != nil {
									break
								}
							}
							if err == io.EOF {
								println("MD5: " + hex.EncodeToString(hashmd5.Sum(nil)))
								println("SHA256: " + hex.EncodeToString(hashsha256.Sum(nil)))
							}

						} else {
							println(err.Error())
						}
					} else {
						var buffersize int = 2 * 1024 * 1024
						stream := blob.NewBlobStream(firstfile.Blob, buffersize)

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
							println("MD5: " + hex.EncodeToString(hashmd5.Sum(nil)))
						}

						str, _ := firstfile.Blob.Text()
						println("content " + str)
						//println(err.Error())
					}
				} else {
					println(err.Error())
				}

			}
		}

		return nil
	})
}

func test() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		endpoint, _ := url.Parse("http://localhost:9090/po")
		doc, _ := document.New()

		files, _ := doc.QuerySelector("[name=file]")
		if h, err := htmlinputelement.NewFromJSObject(files); err == nil {
			if file, err := h.Files(); err == nil {
				f, _ := formdata.New()
				f.AppendString("po", "po")
				fi, _ := file.Item(0)

				f.AppendJSObject("avatar", fi.JSObject())
				if size, err := fi.Size(); err == nil {
					fmt.Printf("size of the file :%d --%s\n", size, fi)
				} else {
					fmt.Println(err.Error())
				}

				if xhr, err := xmlhttprequest.New(); err == nil {

					xhr.Open("POST", endpoint)

					xhr.SetOnload(func(x xmlhttprequest.XMLHTTPRequest) {

						fmt.Printf("XML HTTPRequest Loaded\n")

						if text, err := x.ResponseText(); err == nil {
							fmt.Printf("Resultat: %s\n", text)
						}

						if header, err := x.GetResponseHeader("Content-Type"); err == nil {
							fmt.Printf("Resultat: %s\n", header)
						}

					})
					xhr.SendForm(f)
				}

			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return nil
	})
}

func main() {

	if j, err := json.Parse("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		fmt.Printf("Value of complex[\"toto\"] %s\n", p.Get("complex").Get("toto"))
		fmt.Printf("---->%s\n", p)
		j.Export("bastien")
	} else {
		fmt.Printf("erreur %s\n", err)
	}

	endpoint, _ := url.Parse("http://localhost:9090/static.json")

	fetchsync := make(chan bool)

	fetch.NewFetch(endpoint, "GET", nil, nil, func(r response.Response, err error) {

		if err == nil {
			if s, _ := r.Status(); s == 200 {
				if text, err := r.Text(); err == nil {

					if j, err := json.Parse(text); err == nil {
						jsonGo := j.GoJson()
						fmt.Printf("Hello %s\n", jsonGo.Get("hello"))
					} else {
						fmt.Printf("erreur %s", err)
					}

				} else {
					fmt.Println(err.Error())
				}
			}
		}

		fetchsync <- true
	})

	<-fetchsync

	fetch.NewFetch(endpoint, "GET", nil, nil, func(r response.Response, err error) {

		if err == nil {
			if s, _ := r.Status(); s == 200 {
				if b, err := r.ArrayBufferBytes(); err == nil {

					fmt.Printf("-----------------------Bytes: %s", string(b))

				} else {
					fmt.Println(err.Error())
				}
			}
		}

		fetchsync <- true
	})
	<-fetchsync

	dataPost := url.Values{}

	dataPost.Set("test", "ok")

	fetch.NewFetch(endpoint,
		"POST",
		&map[string]interface{}{"content-type": "application/x-www-form-urlencoded", "User-Agent": "Tester"},
		&dataPost, func(r response.Response, err error) {

			if err == nil {
				if s, _ := r.Status(); s == 200 {
					if text, err := r.Text(); err == nil {

						if j, err := json.Parse(text); err == nil {
							jsonGo := j.GoJson()
							fmt.Printf("Hello %s\n", jsonGo.Get("hello"))
						} else {
							fmt.Printf("erreur %s", err)
						}

					} else {
						fmt.Println(err.Error())
					}
				}
			}

			fetchsync <- true
		})

	<-fetchsync

	event, _ := customevent.New("TestEvent", "detail du text")
	doc, _ := document.New()
	doc.DispatchEvent(event.Event)

	event.Export("romain")

	indexeddb.Open("test", 3, func(i idbdatabase.IDBDatabase) error {

		if store, err := i.CreateStore("utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {
			store.CreateIndex("email", "emailkey", map[string]interface{}{"unique": true})
			store.CreateIndex("nom", "nom", nil)
		}
		return nil
	}, func(i idbdatabase.IDBDatabase) error {

		if store, err := i.GetObjectStore("utilisateur", "readwrite"); err == nil {
			fmt.Printf("get store..\n")
			if objadd, err := store.Add(map[string]interface{}{"email": "oui", "prenom": "manu"}); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Object add: %d\n", objadd)
				store.Put(map[string]interface{}{"id": objadd, "email": "oui", "prenom": "lea"})
			}
			a, _ := store.GetAllKeys()
			fmt.Printf("%s\n", a)
			if b, err := store.Get(1); err == nil {
				fmt.Printf("object 1: %s\n", b)
			} else {
				fmt.Println(err.Error())
			}

			if all, err := store.GetAll(); err == nil {
				fmt.Printf("all: %s\n", all)
			} else {
				fmt.Println(err.Error())
			}

			if count, err := store.Count(); err == nil {
				if (count % 10) == 0 {
					store.Clear()
				}
				fmt.Printf("Store count: %d\n", count)
			} else {
				fmt.Println(err.Error())
			}

		} else {
			fmt.Println("--->" + err.Error())
		}

		return nil
	}, func(err error) {
		fmt.Printf("erreur: %s\n", err.Error())
	},
	)

	localstore, _ := storage.New("session")
	localstore.SetItem("dog", "dalmatien")

	fmt.Println("-----------Test Channels---------")
	if channel, err := broadcastchannel.New("TestChannel"); err == nil {
		channel.SetOnMessage(func(c broadcastchannel.Channel, m messageevent.MessageEvent) {
			if dataObject, err := m.Data(); err == nil {
				fmt.Printf("--->%s---\n", dataObject.String())
			}

		})

		if err := channel.PostMessage("New wasm loaded"); err != nil {
			fmt.Println(err.Error())
		}
		channel.Export("monchannel")
	} else {
		fmt.Println(err.Error())
	}

	if xhr, err := xmlhttprequest.New(); err == nil {
		endpoint, _ := url.Parse("http://localhost:9090/static.json")
		xhr.Open("GET", endpoint)
		xhr.SetOnload(func(x xmlhttprequest.XMLHTTPRequest) {

			fmt.Printf("XML HTTPRequest Loaded\n")

			if text, err := x.ResponseText(); err == nil {
				fmt.Printf("Resultat: %s\n", text)
			}

			if header, err := x.GetResponseHeader("Content-Type"); err == nil {
				fmt.Printf("Resultat: %s\n", header)
			}

		})

		xhr.SetOnProgress(func(x xmlhttprequest.XMLHTTPRequest, g object.GOMap) {
			fmt.Printf("On progress :%s\n", g)
		})
		xhr.Send()

	}

	if xhr, err := xmlhttprequest.New(); err == nil {

		xhr.Open("POST", endpoint)
		f, _ := formdata.New()
		f.AppendString("data", "pouet")
		xhr.SetOnload(func(x xmlhttprequest.XMLHTTPRequest) {

			fmt.Printf("XML HTTPRequest Loaded\n")

			if text, err := x.ResponseText(); err == nil {
				fmt.Printf("Resultat: %s\n", text)
			}

			if header, err := x.GetResponseHeader("Content-Type"); err == nil {
				fmt.Printf("Resultat: %s\n", header)
			}

		})
		xhr.SendForm(f)
	}

	if ws, err := websocket.New("wss://echo.websocket.org"); err == nil {
		ws.SetOnMessage(func(w websocket.WebSocket, message interface{}) {
			if a, ok := message.(arraybuffer.ArrayBuffer); ok {
				println("Ws receive:" + a.String())
			}
			if s, ok := message.(string); ok {
				println("Ws receive:" + s)
			}

		})

		ws.SetOnOpen(func(ws websocket.WebSocket) {

			ws.Send("hello")
		})

	}

	js.Global().Set("test", TestBlob())
	println("loaded")
	ch := make(chan struct{})
	<-ch

}
