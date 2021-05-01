package object

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type GOMap struct {
	value map[string]GOValue
}

func (g GOMap) String() string {

	var str string = "["
	var i int = 0
	var l int = len(g.value)
	for key, value := range g.value {
		str = fmt.Sprintf("%s%s:%s", str, key, value.String())
		if i < (l - 1) {
			str = str + " "
		}
		i++
	}
	str = fmt.Sprintf("%s]", str)
	return str
}

func Map(object js.Value) GOMap {

	var m map[string]GOValue = make(map[string]GOValue)
	if object.Type() == js.TypeObject {
		for i := 0; i < object.Length(); i++ {
			keypair := object.Index(i)
			if key, value := Pair(keypair); true {
				m[key.String()] = value
			}

		}

	}
	return GOMap{value: m}
}
