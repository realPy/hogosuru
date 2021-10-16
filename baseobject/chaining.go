package baseobject

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

func (b BaseObject) GetAttributeString_(attribute string) string {
	var c string
	var err error

	if c, err = b.GetAttributeString(attribute); err != nil {
		b.Debug(err.Error())
	}

	return c
}
