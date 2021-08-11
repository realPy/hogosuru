package hogosuru

import (
	"syscall/js"

	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/webassembly"
)

//LoadWasm dynamic fetch and start a wasm binary

func LoadWasm(urlfetch string) (promise.Promise, error) {

	return promise.New(func(p promise.Promise) (interface{}, error) {

		var err error
		var w webassembly.WebAssembly
		var gobjinterface js.Value

		if w, err = webassembly.New(); err == nil {

			if gobjinterface, err = js.Global().GetWithErr("Go"); err == nil {
				gobj := gobjinterface.New()
				var importobj js.Value
				if importobj, err = gobj.GetWithErr("importObject"); err == nil {
					var f fetch.Fetch

					if f, err = fetch.NewFetch(urlfetch, "GET", nil, nil, nil); err == nil {
						var p promise.Promise

						if p, err = w.InstantiateStreaming(f.Promise, importobj); err == nil {
							module, _ := p.Await()
							var instance js.Value

							if instance, err = module.JSObject().GetWithErr("instance"); err == nil {
								_, err = gobj.JSValue().CallWithErr("run", instance)
							}

						}

					}

				}
			}

		}
		return nil, err
	})
}
