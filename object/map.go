package object

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type GOMap struct {
	value map[string]GOValue
}

func (g GOMap) Get(key string) GOValue {
	return g.value[key]
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
		if Object, err := NewObject(); err == nil {
			if entries, err := Object.Entries(object); err == nil {
				for i := 0; i < entries.Length(); i++ {

					keypair := entries.Index(i)
					if key, value := Pair(keypair); true {
						if value.IsObject() {
							m[key.String()] = GOValue{value: Map(value.Object())}
						} else {
							m[key.String()] = value
						}

					}

				}
			}
		}

	}
	return GOMap{value: m}
}
