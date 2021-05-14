package broadcastchannel

//Full implemented
// https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/messageevent"
	"github.com/realPy/hogosuru/object"
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
	onmessage js.Func
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
func (c Channel) SetOnMessage(handler func(Channel, messageevent.MessageEvent)) {
	c.onmessage = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if msgEvent, err := messageevent.NewFromJSObject(args[0]); err == nil {
			handler(c, msgEvent)
		}

		return nil
	})

	c.JSObject().Set("onmessage", c.onmessage)

}

//SetOnError Set the receiver method on channel
func (c Channel) SetOnMessageError(handler func(Channel, messageevent.MessageEvent)) {
	onmessageerror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if msgEvent, err := messageevent.NewFromJSObject(args[0]); err == nil {
			handler(c, msgEvent)
		}

		return nil
	})

	c.JSObject().Set("onmessageerror", onmessageerror)

}

//PostMessage Post a message on channel
func (c Channel) PostMessage(message string) error {
	var err error
	_, err = c.JSObject().CallWithErr("postMessage", js.ValueOf(message))

	return err
}

//Close Close the channel
func (c Channel) Close() error {
	var err error
	_, err = c.JSObject().CallWithErr("close")

	return err
}

func (c Channel) Name() (string, error) {
	var err error
	var obj js.Value

	if obj, err = c.JSObject().GetWithErr("name"); err == nil {

		return obj.String(), nil
	}
	return "", err
}
