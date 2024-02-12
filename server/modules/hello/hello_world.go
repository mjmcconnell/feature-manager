package hello

import (
	"github.com/labstack/echo/v4"

	"github.com/mjmcconnell/feature-manager/internal"
)

type HelloModule struct{}

func init() {
	internal.Modules = append(internal.Modules, NewHelloModule)
}

func NewHelloModule() internal.Module {
	return &HelloModule{}
}

func (m HelloModule) Name() string {
	return "Hello Module"
}

func (m HelloModule) InitRoutes(e *echo.Echo) {
	e.GET("/hello", helloHandler)
}
