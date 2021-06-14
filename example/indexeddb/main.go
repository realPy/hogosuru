package main

import (
	"github.com/realPy/hogosuru/event"
	indexeddb "github.com/realPy/hogosuru/indexeddb"
)

func main() {

	if factory, err := indexeddb.GetIDBFactory(); err != nil {
		println("erreur", err.Error())
	} else {
		if openrequest, err := factory.Open("manu", "3"); err == nil {
			openrequest.OnSuccess(func(e event.Event) {
				println("success")
				if deleterequest, err := factory.DeleteDatabase("manu", "3"); err != nil {
					println("erreur", err.Error())
				} else {

					println("Delete successfull")
					deleterequest.OnSuccess(func(e event.Event) {
						println("succesfull delete")
					})
					deleterequest.OnError(func(e event.Event) {
						println("error delete")
					})
				}

			})
			openrequest.Export("manu")
		} else {
			println("erreur", err.Error())

		}

	}

	ch := make(chan struct{})
	<-ch

}
