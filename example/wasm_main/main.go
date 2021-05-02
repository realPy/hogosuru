package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/jswasm/broadcastchannel"
	"github.com/realPy/jswasm/document"
	"github.com/realPy/jswasm/http"
	"github.com/realPy/jswasm/storage"

	"github.com/realPy/jswasm/customevent"
	"github.com/realPy/jswasm/indexeddb"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/json"
)

func main() {

	if j, err := json.NewJsonFromString("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		fmt.Printf("Value of complex[\"toto\"] %s\n", p.Get("complex").Get("toto"))
		fmt.Printf("---->%s\n", p)
	} else {
		fmt.Printf("erreur %s", err)
	}
	endpoint, _ := url.Parse("http://localhost:9090/static.json")
	http.HTTPGetText(endpoint, func(status int, text string) {
		if status == 200 {
			/*
				if j, err := json.JsonParse(text); err == nil {
					jsonGo := j.GoJson()
					fmt.Printf("Hello %s\n", jsonGo["hello"])
				} else {
					fmt.Printf("erreur %s", err)
				}*/
		}
	})

	event, _ := customevent.NewJSCustomEvent("TestEvent", "detail du text")
	event.DispatchEvent(document.Root())

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
		channel.SetReceiveMessage(func(obj js.Value) {
			fmt.Printf("--->%s---\n", obj.String())
		})

		if err := channel.PostMessage("New wasm loaded"); err != nil {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

	ch := make(chan struct{})
	<-ch

}
