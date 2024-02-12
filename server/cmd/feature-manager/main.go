package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mjmcconnell/feature-manager/internal"
	// modules
	_ "github.com/mjmcconnell/feature-manager/modules/hello"
)

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func initModules(e *echo.Echo) error {
	for _, f := range internal.Modules {
		module := f()
		module.InitRoutes(e)
	}
	return nil
}

func main() {

	e := echo.New()
	e.GET("/", root)

	initModules(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
