package fetch

import (
	"github.com/realPy/hogosuru/abortcontroller"
	"github.com/realPy/hogosuru/abortsignal"
)

//FetchCancellable struct
type FetchCancellable struct {
	Fetch
	abortctrl abortcontroller.AbortController
}

func NewCancellable(urlfetch string, opts ...interface{}) (FetchCancellable, error) {

	//var arrayJS []interface{}
	var f FetchCancellable
	var err error
	var init interface{}
	var s abortsignal.AbortSignal

	if f.abortctrl, err = abortcontroller.New(); err == nil {
		if s, err = f.abortctrl.Signal(); err == nil {
			if len(opts) == 0 {

				init = map[string]interface{}{"signal": s.JSObject()}

			} else {
				if initarray, ok := opts[0].(map[string]interface{}); ok {
					if _, ok := initarray["signal"]; !ok {
						initarray["signal"] = s.JSObject()

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

func (f FetchCancellable) Abort() error {
	return f.abortctrl.Abort()
}
