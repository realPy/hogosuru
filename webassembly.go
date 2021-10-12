package hogosuru

import (
	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
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

					if ok, err := w.Implement("instantiateStreaming"); err == nil && ok {

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
					} else {

						p, err = promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

							f.Then(func(r response.Response) *promise.Promise {
								var arr arraybuffer.ArrayBuffer
								var err error
								if arr, err = r.ArrayBuffer(); err == nil {
									var p1 promise.Promise
									if p1, err = w.Instantiate(arr, importobj); err == nil {
										p1.Then(func(obj interface{}) *promise.Promise {
											var instance js.Value

											if module, ok := obj.(baseobject.ObjectFrom); ok {

												if instance, err = module.JSObject().GetWithErr("instance"); err == nil {
													if _, err = gobj.JSValue().CallWithErr("run", instance); err == nil {
														resolvefunc.Invoke()
													}

												}
											}
											if err != nil {
												var errjs js.Value
												if errjs, err = baseobject.ErrorToJS(err); err == nil {
													errfunc.Invoke(errjs)
												} else {
													AssertErr(err)
												}
											}
											return nil
										}, func(e error) {

											if errjs, err := baseobject.ErrorToJS(e); err == nil {
												errfunc.Invoke(errjs)
											} else {
												AssertErr(err)
											}

										})

									}

								}

								if err != nil {
									var errjs js.Value
									if errjs, err = baseobject.ErrorToJS(err); err == nil {
										errfunc.Invoke(errjs)
									} else {
										AssertErr(err)
									}
								}

								return nil
							}, nil)

							return nil, nil
						})

					}

				}

			}
		}

	}

	return f, p, err
}
