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

				if result, err := openrequest.Result(); err == nil {

					if db, err := indexeddb.IDBDatabaseNewFromObject(result); err == nil {
						if transaction, err := db.Transaction("utilisateur", "readwrite"); err == nil {
							if store, err := transaction.ObjectStore("utilisateur"); err == nil {

								if req, err := store.Add(map[string]interface{}{"email:": "manu"}); err != nil {
									println("erreur", err.Error())
								} else {

									req.OnSuccess(func(e event.Event) {

										println("Add successfull")
									})

								}
							} else {
								println("erreur", err.Error())

							}
						} else {
							println("erreur", err.Error())

						}

					} else {
						println("erreur", err.Error())

					}

				} else {
					println("erreur", err.Error())

				}
			})

			openrequest.OnUpgradeNeeded(func(e event.Event) {

				if result, err := openrequest.Result(); err == nil {

					if db, err := indexeddb.IDBDatabaseNewFromObject(result); err == nil {
						if store, err := db.CreateObjectStore("utilisateur", map[string]interface{}{"keyPath": "id", "autoIncrement": true}); err == nil {

							store.CreateIndex("email", "email", map[string]interface{}{"unique": true})
						} else {
							println("erreur", err.Error())

						}
					}

				} else {
					println("erreur", err.Error())

				}

			})

		} else {
			println("erreur", err.Error())

		}

	}

	ch := make(chan struct{})
	<-ch

}
