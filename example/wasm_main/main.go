package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/jswasm/http"
	"github.com/realPy/jswasm/indexdb"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/json"
)

func main() {

	if j, err := json.JsonParse("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		fmt.Printf("Value of complex[\"toto\"] %s\n", p["complex"].(map[string]interface{})["toto"])
	} else {
		fmt.Printf("erreur %s", err)
	}
	endpoint, _ := url.Parse("http://localhost:9090/static.json")
	http.HTTPGetText(endpoint, func(status int, text string) {
		if status == 200 {
			if j, err := json.JsonParse(text); err == nil {
				jsonGo := j.GoJson()
				fmt.Printf("Hello %s\n", jsonGo["hello"])
			} else {
				fmt.Printf("erreur %s", err)
			}
		}
	})

	c, _ := indexdb.OpenIndexDB("test", func(db js.Value) error {

		if store, err := indexdb.CreateStore(db, "utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {
			store.CreateIndex("email", "emailkey", map[string]interface{}{"unique": true})
			store.CreateIndex("nom", "nom", nil)
		}
		return nil
	})

	if err := c.Store("utilisateur", map[string]interface{}{"email": "oui", "prenom": "manu"}); err != nil {
		fmt.Println(err.Error())
	}
	ch := make(chan struct{})
	<-ch

}
