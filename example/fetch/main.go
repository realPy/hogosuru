package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/json"
	"github.com/realPy/hogosuru/response"
)

func main() {

	//fetch is async so you must use a channel to wait response
	fetchsync := make(chan bool)
	//dont forget that fetch url need cors
	endpoint, _ := url.Parse("http://localhost:9090/static.json")

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

	//fetch same ressource but as Bytes
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

	//post some data with url encoded
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

	ch := make(chan struct{})
	<-ch

}
