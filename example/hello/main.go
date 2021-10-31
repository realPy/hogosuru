package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmlheadingelement"
)

func main() {
	hogosuru.Init()
	//we get the current document if an error occur the err is draw to the console thank to AssertErr
	if doc, err := document.New(); hogosuru.AssertErr(err) {

		//we get the body of the document if an error occur the err is draw to the console thank to AssertErr
		if body, err := doc.Body(); hogosuru.AssertErr(err) {

			//now we create dynamiclly the h1 element
			if h1, err := htmlheadingelement.NewH1(doc); hogosuru.AssertErr(err) {

				//We set the text content with Hello World
				h1.SetTextContent("Hello world")
				// and append to the body

				body.AppendChild(h1.Node)

			}

		}
	}

	ch := make(chan struct{})
	<-ch

}
