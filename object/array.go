package object

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type GOArray struct {
	value []GOValue
}

func Array(object js.Value) GOArray {

	var goArray GOArray

	if object.Type() == js.TypeObject {
		for i := 0; i < object.Length(); i++ {
			jsvalue := object.Index(i)
			goArray.value = append(goArray.value, NewGOValue(jsvalue))
		}

	}
	return goArray
}

func ParseArray(object js.Value, handler func(js.Value)) error {

	if object.Type() == js.TypeObject {
		for i := 0; i < object.Length(); i++ {

			handler(object.Index(i))
		}
	}

	return ErrNotAnObject
}
func (g GOArray) Length() int {

	return len(g.value)
}

func (g GOArray) Index(i int) GOValue {

	return g.value[i]
}

func (g GOArray) String() string {

	var str string = "["
	var i int = 0
	var l int = len(g.value)
	for _, value := range g.value {
		str = fmt.Sprintf("%s%s", str, value.String())
		if i < (l - 1) {
			str = str + " "
		}
		i++
	}
	str = fmt.Sprintf("%s]", str)
	return str
}
