package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
)

func main() {
	hogosuru.Init()
	d := document.New_()

	body := d.Body_()

	if div1, err := htmldivelement.New(d); hogosuru.AssertErr(err) {

		body.AppendChild(div1.Node)
		hogosuru.KeyObservable().RegisterFunc("prenom", func(value interface{}) {

			if prenom, ok := value.(string); ok {
				div1.SetTextContent(prenom)
			}

		})

	}

	if div2, err := htmldivelement.New(d); hogosuru.AssertErr(err) {

		body.AppendChild(div2.Node)
		hogosuru.KeyObservable().RegisterFunc("prenom", func(value interface{}) {

			if prenom, ok := value.(string); ok {
				div2.SetTextContent(prenom)
			}

		})

	}

	hogosuru.KeyObservable().Set("prenom", "manu", true)

	if _, err := hogosuru.KeyObservable().Get("nom"); hogosuru.AssertErr(err) {
		println("Ok->")
	}
	ch := make(chan struct{})
	<-ch

}
