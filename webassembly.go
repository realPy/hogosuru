package hogosuru

import (
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/webassembly"
)

//LoadWasm dynamic fetch and start a wasm binary

func LoadWasm(urlfetch string) (fetch.Fetch, promise.Promise, error) {

	var err error
	var f fetch.Fetch
	var p promise.Promise

	var w webassembly.WebAssembly
	var gobjinterface js.Value

	if w, err = webassembly.New(); err == nil {

		if gobjinterface, err = js.Global().GetWithErr("Go"); err == nil {
			gobj := gobjinterface.New()
			var importobj js.Value
			if importobj, err = gobj.GetWithErr("importObject"); err == nil {

				if f, err = fetch.New(urlfetch, map[string]interface{}{"method": "GET"}); err == nil {
					if p, err = w.InstantiateStreaming(f.Promise, importobj); err == nil {

						p.Then(func(obj interface{}) *promise.Promise {
							var instance js.Value

							if module, ok := obj.(baseobject.ObjectFrom); ok {

								if instance, err = module.JSObject().GetWithErr("instance"); err == nil {
									_, err = gobj.JSValue().CallWithErr("run", instance)
								}
							}

							return nil
						}, nil)

					}
				}

			}
		}

	}

	return f, p, err
}
