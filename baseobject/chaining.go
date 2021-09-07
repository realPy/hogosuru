package baseobject

func (o BaseObject) Class_() string {

	var c string
	var err error

	if c, err = o.Class(); err != nil {
		o.Debug(err.Error())
	}

	return c
}
