package filelist

func (f FileList) Length_() int {
	var i int
	var err error

	if i, err = f.Length(); err != nil {
		f.Debug(err.Error())
	}
	return i
}
