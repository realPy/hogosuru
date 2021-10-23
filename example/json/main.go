package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/json"
)

var str2 = `{
	"name":"John",
	"age":30,
	"cars":[ "Ford", "BMW", "Fiat" ]
	}`

func main() {
	hogosuru.Init()
	//var str string = "{\"test\":true,\"o\":\"poi\",\"nani\":1.5}"
	//var str string = "{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"

	if j, err := json.Parse(str2); err == nil {

		p := j.Map()

		println(p.(map[string]interface{})["cars"].([]interface{})[0])

		j.Export("adrien")
	} else {
		println("erreur " + err.Error())

	}

	ch := make(chan struct{})
	<-ch

}
