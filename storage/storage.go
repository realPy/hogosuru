package storage

import (
	"errors"
	"reflect"
	"strings"
	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/gomap"
	"github.com/realPy/hogosuru/base/indexeddb"
	"github.com/realPy/hogosuru/base/jserror"
	"github.com/realPy/hogosuru/base/object"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/window"
)

/*
	type Pattern struct {
		Id     int      `indexeddb:"store=pattern,version=1,keypath=id,autoincrement"`
		Color  string   `indexeddb:"index=indexcolor,unique"`
		Color2 string   `indexeddb:"index=color2:indexcolor2"`
		Color3 []string `indexeddb:"index=indexcolor3,multientry"`
		Value  string   `indexeddb:""`
	}
*/

var storeFactory = map[string]reflect.Value{}
var stores = map[string]indexeddb.IDBObjectStore{}
var keyFactory = map[string]string{}

/*
func objectStoreDescByType(t reflect.Type) (version string, name string, keypath string, fieldname string, autoincrement bool) {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("indexeddb")
		stag := strings.Split(tag, ",")
		for i := 0; i < len(stag); i++ {
			sstag := strings.Split(stag[i], "=")

			if len(sstag) > 1 {
				switch sstag[0] {
				case "version":
					version = sstag[1]
				case "store":
					name = sstag[1]
				case "keypath":
					keypath = sstag[1]
					fieldname = field.Name
				}

			} else {
				switch stag[i] {
				case "autoincrement":
					autoincrement = true
				}
			}

		}

	}
	if version == "" {
		version = "1"
	}
	return
}
*/

func objectStoreDesc(s interface{}) (reflectname string, version string, name string, keypath string, fieldname string, autoincrement bool) {
	reflectStruct := reflect.ValueOf(s)
	reflectname = reflectStruct.Type().String()

	if _, ok := storeFactory[reflectname]; !ok {
		storeFactory[reflectname] = reflectStruct
	}
	for i := 0; i < reflectStruct.NumField(); i++ {
		field := reflectStruct.Type().Field(i)
		tag := field.Tag.Get("indexeddb")
		stag := strings.Split(tag, ",")
		for i := 0; i < len(stag); i++ {
			sstag := strings.Split(stag[i], "=")

			if len(sstag) > 1 {
				switch sstag[0] {
				case "version":
					version = sstag[1]
				case "store":
					name = sstag[1]
				case "keypath":
					keypath = sstag[1]
					fieldname = field.Name
				}

			} else {
				switch stag[i] {
				case "autoincrement":
					autoincrement = true
				}
			}

		}

	}
	if version == "" {
		version = "1"
	}
	return
}

func objectStoreForeachIndex(s interface{}, f func(nameindex, key string, unique, multientry bool)) {
	var (
		nameindex, key string
		unique         bool
		multientry     bool
	)
	reflectStruct := reflect.ValueOf(s)
	for i := 0; i < reflectStruct.NumField(); i++ {
		field := reflectStruct.Type().Field(i)
		tag := field.Tag.Get("indexeddb")
		stag := strings.Split(tag, ",")
		for i := 0; i < len(stag); i++ {
			sstag := strings.Split(stag[i], "=")
			if len(sstag) > 1 {
				switch sstag[0] {
				case "index":
					ssstag := strings.Split(sstag[1], ":")
					if len(ssstag) > 1 {
						//first is indexname and second key
						nameindex = ssstag[1]
						key = ssstag[0]

					} else {
						key = strings.ToLower(field.Name)
						nameindex = sstag[1]
					}
				}

			} else {
				switch stag[i] {
				case "unique":
					unique = true
				case "multientry":
					multientry = true
				}
			}

		}

		if key != "" {
			f(nameindex, key, unique, multientry)
		}
		key = ""
		nameindex = ""
		unique = false
		multientry = false
	}
}

