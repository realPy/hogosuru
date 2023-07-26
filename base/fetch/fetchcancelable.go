package fetch

import (
	"github.com/realPy/hogosuru/base/abortcontroller"
	"github.com/realPy/hogosuru/base/abortsignal"
)

// FetchCancelable struct
type FetchCancelable struct {
	Fetch
	abortctrl *abortcontroller.AbortController
}

func NewCancelable(urlfetch string, opts ...interface{}) (FetchCancelable, error) {

	//var arrayJS []interface{}
	var f FetchCancelable
	var err error
	var init interface{}
	var s abortsignal.AbortSignal
	var abortctrl abortcontroller.AbortController

	if abortctrl, err = abortcontroller.New(); err == nil {
		if s, err = abortctrl.Signal(); err == nil {
			if len(opts) == 0 {

				init = map[string]interface{}{"signal": s.JSObject()}
				f.abortctrl = &abortctrl

			} else {
				if initarray, ok := opts[0].(map[string]interface{}); ok {
					if _, ok := initarray["signal"]; !ok {
						initarray["signal"] = s.JSObject()
						f.abortctrl = &abortctrl
					}
				}
				init = opts[0]
			}

			if init == nil {
				f.Fetch, err = New(urlfetch)
			} else {
				f.Fetch, err = New(urlfetch, init)
			}
		}
	}

	return f, err

}

func (f FetchCancelable) Abort() error {

	if f.abortctrl == nil {
		return ErrSignalNotManaged
	}
	return f.abortctrl.Abort()
}
