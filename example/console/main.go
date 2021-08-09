package main

import (
	"github.com/realPy/hogosuru/console"
)

func main() {

	if c, err := console.New(); err == nil {
		c.Time("console")
		c.Debug("Debug message")
		c.Dir(c.BaseObject)

		c.Error("error message")
		c.Info("info message :)")

		c.GroupCollapsed("Data collapsed")
		c.Info("message1")
		c.Info("message2")
		c.GroupEnd()

		c.Group("Data no collapsed")
		c.Info("message1")
		c.Info("message2")
		c.Count("countme")
		c.GroupEnd()
		c.Warn("Warn message")
		c.Log("Log message")
		c.TimeEnd("console")
		c.Count("countme")
	} else {
		println("erreur", err.Error())
	}

	ch := make(chan struct{})
	<-ch

}
