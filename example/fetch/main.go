package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/json"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
	"github.com/realPy/hogosuru/typedarray"
)

func main() {
	hogosuru.Init()
	//fetch is async so you must use a channel to wait response
	fetchsync := make(chan bool)

	//dont forget that fetch url need cors
	d, _ := document.New()

	hogosuru.AssertErr(d.SetDomain("toto.com"))

	endpoint, _ := url.Parse("http://localhost:9090/static.json")

	if p, err := fetch.New(endpoint.String()); hogosuru.AssertErr(err) {

		p2, _ := p.Then(func(r response.Response) *promise.Promise {

			var ptext promise.Promise
			var err error
			var s int
			if s, err = r.Status(); err == nil && s == 400 {

				ptext, _ = r.Text()

			} else {

				if err == nil {
					ptext, _ = promise.Reject(errors.New("NO HTTP CODE 200"))
				} else {
					ptext, _ = promise.Reject(err)
				}

			}
			return &ptext
		}, nil)

		p2.Then(func(i interface{}) *promise.Promise {

			fetchsync <- true
			return nil
		}, func(e error) {
			fetchsync <- true
		})

	}

	<-fetchsync

	println("-->")
	if p, err := fetch.New(endpoint.String()); hogosuru.AssertErr(err) {

		p.Then(func(r response.Response) *promise.Promise {
			if s, err := r.Status(); hogosuru.AssertErr(err) && s == 200 {

				if p2, err := r.Text(); hogosuru.AssertErr(err) {

					p2.Then(func(i interface{}) *promise.Promise {

						if j, err := json.Parse(i.(string)); hogosuru.AssertErr(err) {
							json := j.Map()
							fmt.Printf("Hello %s\n", json.(map[string]interface{})["hello"])
						}
						fetchsync <- true
						return nil
					}, func(e error) {
						hogosuru.AssertErr(err)
						fetchsync <- true
					})

				}

			}
			return nil
		}, func(e error) {

			hogosuru.AssertErr(e)
			fetchsync <- true
		})

	}

	<-fetchsync

	//same thing with an another way with  linked promise

	if p, err := fetch.New(endpoint.String()); hogosuru.AssertErr(err) {

		lp, _ := p.Then(func(r response.Response) *promise.Promise {
			var p2 promise.Promise
			var err error
			var s int
			if s, err = r.Status(); hogosuru.AssertErr(err) && s == 200 {

				p2, err = r.Text()
				hogosuru.AssertErr(err)

			} else {
				p2, _ = promise.Reject(errors.New("Unvalid HTTP status"))
			}

			if err != nil {
				p2, _ = promise.Reject(err)
			}
			return &p2
		}, func(e error) {

			hogosuru.AssertErr(e)
			fetchsync <- true
		})

		lp.Then(func(i interface{}) *promise.Promise {

			if j, err := json.Parse(i.(string)); hogosuru.AssertErr(err) {
				json := j.Map()
				fmt.Printf("Hello %s\n", json.(map[string]interface{})["hello"])
			}

			fetchsync <- true
			return nil
		}, func(e error) {
			hogosuru.AssertErr(e)
			fetchsync <- true
		})

	}

	<-fetchsync

	//same with get data as Arraybuffer
	println("------array buffer---")
	if p, err := fetch.New(endpoint.String()); hogosuru.AssertErr(err) {

		p.Then(func(r response.Response) *promise.Promise {
			if s, err := r.Status(); hogosuru.AssertErr(err) && s == 200 {

				if p2, err := r.ArrayBuffer(); hogosuru.AssertErr(err) {

					p2.Then(func(i interface{}) *promise.Promise {

						a, _ := typedarray.NewUint8Array(i.(arraybuffer.ArrayBuffer).JSObject())
						println("---->" + a.ToString_())
						fetchsync <- true
						return nil
					}, func(e error) {
						hogosuru.AssertErr(e)
						fetchsync <- true
					})

				}

			}
			return nil
		}, func(e error) {

			hogosuru.AssertErr(e)
			fetchsync <- true
		})

	}
	<-fetchsync

	println("------")
	if p, err := fetch.New(endpoint.String()); hogosuru.AssertErr(err) {

		p.Then(func(r response.Response) *promise.Promise {
			if s, err := r.Status(); hogosuru.AssertErr(err) && s == 200 {

				if p2, err := r.Text(); hogosuru.AssertErr(err) {

					p2.Then(func(i interface{}) *promise.Promise {

						if j, err := json.Parse(i.(string)); hogosuru.AssertErr(err) {
							json := j.Map()
							fmt.Printf("Hello3 %s\n", json.(map[string]interface{})["hello"])
						}

						fetchsync <- true
						return nil
					}, nil)

				}

			}
			return nil
		}, func(e error) {

			hogosuru.AssertErr(e)
			fetchsync <- true
		})

	}
	<-fetchsync

	var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
		"XCustomValue": "Test"}

	var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test"}

	//Start promise and wait result
	if f, err := fetch.New(endpoint.String(), fetchOpts); err == nil {

		if RespText, err := f.Then(func(r response.Response) *promise.Promise {

			if s, err := r.Status(); hogosuru.AssertErr(err) && s == 200 {
				var ptext promise.Promise
				ptext, err = r.Text()
				hogosuru.AssertErr(err)

				return &ptext
			} else {
				if err == nil {
					if errorpromise, err2 := promise.Reject(errors.New("Bad HTTP Status")); hogosuru.AssertErr(err2) {
						return &errorpromise
					}
				}
			}

			return nil
		}, nil); hogosuru.AssertErr(err) {

			RespText.Then(func(i interface{}) *promise.Promise {

				if j, err := json.Parse(i.(string)); err == nil {
					json := j.Map()
					fmt.Printf("[Hello POST] %s\n", json.(map[string]interface{})["hello"])
				} else {
					fmt.Printf("erreur %s", err)
				}

				return nil
			}, func(e error) {

				hogosuru.AssertErr(e)
			})

		}

	}

	//test Posting data

	/*
		fetch.Fetch(endpoint.String(), "GET", nil, nil, func(r response.Response, err error) {

				if s, _ := r.Status(); s == 200 {

					if p, err := r.Text(); err == nil {
						p.Then(func(i interface{}) *promise.Promise {


						},func(e error) {


						}),

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

				if header, err := r.Response_().Headers(); hogosuru.AssertErr(err) {
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
	*/
	ch := make(chan struct{})
	<-ch

}
