package date

func New_(values ...interface{}) Date {
	d, err := New(values...)

	if err != nil {
		d.Debug(err.Error())
	}
	return d
}

func (d Date) ToISOString_() string {

	var s string
	var err error

	if s, err = d.ToISOString(); err != nil {
		d.Debug(err.Error())
	}

	return s
}
