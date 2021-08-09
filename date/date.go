package date

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var dateinterface js.Value

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if dateinterface, err = js.Global().GetWithErr("Date"); err != nil {
			dateinterface = js.Null()
		}
	})
	baseobject.Register(dateinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return dateinterface
}

//Date struct
type Date struct {
	object.Object
}

type DateFrom interface {
	Date() Date
}

func (d Date) Date() Date {
	return d
}

func New(values ...interface{}) (Date, error) {
	var d Date
	var err error
	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}
	if di := GetInterface(); !di.IsNull() {

		d.BaseObject = d.SetObject(di.New(arrayJS...))

	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (Date, error) {
	var d Date
	var err error

	if di := GetInterface(); !di.IsNull() {
		if obj.InstanceOf(di) {
			d.BaseObject = d.SetObject(obj)
		} else {
			err = ErrNotADate
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d Date) callString(method string, opts ...interface{}) (string, error) {

	var err error
	var obj js.Value
	var ret string

	var optJSValue []interface{}

	for _, opt := range opts {
		optJSValue = append(optJSValue, js.ValueOf(opt))
	}

	if obj, err = d.JSObject().CallWithErr(method, optJSValue...); err == nil {
		if obj.Type() == js.TypeString {
			ret = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return ret, err
}

func (d Date) GetDate() (int64, error) {
	return d.CallInt64("getDate")
}

func (d Date) GetDay() (int64, error) {
	return d.CallInt64("getDay")
}

func (d Date) GetFullYear() (int64, error) {
	return d.CallInt64("getFullYear")
}

func (d Date) GetHours() (int64, error) {
	return d.CallInt64("getHours")
}

func (d Date) GetMilliseconds() (int64, error) {
	return d.CallInt64("getMilliseconds")
}

func (d Date) GetMinutes() (int64, error) {
	return d.CallInt64("getMinutes")
}

func (d Date) GetSeconds() (int64, error) {
	return d.CallInt64("getSeconds")
}

func (d Date) GetTime() (int64, error) {
	return d.CallInt64("getTime")
}

func (d Date) GetTimezoneOffset() (int64, error) {
	return d.CallInt64("getTimezoneOffset")
}

func (d Date) GetUTCDate() (int64, error) {
	return d.CallInt64("getUTCDate")
}

func (d Date) GetUTCDay() (int64, error) {
	return d.CallInt64("getUTCDay")
}

func (d Date) GetUTCFullYear() (int64, error) {
	return d.CallInt64("getUTCFullYear")
}

func (d Date) GetUTCHours() (int64, error) {
	return d.CallInt64("getUTCHours")
}

func (d Date) GetUTCMilliseconds() (int64, error) {
	return d.CallInt64("getUTCMilliseconds")
}

func (d Date) GetUTCMinutes() (int64, error) {
	return d.CallInt64("getUTCMinutes")
}

func (d Date) GetUTCMonth() (int64, error) {
	return d.CallInt64("getUTCMonth")
}

func (d Date) GetUTCSeconds() (int64, error) {
	return d.CallInt64("getUTCSeconds")
}

func (d Date) SetDate(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setDate", js.ValueOf(value))
	return err
}

func (d Date) SetFullYear(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setFullYear", js.ValueOf(value))
	return err
}

func (d Date) SetHours(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setHours", js.ValueOf(value))
	return err
}

func (d Date) SetMilliseconds(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setMilliseconds", js.ValueOf(value))
	return err
}

func (d Date) SetMinutes(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setMinutes", js.ValueOf(value))
	return err
}

func (d Date) SetSeconds(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setSeconds", js.ValueOf(value))
	return err
}

func (d Date) SetTime(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setTime", js.ValueOf(value))
	return err
}

func (d Date) SetUTCDate(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCDate", js.ValueOf(value))
	return err
}

func (d Date) SetUTCFullYear(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCFullYear", js.ValueOf(value))
	return err
}

func (d Date) SetUTCHours(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCHours", js.ValueOf(value))
	return err
}

func (d Date) SetUTCMilliseconds(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCMilliseconds", js.ValueOf(value))
	return err
}

func (d Date) SetUTCMinutes(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCMinutes", js.ValueOf(value))
	return err
}

func (d Date) SetUTCSeconds(value int64) error {
	var err error

	_, err = d.JSObject().CallWithErr("setUTCSeconds", js.ValueOf(value))
	return err
}

func Parse(value string) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if di := GetInterface(); !di.IsNull() {

		if obj, err = di.CallWithErr("parse", js.ValueOf(value)); err == nil {
			if obj.Type() == js.TypeNumber {
				ret = int64(obj.Float())
			} else {
				err = baseobject.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func Now() (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	if di := GetInterface(); !di.IsNull() {

		if obj, err = di.CallWithErr("now"); err == nil {
			if obj.Type() == js.TypeNumber {
				ret = int64(obj.Float())
			} else {
				err = baseobject.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func (d Date) ToDateString() (string, error) {
	return d.callString("toDateString")
}

func (d Date) ToISOString() (string, error) {
	return d.callString("toISOString")
}

func (d Date) ToJSON() (string, error) {
	return d.callString("toJSON")
}

func (d Date) ToLocaleDateString(locale string, options map[string]interface{}) (string, error) {

	return d.callString("toLocaleDateString", locale, options)
}

func (d Date) ToLocaleString(locale string, options map[string]interface{}) (string, error) {

	return d.callString("toLocaleString", locale, options)
}

func (d Date) ToLocaleTimeString(locale string, options map[string]interface{}) (string, error) {

	return d.callString("toLocaleTimeString", locale, options)
}

func (d Date) ToString() (string, error) {
	return d.callString("toString")
}

func (d Date) ToTimeString() (string, error) {
	return d.callString("toTimeString")
}

func (d Date) ToUTCString() (string, error) {
	return d.callString("toUTCString")
}

func (d Date) ValueOf() (int64, error) {
	return d.CallInt64("valueOf")
}

func UTC(values ...interface{}) (int64, error) {
	var err error
	var obj js.Value
	var ret int64
	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}

	if di := GetInterface(); !di.IsNull() {

		if obj, err = di.CallWithErr("UTC", arrayJS...); err == nil {
			if obj.Type() == js.TypeNumber {
				ret = int64(obj.Float())
			} else {
				err = baseobject.ErrObjectNotNumber
			}
		}
		return ret, err

	} else {
		err = ErrNotImplemented
	}

	return ret, err
}
