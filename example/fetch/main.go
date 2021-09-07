package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/json"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

func main() {

	//fetch is async so you must use a channel to wait response
	fetchsync := make(chan bool)
	//dont forget that fetch url need cors
	endpoint, _ := url.Parse("http://localhost:9090/static.json")

	fetch.NewFetch(endpoint.String(), "GET", nil, nil, func(r response.Response, err error) {

		if err == nil {
			if s, _ := r.Status(); s == 200 {
				if text, err := r.Text(); err == nil {

					if j, err := json.Parse(text); err == nil {
						json := j.Map()
						fmt.Printf("Hello %s\n", json.(map[string]interface{})["hello"])
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
	fetch.NewFetch(endpoint.String(), "GET", nil, nil, func(r response.Response, err error) {

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

	fetch.NewFetch(endpoint.String(),
		"POST",
		&map[string]interface{}{"content-type": "application/x-www-form-urlencoded", "User-Agent": "Tester"},
		&dataPost, func(r response.Response, err error) {

			if err == nil {
				if s, _ := r.Status(); s == 200 {
					if text, err := r.Text(); err == nil {

						if j, err := json.Parse(text); err == nil {
							json := j.Map()
							fmt.Printf("[Hello] %s\n", json.(map[string]interface{})["hello"])
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

	if f, err := fetch.New(endpoint.String(), map[string]interface{}{"method": "GET"}); hogosuru.AssertErr(err) {

		f.Then(func(r response.Response) *promise.Promise {

			if header, err := r.Response().Headers(); hogosuru.AssertErr(err) {
				it, _ := header.Entries()
				for key, value, err := it.Next(); err == nil; key, value, err = it.Next() {
					println(key, ":", value)

				}

			}

			return nil

		}, func(e error) {
			hogosuru.AssertErr(e)
		})

	}

	ch := make(chan struct{})
	<-ch

}
