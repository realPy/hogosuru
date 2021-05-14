package broadcastchannel

import (
	"testing"
)

func TestChasnnel(t *testing.T) {

	//	var io chan string = make(chan string)
	//	var c1 Channel
	var err error
	if _, err = New("alpha"); err == nil {
		/*
			if c2, err = New("alpha"); err == nil {

				c2.SetOnMessage(func(ch Channel, ev messageevent.MessageEvent) {
					if dataObject, err := ev.Data(); err == nil {
						io <- dataObject.String()
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
				fmt.Printf("receive %s\n", r)

			case <-time.After(time.Duration(100) * time.Millisecond):
				t.Errorf("No message channel receive")
			}
			//c2.Release()
		*/
	} else {
		t.Errorf("Channel Alpha not created %s", err.Error())
	}

}
