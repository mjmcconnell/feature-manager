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

	r.GET("/health_check", health_check)

	return r
}

func (r Router) StartServer() error {
	return r.Start("0.0.0.0:8080")
}

func health_check(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
