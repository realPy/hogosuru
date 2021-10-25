package broadcastchannel

/*
func TestChannelIO(t *testing.T) {

	var io chan string = make(chan string)
	var c1, c2 BroadcastChannel
	var err error

	if c1, err = New("alpha"); err == nil {

		if c2, err = New("alpha"); err == nil {

			c2.OnMessage(func(ev messageevent.MessageEvent) {
				if dataObject, err := ev.Data(); err == nil {
					io <- dataObject.(baseobject.ObjectFrom).BaseObject_().String()
				} else {
					t.Errorf(err.Error())
				}

			})
		} else {
			t.Errorf("Channel Alpha not created")
		}
		c1.PostMessage("hello")

		select {
		case r := <-io:
			if r != "hello" {
				t.Errorf("Must receive hello")
			}

		case <-time.After(time.Duration(100) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
		if err := c2.Close(); err != nil {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf("Channel Alpha not created %s", err.Error())
	}
	if err := c1.Close(); err != nil {
		t.Errorf(err.Error())
	}
}
*/
