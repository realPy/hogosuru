package element

func (e Element) ID_() string {

	var c string
	var err error

	if c, err = e.ID(); err != nil {
		e.Debug(err.Error())
	}

	return c
}
