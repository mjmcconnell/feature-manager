package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mjmcconnell/feature-manager/internal/middleware"
)

type Router struct {
	*echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	e.Use(middleware.NewLogMiddleware())
	r := &Router{e}

	initModuleRoutes(r)

	r.GET("/health_check", health_check)

	return r
}

func (r Router) StartServer() error {
	return r.Start("localhost:8080")
}

func health_check(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func initModuleRoutes(r *Router) error {
	for _, f := range Modules {
		module := f()
		module.InitRoutes(r)
	}
	return nil
}
