package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/jswasm/formdata"
	"github.com/realPy/jswasm/htmlinputelement"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/response"

	"github.com/realPy/jswasm/broadcastchannel"
	"github.com/realPy/jswasm/customevent"
	"github.com/realPy/jswasm/document"
	"github.com/realPy/jswasm/fetch"
	"github.com/realPy/jswasm/indexeddb"
	"github.com/realPy/jswasm/object"
	"github.com/realPy/jswasm/storage"
	"github.com/realPy/jswasm/xmlhttprequest"

	"github.com/realPy/jswasm/json"
)

func test() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		endpoint, _ := url.Parse("http://localhost:9090/po")
		files, _ := document.QuerySelector("[name=file]")
		if h, err := htmlinputelement.NewFromJSObject(files); err == nil {
			if file, err := h.Files(); err == nil {
				f, _ := formdata.NewFormData()
				f.AppendString("po", "po")
				fi, _ := file.Item(0)
				f.AppendJSObject("avatar", fi.JSObject())
				if size, err := fi.Size(); err == nil {
					fmt.Printf("size of the file :%d --%s\n", size, fi)
				} else {
					fmt.Println(err.Error())
				}

				if xhr, err := xmlhttprequest.NewXMLHTTPRequest(); err == nil {

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

	if j, err := json.NewJsonFromString("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		fmt.Printf("Value of complex[\"toto\"] %s\n", p.Get("complex").Get("toto"))
		fmt.Printf("---->%s\n", p)
	} else {
		fmt.Printf("erreur %s", err)
	}

	endpoint, _ := url.Parse("http://localhost:9090/static.json")

	fetchsync := make(chan bool)

	fetch.NewFetch(endpoint, "GET", nil, nil, func(r response.Response) {

		if r.Status() == 200 {
			if text, err := r.Text(); err == nil {

				if j, err := json.NewJsonFromString(text); err == nil {
					jsonGo := j.GoJson()
					fmt.Printf("Hello %s\n", jsonGo.Get("hello"))
				} else {
					fmt.Printf("erreur %s", err)
				}

			} else {
				fmt.Println(err.Error())
			}
		}
		fetchsync <- true
	})

	<-fetchsync

	fetch.NewFetch(endpoint, "GET", nil, nil, func(r response.Response) {

		if r.Status() == 200 {
			if b, err := r.ArrayBufferBytes(); err == nil {

				fmt.Printf("Bytes: %s", string(b))

			} else {
				fmt.Println(err.Error())
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
		&dataPost, func(r response.Response) {

			if r.Status() == 200 {
				if text, err := r.Text(); err == nil {

					if j, err := json.NewJsonFromString(text); err == nil {
						jsonGo := j.GoJson()
						fmt.Printf("Hello %s\n", jsonGo.Get("hello"))
					} else {
						fmt.Printf("erreur %s", err)
					}

				} else {
					fmt.Println(err.Error())
				}
			}
			fetchsync <- true
		})

	<-fetchsync

	event, _ := customevent.NewJSCustomEvent("TestEvent", "detail du text")
	event.DispatchEvent(document.Root())

	event.Export("romain")

	if c, err := indexeddb.OpenIndexedDB("test", 3, func(db js.Value) error {

		if store, err := indexeddb.CreateStore(db, "utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {
			store.CreateIndex("email", "emailkey", map[string]interface{}{"unique": true})
			store.CreateIndex("nom", "nom", nil)
		}
		return nil
	}); err == nil {

		if store, err := c.GetObjectStore("utilisateur", "readwrite"); err == nil {
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
			fmt.Println(err.Error())
		}

	} else {
		fmt.Printf("erreur: %s\n", err.Error())
	}

	localstore, _ := storage.GetLocalStorage("session")
	localstore.SetItem("dog", "dalmatien")

	fmt.Println("-----------Test Channels---------")
	if channel, err := broadcastchannel.NewBroadcastChannel("TestChannel"); err == nil {
		channel.SetReceiveMessage(func(c broadcastchannel.Channel, obj object.GOMap) {
			fmt.Printf("--->%s---\n", obj.Get("data").String())
		})

		if err := channel.PostMessage("New wasm loaded"); err != nil {
			fmt.Println(err.Error())
		}
		channel.Export("monchannel")
	} else {
		fmt.Println(err.Error())
	}

	if xhr, err := xmlhttprequest.NewXMLHTTPRequest(); err == nil {
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

	if xhr, err := xmlhttprequest.NewXMLHTTPRequest(); err == nil {

		xhr.Open("POST", endpoint)
		f, _ := formdata.NewFormData()
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

	js.Global().Set("test", test())

	ch := make(chan struct{})
	<-ch

}
