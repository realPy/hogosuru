package main

import "github.com/realPy/jswasm/json"

func main() {

	if j, err := json.Parse("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		println("Value of complex[\"toto\"]" + p.Get("complex").Get("toto").String())
		println("---->" + p.String())
	} else {
		println("erreur " + err.Error())

	}
	ch := make(chan struct{})
	<-ch

}
