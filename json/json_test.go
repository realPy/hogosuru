package json

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
func TestParse(t *testing.T) {

	var str = `{
		"name":"John",
		"age":30,
		"cars":[ "Ford", "BMW", "Fiat" ]
		}`

	var badstr = `{
			"name":"John",
			"age":30,
			"cars:[ "Ford", "BMW", "Fiat" ]
			}`

	if j, err := Parse(str); err == nil {
		goValue := j.Map()

		if goValue.(map[string]interface{})["name"] != "John" {
			t.Errorf("Name not match")
		}
	} else {
		t.Errorf(err.Error())
	}

	if _, err := Parse(badstr); err == nil {
		t.Error("Must give an error")
	}
}
func TestStringify(t *testing.T) {

	if str, err := Stringify(1, "hello", true); err == nil {

		if str != "[1,\"hello\",true]" {
			t.Errorf("Value not match %s", str)
		}

	} else {
		t.Errorf(err.Error())
	}
}
