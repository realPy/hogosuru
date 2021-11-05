package baseobject

import "syscall/js"

func GoValue_(object js.Value) interface{} {

	var i interface{}
	var err error

	if i, err = GoValue(object); err != nil {
		debug(err.Error())
	}

	return i
}

func (b BaseObject) Class_() string {

	var c string
	var err error

	if c, err = b.Class(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b BaseObject) ToString_() string {

	var c string
	var err error

	if c, err = b.ToString(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b BaseObject) ConstructName_() string {

	var c string
	var err error

	if c, err = b.ConstructName(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b BaseObject) GetAttributeString_(attribute string) string {
	var c string
	var err error

	if c, err = b.GetAttributeString(attribute); err != nil {
		b.Debug(err.Error())
	}

	return c
}
