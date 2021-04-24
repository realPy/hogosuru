package main

import (
	"fmt"

	"github.com/realPy/jswasm"
)

func main() {

	if j, err := jswasm.JsonParse("{\"test\":true,\"o\":\"poi\",\"nani\":1.5,\"complex\":{ \"toto\":\"yes\"}}"); err == nil {
		p := j.GoJson()
		fmt.Printf("Value of complex[\"toto\"] %s\n", p["complex"].(map[string]interface{})["toto"])
	} else {
		fmt.Printf("erreur %s", err)
	}

	ch := make(chan struct{})
	<-ch

}
