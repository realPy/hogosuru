package broadcastchannel

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var singleton sync.Once

var bcinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Channel struct
type Channel struct {
	object js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var bcinstance JSInterface
		var err error
		if bcinstance.objectInterface, err = js.Global().GetWithErr("BroadcastChannel"); err == nil {
			bcinterface = &bcinstance
		}
	})

	return bcinterface
}

//NewBroadcastChannel Get a new channel broadcast
func NewBroadcastChannel(channelname string) (Channel, error) {
	var channel Channel

	if bci := GetJSInterface(); bci != nil {
		channel.object = bci.objectInterface.New(js.ValueOf(channelname))
		return channel, nil
	}
	return channel, ErrNotImplemented
}

//SetReceiveMessage Set the receiver method on channel
func (c Channel) SetReceiveMessage(handler func(js.Value)) {
	onmessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if obj, err := object.DataFromMessageEvent(args[0]); err == nil {
			handler(obj)
		}

		return nil
	})

	c.object.Set("onmessage", onmessage)

}

//PostMessage Post a message on channel
func (c Channel) PostMessage(message string) error {
	var err error
	_, err = c.object.CallWithErr("postMessage", js.ValueOf(message))

	return err
}
