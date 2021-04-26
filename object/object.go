package object

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

//String return the string javascript value represent the object
func String(object js.Value) (string, error) {

	if object.Type() == js.TypeObject {
		if value, err := object.CallWithErr("toString"); err == nil {
			return value.String(), nil
		} else {
			return "", err
		}

	}

	return "", fmt.Errorf("The given value must be an object")
}
