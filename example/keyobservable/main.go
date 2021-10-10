package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
)

func main() {

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

	hogosuru.KeyObservable().Send("prenom", "manu", true)

	ch := make(chan struct{})
	<-ch

}
