package broadcastchannel

//Full implemented
// https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/messageevent"
)

var singleton sync.Once

var bcinterface js.Value

//BroadcastChannel struct
type BroadcastChannel struct {
	eventtarget.EventTarget
}

type BroadcastChannelFrom interface {
	BroadcastChannel() BroadcastChannel
}

func (b BroadcastChannel) BroadcastChannel() BroadcastChannel {
	return b
}

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if bcinterface, err = js.Global().GetWithErr("BroadcastChannel"); err != nil {
			bcinterface = js.Null()
		}

	})
	//for autodiscover
	messageevent.GetInterface()

	return bcinterface
}

//New Get a new channel broadcast
func New(channelname string) (BroadcastChannel, error) {
	var channel BroadcastChannel
	var err error
	if bci := GetInterface(); !bci.IsNull() {
		channel.BaseObject = channel.SetObject(bci.New(js.ValueOf(channelname)))
	} else {
		err = ErrNotImplemented
	}
	return channel, err
}

//PostMessage Post a message on channel
func (c BroadcastChannel) PostMessage(message string) error {
	var err error
	_, err = c.JSObject().CallWithErr("postMessage", js.ValueOf(message))

	return err
}

//Close Close the channel
func (c BroadcastChannel) Close() error {
	var err error
	_, err = c.JSObject().CallWithErr("close")

	return err
}

func (c BroadcastChannel) Name() (string, error) {

	return c.GetAttributeString("name")
}
