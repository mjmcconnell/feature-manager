package internal

import (
	"github.com/labstack/echo/v4"
)

type Module interface {
	Name() string
	InitRoutes(e *echo.Echo)
}

var Modules []func() Module
