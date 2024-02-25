package internal

type Module interface {
	Name() string
	Init(r *Router)
}

var Modules []func() Module

func InitModules(r *Router) error {
	for _, f := range Modules {
		module := f()
		module.Init(r)
	}
	return nil
}
