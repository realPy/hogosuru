package testingutils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
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

func AssertErr(t *testing.T, err error, skips ...int) bool {

	if err != nil {

		skip := 1
		if len(skips) > 0 {
			skip = skips[0]
		}
		_, file, line, _ := runtime.Caller(skip)

		t.Errorf("%s:%d >> %s", filepath.Base(file), line, err.Error())
		return false
	}

	return true
}

func AssertExpect(t *testing.T, exp interface{}, get interface{}, skips ...int) bool {

	if !reflect.DeepEqual(exp, get) {
		skip := 1
		if len(skips) > 0 {
			skip = skips[0]
		}

		_, file, line, _ := runtime.Caller(skip)
		prefix := fmt.Sprintf("%s:%d >> ", filepath.Base(file), line)
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

func AssertStringContains(t *testing.T, exp interface{}, get interface{}) bool {

	var aType = "<nil>"
	var bType = "<nil>"

	if reflect.ValueOf(exp).IsValid() {
		aType = reflect.TypeOf(exp).Name()
	}
	if reflect.ValueOf(get).IsValid() {
		bType = reflect.TypeOf(get).Name()
	}

	if _, ok := exp.(string); !ok {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d >> Expect value must be a atring", filepath.Base(file), line)
	}

	if aType != bType {
		_, file, line, _ := runtime.Caller(1)

		t.Errorf("%s:%d >>  Expect %v Have %v ", filepath.Base(file), line, aType, bType)
		return false
	}

	if !strings.Contains(get.(string), exp.(string)) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d >> %s not contains in %s", filepath.Base(file), line, exp.(string), get.(string))
		return false
	}

	return true
}