func Register(w window.Window, data interface{}, databasename string) (promise.Promise, error) {

	//check if the schema is same as stored, if no update schema or create one
	factory, err := w.IndexdedDB()
	if err != nil {
		return promise.Promise{}, err
	}
	reflectname, version, objectname, keypath, fieldname, autoincrement := objectStoreDesc(data)

	openrequest, err := factory.Open(databasename, version)
	if err != nil {
		return promise.Promise{}, err
	}

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		openrequest.OnSuccess(func(e event.Event) {
			result, _ := openrequest.Result()
			if db, ok := result.(indexeddb.IDBDatabase); ok {
				if transaction, err := db.Transaction(objectname, "readwrite"); err == nil {
					if store, err := transaction.ObjectStore(objectname); err == nil {
						stores[reflectname] = store
						keyFactory[reflectname] = fieldname
						resolvefunc.Invoke(store.JSObject())
					}
				}

			} else {
				if errjs, err := jserror.New(errors.New("result open request is not an indexeddb object")); err == nil {
					errfunc.Invoke(errjs.JSObject())
				}
			}

		})

		openrequest.OnUpgradeNeeded(func(e event.Event) {
			result, _ := openrequest.Result()
			desc := map[string]interface{}{"keyPath": keypath}
			if autoincrement {
				desc["autoIncrement"] = true
			}

			if db, ok := result.(indexeddb.IDBDatabase); ok {
				if store, err := db.CreateObjectStore(objectname, desc); err == nil {

					objectStoreForeachIndex(data, func(nameindex, key string, unique, multientry bool) {
						opt := map[string]interface{}{}
						if unique {
							opt["unique"] = true
						}
						if multientry {
							opt["multientry"] = true
						}
						store.CreateIndex(nameindex, key, opt)
					})
				} else {

				}
			} else {

			}
		})
		openrequest.OnError(func(e event.Event) {
			errfunc.Invoke(e.JSObject())
		})
		return nil, nil
	})

}

func extractObjectFromReflection(elem reflect.Type, b object.Object) reflect.Value {
	m, _ := gomap.MapFromJSObject(b.JSObject()).(map[string]interface{})
	newobj := reflect.New(elem)
	inewobj := reflect.Indirect(newobj)

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		f := inewobj.FieldByName(field.Name)

		if f.Kind() == reflect.Int {
			v := m[strings.ToLower(field.Name)]
			iint := int(0)
			switch i := v.(type) {
			case int:
				iint = int(i)
			case int32:
				iint = int(i)
			case int64:
				iint = int(i)
			case uint:
				iint = int(i)
			case uint32:
				iint = int(i)
			case uint64:
				iint = int(i)
			}

			if f.CanSet() {
				f.Set(reflect.ValueOf(iint))
			}
		}

		if f.Kind() == reflect.Int32 {
			v := m[strings.ToLower(field.Name)]
			iint := int32(0)
			switch i := v.(type) {
			case int:
				iint = int32(i)
			case int32:
				iint = int32(i)
			case int64:
				iint = int32(i)
			case uint:
				iint = int32(i)
			case uint32:
				iint = int32(i)
			case uint64:
				iint = int32(i)
			}

			if f.CanSet() {
				f.Set(reflect.ValueOf(iint))
			}
		}
		if f.Kind() == reflect.Slice {
			v := m[strings.ToLower(field.Name)]
			switch f.Type().Elem().Kind() {
			case reflect.String:
				a := []string{}
				switch i := v.(type) {
				case []interface{}:
					for _, obj := range i {
						if s, ok := obj.(string); ok {
							a = append(a, s)
						}
					}
				}
				if f.CanSet() {
					f.Set(reflect.ValueOf(a))
				}
			}

		}
		if f.Kind() == reflect.String {
			v, _ := m[strings.ToLower(field.Name)].(string)
			if f.CanSet() {
				f.Set(reflect.ValueOf(v))
			}

		}

	}
	return inewobj
}

