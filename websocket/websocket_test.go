package websocket

import (
	"testing"
	"time"

	"github.com/realPy/hogosuru/messageevent"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestEcho(t *testing.T) {
	var io chan bool = make(chan bool)
	var nbmsg int = 0
	if w, err := New("wss://ws.ifelse.io"); err == nil {

		w.SetOnMessage(func(e messageevent.MessageEvent) {
			if nbmsg == 0 {
				w.Send("hogosuru")
			} else {
				if message, err := e.Data(); err == nil {
					if s, ok := message.(string); ok {
						if s == "hogosuru" {
							io <- true
						} else {
							t.Errorf("Serveur must send hogosuru pong receive %s", s)
						}

					} else {
						t.Error("Response must be a string")
					}

				} else {
					t.Error(err.Error())
				}
			}
			nbmsg++

		})

	} else {
		t.Error(err.Error())
	}
	select {
	case <-io:
	case <-time.After(time.Duration(2000) * time.Millisecond):
		t.Errorf("No message channel receive")
	}
}
