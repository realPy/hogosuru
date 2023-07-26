package file

func (f File) Name_() string {
	var c string
	var err error

	if c, err = f.Name(); err != nil {
		f.Debug(err.Error())
	}
	return c
}

func (f File) Type_() string {
	var c string
	var err error

	if c, err = f.Type(); err != nil {
		f.Debug(err.Error())
	}
	return c
}
