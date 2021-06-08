package main

import "github.com/realPy/hogosuru/date"

func main() {
	if d, err := date.New(); err == nil {
		value, _ := d.GetMilliseconds()
		println("-->", value)

	} else {
		println("-->", err.Error())
	}

	value, _ := date.UTC(2012, 11, 20, 3, 0, 0)
	if d2, err := date.New(value); err != nil {
		println("erreur", err.Error())
	} else {
		d2.Export("oto")
		if t, _ := d2.ValueOf(); t == value {
			println("Ok")
		}
	}

	println("------>", value)
	d1, _ := date.New()
	ret, _ := d1.ToLocaleString("en-GB", map[string]interface{}{"timeZone": "UTC"})
	println(ret)

	ch := make(chan struct{})
	<-ch

}
