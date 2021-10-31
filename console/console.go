package console

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var consoleinterface js.Value

//Console Console struct
type Console struct {
	baseobject.BaseObject
}

type ConsoleFrom interface {
	Console_() Console
}

func (c Console) Console_() Console {
	return c
}

//GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if consoleinterface, err = baseobject.Get(js.Global(), "console"); err != nil {
			consoleinterface = js.Undefined()
		}

		baseobject.Register(consoleinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return consoleinterface
}

func New() (Console, error) {

	var c Console
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		c.BaseObject = c.SetObject(di)

	} else {

		err = ErrNotImplemented
	}

	return c, err
}
func NewFromJSObject(obj js.Value) (Console, error) {
	var c Console
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				c.BaseObject = c.SetObject(obj)

			} else {
				err = ErrNotAConsole
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return c, err
}

func (c Console) Assert(assertion bool, opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	arrayJS = append(arrayJS, js.ValueOf(assertion))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("assert", arrayJS...)
	return err
}

func (c Console) Clear() error {
	var err error
	_, err = c.Call("clear")
	return err
}

func (c Console) Count(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	_, err = c.Call("count", arrayJS...)
	return err
}

func (c Console) CountReset(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	_, err = c.Call("countReset", arrayJS...)
	return err
}

func (c Console) Debug(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("debug", arrayJS...)
	return err
}

func (c Console) Dir(obj baseobject.BaseObject) error {

	var err error
	_, err = c.Call("dir", obj.JSObject())
	return err
}

func (c Console) DirXml(obj baseobject.BaseObject) error {

	var err error
	_, err = c.Call("dirxml", obj.JSObject())
	return err
}

func (c Console) Error(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("error", arrayJS...)
	return err
}

func (c Console) Exception(opts ...interface{}) error {

	return c.Error(opts...)
}

func (c Console) Group(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	_, err = c.Call("group", arrayJS...)
	return err
}

func (c Console) GroupCollapsed(label ...string) error {

	var err error
	var arrayJS []interface{}

	if len(label) > 0 {
		arrayJS = append(arrayJS, label[0])
	}

	_, err = c.Call("groupCollapsed", arrayJS...)
	return err
}

func (c Console) GroupEnd() error {

	var err error
	_, err = c.Call("groupEnd")
	return err
}

func (c Console) Info(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("info", arrayJS...)
	return err
}

func (c Console) Log(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("log", arrayJS...)
	return err
}

func (c Console) Time(label string) error {

	var err error
	_, err = c.Call("time", js.ValueOf(label))
	return err
}

func (c Console) TimeEnd(label string) error {

	var err error
	_, err = c.Call("timeEnd", js.ValueOf(label))
	return err
}

func (c Console) TimeLog(label string) error {

	var err error
	_, err = c.Call("timeLog", js.ValueOf(label))
	return err
}

func (c Console) Trace() error {

	var err error
	_, err = c.Call("trace")
	return err
}

func (c Console) Warn(opts ...interface{}) error {

	var arrayJS []interface{}
	var err error

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	_, err = c.Call("warn", arrayJS...)
	return err
}
