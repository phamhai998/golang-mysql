package routers

import (
	"net/http"

	"github.com/docker-golang-mysql/controller"
	"github.com/gorilla/mux"
)

type route struct {
	Router *mux.Router
	Path   string
	Func   func(http.ResponseWriter, *http.Request)
	Method string
}

var routes []route

func PersonRouters(router *mux.Router) []route {
	personRoutes := router.PathPrefix("/person").Subrouter()
	routes = append(routes, route{Router: personRoutes, Path: "/get", Func: controller.GetPerson, Method: "GET"})
	routes = append(routes, route{Router: personRoutes, Path: "/insert", Func: controller.CreatePerson, Method: "POST"})
	routes = append(routes, route{Router: personRoutes, Path: "/search", Func: controller.SearchPerson, Method: "POST"})
	routes = append(routes, route{Router: personRoutes, Path: "/update", Func: controller.UpdatePerson, Method: "POST"})
	routes = append(routes, route{Router: personRoutes, Path: "/delete", Func: controller.DeletePerson, Method: "POST"})

	for _, r := range routes {
		r.Router.HandleFunc(r.Path, r.Func).Methods(r.Method)
	}

	return routes
}
