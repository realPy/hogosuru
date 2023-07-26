package wasm

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/arraybuffer"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/fetch"
	"github.com/realPy/hogosuru/base/jserror"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/response"
	"github.com/realPy/hogosuru/base/webassembly"
)

//LoadWasm dynamic fetch and start a wasm binary

func LoadWasm_(urlfetch string) (fetch.Fetch, promise.Promise, error) {

	var err error
	var f fetch.Fetch
	var p promise.Promise

	var w webassembly.WebAssembly
	var gobjinterface js.Value

	if w, err = webassembly.New(); err == nil {

		if gobjinterface, err = baseobject.Get(js.Global(), "Go"); err == nil {
			gobj := gobjinterface.New()
			var importobj js.Value

			if importobj, err = baseobject.Get(gobj, "importObject"); err == nil {

				if f, err = fetch.New(urlfetch, map[string]interface{}{"method": "GET"}); err == nil {

					if ok, err := w.Implement("instantiateStreaming"); err == nil && ok {

						if p, err = w.InstantiateStreaming(f.Promise, importobj); err == nil {

							p.Then(func(obj interface{}) *promise.Promise {
								var instance js.Value

								if module, ok := obj.(baseobject.ObjectFrom); ok {

									if instance, err = module.BaseObject_().Get("instance"); err == nil {

										_, err = baseobject.Call(gobj, "run", instance)

									}
								}

								return nil
							}, nil)

						}
					} else {

						p, err = promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

							f.Then(func(r response.Response) *promise.Promise {
								var parr promise.Promise
								var arr arraybuffer.ArrayBuffer
								var ok bool
								var err error

								if parr, err = r.ArrayBuffer(); err == nil {

									parr.Then(func(i interface{}) *promise.Promise {

										if arr, ok = i.(arraybuffer.ArrayBuffer); ok {

											var p1 promise.Promise
											if p1, err = w.Instantiate(arr, importobj); err == nil {
												p1.Then(func(obj interface{}) *promise.Promise {
													var instance js.Value

													if module, ok := obj.(baseobject.ObjectFrom); ok {

														if instance, err = module.BaseObject_().Get("instance"); err == nil {
															if _, err = baseobject.Call(gobj, "run", instance); err == nil {
																resolvefunc.Invoke()
															}

														}
													}
													if err != nil {
														var errjs jserror.JSError
														if errjs, err = jserror.New(err); err == nil {
															errfunc.Invoke(errjs.JSObject())
														} else {
															hogosuru.AssertErr(err)
														}
													}
													return nil
												}, func(e error) {

													if errjs, err := jserror.New(e); err == nil {
														errfunc.Invoke(errjs.JSObject())
													} else {
														hogosuru.AssertErr(err)
													}

												})

											}

										} else {

											if errjs, err := jserror.New(errors.New("Response is not a buffer")); err == nil {
												errfunc.Invoke(errjs.JSObject())
											} else {
												hogosuru.AssertErr(err)
											}

										}

										return nil

									}, func(e error) {
										if errjs, err := jserror.New(e); err == nil {
											errfunc.Invoke(errjs.JSObject())
										} else {
											hogosuru.AssertErr(err)
										}
									})

								}

								if err != nil {
									var errjs jserror.JSError
									if errjs, err = jserror.New(err); err == nil {
										errfunc.Invoke(errjs.JSObject())
									} else {
										hogosuru.AssertErr(err)
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

func LoadWasm(urlfetch string) (fetch.Fetch, promise.Promise, error) {

	var err error
	var f fetch.Fetch
	var p promise.Promise

	var w webassembly.WebAssembly
	var gobjinterface js.Value

	if w, err = webassembly.New(); err == nil {

		if gobjinterface, err = baseobject.Get(js.Global(), "Go"); err == nil {
			gobj := gobjinterface.New()
			var importobj js.Value

			if importobj, err = baseobject.Get(gobj, "importObject"); err == nil {

				if f, err = fetch.New(urlfetch, map[string]interface{}{"method": "GET"}); err == nil {

					if ok, err := w.Implement("instantiateStreaming"); err == nil && ok {

						if p, err = w.InstantiateStreaming(f.Promise, importobj); err == nil {

							p.Then(func(obj interface{}) *promise.Promise {
								var instance js.Value

								if module, ok := obj.(baseobject.ObjectFrom); ok {

									if instance, err = module.BaseObject_().Get("instance"); err == nil {
										_, err = baseobject.Call(gobj, "run", instance)

									}
								}

								return nil
							}, nil)

						}
					} else {

						p, err = promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
							var pab promise.Promise

							pab, err = f.Then(func(r response.Response) *promise.Promise {
								var parr promise.Promise
								var err error
								if parr, err = r.ArrayBuffer(); err != nil {

									parr, err = promise.Reject(err)

								}
								return &parr

							}, nil)
							var pi promise.Promise

							pi, err = pab.Then(func(i interface{}) *promise.Promise {
								var arr arraybuffer.ArrayBuffer
								var p1 promise.Promise

								if arr, ok = i.(arraybuffer.ArrayBuffer); ok {

									p1, err = w.Instantiate(arr, importobj)

								} else {

									p1, err = promise.Reject(err)

								}

								return &p1

							}, nil)

							pi.Then(func(obj interface{}) *promise.Promise {
								var instance js.Value

								if module, ok := obj.(baseobject.ObjectFrom); ok {

									if instance, err = module.BaseObject_().Get("instance"); err == nil {
										if _, err = baseobject.Call(gobj, "run", instance); err == nil {
											resolvefunc.Invoke()
										}

									}
								}
								if err != nil {
									rej, _ := promise.Reject(err)
									return &rej
								}
								return nil
							}, func(e error) {

								if errjs, err := jserror.New(e); err == nil {
									errfunc.Invoke(errjs.JSObject())
								} else {
									hogosuru.AssertErr(err)
								}

							})

							return nil, nil
						})

					}

				}

			}
		}

	}

	return f, p, err
}
