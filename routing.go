package hogosuru

import (
	"errors"
	"sync"

	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/window"
)

var (
	//ErrRoutingAlreadyPresent ErrRoutingAlreadyPresent err
	ErrRoutingAlreadyPresent = errors.New("This route is already present")
)

const (
	//STDROUTE
	STDROUTE = iota
	//HASHROUTE
	HASHROUTE
)

var singletonRoute sync.Once
var route RouteMap

//Rendering interfacee
type Rendering interface {
	OnLoad(d document.Document, n node.Node, route string) []Rendering
	Node() node.Node
	OnUnload()
}

//Router struct
type RouteMap struct {
	mode             int
	defaultRendering Rendering
	currentHashRoute string
	currentStdRoute  string
	currentRendering Rendering
	routing          map[string]Rendering
}

func init() {
	if w, err := window.New(); err == nil {

		//Vérifier si cette fonction a du sens
		w.OnHashChange(func(e event.Event) {
			r := Router()
			if r.mode == HASHROUTE {
				r.onurlchange()
			}

		})

		w.OnPopState(func(e event.Event) {
			r := Router()

			r.onurlchange()

		})

	} else {
		println("Router " + err.Error())
	}
	//get the current location

	//on new event hash , load the new observer (https://developer.mozilla.org/en-US/docs/Web/API/Window/hashchange_event)

}

//Router
func Router() *RouteMap {

	singletonRoute.Do(func() {
		route.routing = make(map[string]Rendering)
	})

	return &route
}

func (r *RouteMap) DefaultRendering(obj Rendering) {
	r.defaultRendering = obj
	if d, err := document.New(); err == nil {
		if body, err := d.Body(); err == nil {
			r.loadChilds(d, obj, body)
		}
	}
}

func (r *RouteMap) Route() string {
	if r.mode == STDROUTE {
		return r.currentStdRoute
	}
	return r.currentHashRoute
}

func (r *RouteMap) SetRoute(route string) {
	if r.mode == STDROUTE {
		r.currentStdRoute = route
	}
	r.currentHashRoute = route
}

func (r *RouteMap) loadChilds(d document.Document, obj Rendering, node node.Node) {
	arrayRendering := obj.OnLoad(d, node, r.Route())
	if arrayRendering != nil {
		for _, render := range arrayRendering {
			r.loadChilds(d, render, obj.Node())
		}
	}

	node.AppendChild(obj.Node())
}

func (r *RouteMap) Go(newroute string) {

	if w, err := window.New(); err == nil {

		if historyObj, err := w.History(); err == nil {
			historyObj.PushState(nil, newroute, newroute)
			r.onurlchange()
		}
	}

}

func (r *RouteMap) onChangeRoute(newroute string) {
	if len(r.routing) == 0 {
		r.SetRoute(newroute)

	}

	for route, render := range r.routing {
		if newroute == route {
			if r.currentRendering != nil {
				r.currentRendering.OnUnload()
			}
			r.SetRoute(newroute)
			r.LoadRendering(render)
		}
	}

}
func (r *RouteMap) LoadRendering(obj Rendering) {

	r.currentRendering = obj
	if d, err := document.New(); err == nil {

		if r.defaultRendering != nil {
			r.loadChilds(d, obj, r.defaultRendering.Node())
		} else {
			if body, err := d.Body(); err == nil {
				r.loadChilds(d, obj, body)

			}
		}
	}
}

func (r *RouteMap) Start(mode int) {
	r.mode = mode
	r.onurlchange()

}

func (r *RouteMap) Add(route string, obj Rendering) error {
	var err error
	if _, ok := r.routing[route]; ok {
		err = ErrRoutingAlreadyPresent
	} else {

		r.routing[route] = obj
	}
	return err
}

func (r *RouteMap) onhashechange() {
	if w, err := window.New(); err == nil {

		if l, err := w.Location(); err == nil {

			if h, err := l.Hash(); err == nil {

				if len(h) > 1 {
					r.onChangeRoute(h[1:])
				} else {
					r.onChangeRoute("")
				}

			} else {
				println("Router " + err.Error())
			}

		} else {
			println("Router " + err.Error())
		}

	} else {
		println("Router " + err.Error())
	}

}

func (r *RouteMap) onurlchange() {
	if w, err := window.New(); err == nil {

		if l, err := w.Location(); err == nil {
			var route string = ""
			var err error
			if r.mode == STDROUTE {
				route, err = l.Pathname()

			} else {
				route, err = l.Hash()
				if len(route) > 1 {
					route = route[1:]

				}
			}
			if err == nil {

				r.onChangeRoute(route)

			} else {
				println("Router " + err.Error())
			}

		} else {
			println("Router " + err.Error())
		}

	} else {
		println("Router " + err.Error())
	}

}
