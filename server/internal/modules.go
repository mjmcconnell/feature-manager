package internal

type Module interface {
	Name() string
	InitRoutes(r *Router)
}

var Modules []func() Module
