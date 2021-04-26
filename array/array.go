package array

import "github.com/realPy/jswasm/js"

func GoArray(array js.Value) []interface{} {
	var goArray []interface{}

	for i := 0; i < array.Length(); i++ {
		jsvalue := array.Index(i)
		switch jsvalue.Type() {
		case js.TypeNumber:
			goArray = append(goArray, jsvalue.Float())
		case js.TypeString:
			goArray = append(goArray, jsvalue.String())
		case js.TypeBoolean:
			goArray = append(goArray, jsvalue.Bool())

		}

	}
	return goArray
}

func GoArrayInt(array js.Value) []int {
	var goArray []int

	for i := 0; i < array.Length(); i++ {
		jsvalue := array.Index(i)
		if jsvalue.Type() == js.TypeNumber {
			goArray = append(goArray, jsvalue.Int())
		}
	}
	return goArray
}
