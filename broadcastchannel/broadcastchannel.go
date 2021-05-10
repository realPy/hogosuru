package broadcastchannel

import (
	"sync"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/object/event/messageevent"
)

var singleton sync.Once

var bcinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Channel struct
type Channel struct {
	object.Object
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

//New Get a new channel broadcast
func New(channelname string) (Channel, error) {
	var channel Channel

	if bci := GetJSInterface(); bci != nil {
		channel.Object = channel.SetObject(bci.objectInterface.New(js.ValueOf(channelname)))
		return channel, nil
	}
	return channel, ErrNotImplemented
}

//SetReceiveMessage Set the receiver method on channel
func (c Channel) SetReceiveMessage(handler func(Channel, object.GOMap)) {
	onmessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if obj, err := messageevent.NewMessageEvent(args[0]); err == nil {
			handler(c, obj)
		}

		return nil
	})

	c.JSObject().Set("onmessage", onmessage)

}

//PostMessage Post a message on channel
func (c Channel) PostMessage(message string) error {
	var err error
	_, err = c.JSObject().CallWithErr("postMessage", js.ValueOf(message))

	return err
}
