package broadcastchannel

import (
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

type BroadcastchannelInterface struct {
	bcInterface js.Value
}

type Channel struct {
	channelObject js.Value
}

func NewBroadcastChannelInterface() (BroadcastchannelInterface, error) {
	var instance BroadcastchannelInterface
	var err error
	instance.bcInterface, err = js.Global().GetWithErr("BroadcastChannel")

	return instance, err
}

func (bc BroadcastchannelInterface) New(channel string) Channel {
	return Channel{channelObject: bc.bcInterface.New(js.ValueOf(channel))}
}

func (c Channel) SetReceiveMessage(handler func(js.Value)) {
	onmessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if obj, err := object.DataFromMessageEvent(args[0]); err == nil {
			handler(obj)
		}

		return nil
	})

	c.channelObject.Set("onmessage", onmessage)

}

func (c Channel) PostMessage(message string) error {
	var err error
	_, err = c.channelObject.CallWithErr("postMessage", js.ValueOf(message))

	return err
}
