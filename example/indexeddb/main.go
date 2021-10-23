package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/event"
	indexeddb "github.com/realPy/hogosuru/indexeddb"
	"github.com/realPy/hogosuru/window"
)

func main() {
	hogosuru.Init()
	if w, err := window.New(); err == nil {

		if factory, err := w.IndexdedDB(); err == nil {
			if openrequest, err := factory.Open("manu", "3"); err == nil {

				openrequest.OnSuccess(func(e event.Event) {

					if result, err := openrequest.Result(); err == nil {

						if db, err := indexeddb.IDBDatabaseNewFromObject(result); err == nil {
							if transaction, err := db.Transaction("utilisateur", "readwrite"); err == nil {
								if store, err := transaction.ObjectStore("utilisateur"); err == nil {
									if req2, err := store.Put(map[string]interface{}{"id": 1, "email": "oui", "prenom": "berndard"}); err == nil {
										req2.OnSuccess(func(e event.Event) {
											//	store.Put(map[string]interface{}{"email": "oui", "prenom": "bernard"})
											println("Put successfull")

										})

										req2.OnError(func(e event.Event) {
											//	e.Export("toto")

											d, _ := e.Target()

											if req, ok := d.(indexeddb.IDBRequest); ok {

												objerr, _ := req.Error()
												strerr, _ := objerr.Message()
												println("-***-->", strerr)
											}

										})
									} else {
										println("erreur", err.Error())
									}

									if req, err := store.Add(map[string]interface{}{"email": "oui", "prenom": "manu"}); err != nil {
										println("erreur", err.Error())
									} else {

										req.OnSuccess(func(e event.Event) {
											//	store.Put(map[string]interface{}{"email": "oui", "prenom": "bernard"})
											println("Add successfull")

										})

										req.OnError(func(e event.Event) {
											//	e.Export("toto")

											d, _ := e.Target()

											if req, ok := d.(indexeddb.IDBRequest); ok {

												objerr, _ := req.Error()
												strerr, _ := objerr.ToString()
												println("----->", strerr)
											}

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

								store.CreateIndex("email", "emailkey", map[string]interface{}{"unique": true})
								store.CreateIndex("nom", "nom")
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

		} else {
			println("erreur", err.Error())
		}
	}

	ch := make(chan struct{})
	<-ch

}
