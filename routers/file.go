package routers

import (
	"github.com/docker-golang-mysql/controller"
	"github.com/gorilla/mux"
)

func FileRouters(router *mux.Router) {
	fileRoutes := router.PathPrefix("/file").Subrouter()

	routes = append(routes, route{Router: fileRoutes, Path: "/upload", Func: controller.Upload, Method: "POST"})
	routes = append(routes, route{Router: fileRoutes, Path: "/readfile/{image}", Func: controller.ReadFile, Method: "GET"})
	routes = append(routes, route{Router: fileRoutes, Path: "/writefile", Func: controller.WriteCSV, Method: "GET"})

	for _, r := range routes {
		r.Router.HandleFunc(r.Path, r.Func).Methods(r.Method)
	}
}