func Delete(data interface{}, rollback bool) (promise.Promise, error) {
	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		o := reflect.ValueOf(data)
		elem := o.Type()

		if store, ok := stores[elem.String()]; ok {
			if fieldname, ok := keyFactory[elem.String()]; ok {

				f := o.FieldByName(fieldname)
				req, err := store.Delete(f.Interface())
				if err != nil {
					if errjs, err := jserror.New(errors.New("unable to call delete")); err == nil {
						errfunc.Invoke(errjs.JSObject())
					} else {
						hogosuru.AssertErr(err)
					}
				}
				req.OnSuccess(func(e event.Event) {
					resolvefunc.Invoke()
				})

				req.OnError(func(e event.Event) {
					if !rollback {
						e.PreventDefault()
					}

					errfunc.Invoke(e.JSObject())
				})
			}
		}

		return nil, nil
	})
}

func opWrite(add bool, data interface{}, rollback bool) (promise.Promise, error) {

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		o := reflect.ValueOf(data)
		elem := o.Type()
		m := map[string]interface{}{}
		if store, ok := stores[elem.String()]; ok {
			for i := 0; i < elem.NumField(); i++ {
				field := elem.Field(i)
				m[strings.ToLower(field.Name)] = o.Field(i).Interface()

			}
			var (
				req indexeddb.IDBRequest
				err error
			)
			if add {
				req, err = store.Add(m)
			} else {
				req, err = store.Put(m)
			}

			if err == nil {
				req.OnSuccess(func(e event.Event) {
					resolvefunc.Invoke()
				})

				req.OnError(func(e event.Event) {
					if !rollback {
						e.PreventDefault()
					}

					errfunc.Invoke(e.JSObject())
				})

			} else {
				if errjs, err := jserror.New(err); err == nil {
					errfunc.Invoke(errjs.JSObject())
				} else {
					hogosuru.AssertErr(err)
				}
			}
		}

		return nil, nil
	})
}

func Add(dst interface{}, rollback bool) (promise.Promise, error) {
	return opWrite(true, dst, rollback)
}

func Put(dst interface{}, rollback bool) (promise.Promise, error) {
	return opWrite(false, dst, rollback)
}

func Get(dst interface{}, key interface{}) (promise.Promise, error) {
	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		o := reflect.ValueOf(dst)
		reflectdst := reflect.Indirect(o)
		elem := reflectdst.Type()
		if store, ok := stores[elem.String()]; ok {
			req, _ := store.Get(key)
			req.OnSuccess(func(e event.Event) {
				result, _ := req.Result()
				if b, ok := result.(object.Object); ok {

					inewobj := extractObjectFromReflection(elem, b)
					reflect.Indirect(o).Set(inewobj)

					resolvefunc.Invoke()

				} else {
					if errjs, err := jserror.New(errors.New("result response is not an object")); err == nil {
						errfunc.Invoke(errjs.JSObject())
					} else {
						hogosuru.AssertErr(err)
					}
				}

			})
			req.OnError(func(e event.Event) {
				errfunc.Invoke(e.JSObject())
			})

		}

		return nil, nil
	})
}

func GetAll(dst interface{}) (promise.Promise, error) {
	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		o := reflect.ValueOf(dst)
		reflectdst := reflect.Indirect(o)

		if reflectdst.Kind() == reflect.Slice {
			elem := reflectdst.Type().Elem()
			if store, ok := stores[elem.String()]; ok {
				req, _ := store.GetAll()
				req.OnSuccess(func(e event.Event) {
					result, _ := req.Result()
					if a, ok := result.(array.Array); ok {

						a.Every(func(i interface{}) bool {
							if b, ok := i.(object.Object); ok {
								inewobj := extractObjectFromReflection(elem, b)
								reflectdst = reflect.Append(reflectdst, inewobj)
							}

							return true
						})

						reflect.Indirect(o).Set(reflectdst)
						resolvefunc.Invoke()

					} else {
						if errjs, err := jserror.New(errors.New("result response is not an array")); err == nil {
							errfunc.Invoke(errjs.JSObject())
						} else {
							hogosuru.AssertErr(err)
						}
					}

				})
				req.OnError(func(e event.Event) {
					errfunc.Invoke(e.JSObject())
				})

			}

		}

		return nil, nil
	})

}
