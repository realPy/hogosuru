package testingutils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func ImplementedExpect(t *testing.T, obj baseobject.BaseObject, methodsname []string) {

	for _, methodname := range methodsname {

		if ok, err := obj.Implement(methodname); AssertErr(t, err) {

			if !ok {
				t.Errorf("%s must implemented %s", obj.ToString_(), methodname)
			}

		}

	}

}

func AssertErr(t *testing.T, err error) bool {

	if err != nil {
		t.Errorf(err.Error())
		return false
	}

	return true
}

func AssertExpect(t *testing.T, exp interface{}, get interface{}) bool {

	if !reflect.DeepEqual(exp, get) {
		_, file, line, _ := runtime.Caller(1)
		prefix := fmt.Sprintf("%s:%d >>", filepath.Base(file), line)
		var aType = "<nil>"
		var bType = "<nil>"
		if reflect.ValueOf(exp).IsValid() {
			aType = reflect.TypeOf(exp).Name()
		}
		if reflect.ValueOf(get).IsValid() {
			bType = reflect.TypeOf(get).Name()
		}
		t.Errorf("%s Expect %s:%v Have %s:%v", prefix, aType, exp, bType, get)
		return false
	}

	return true
}
