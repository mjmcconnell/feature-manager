package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Module")
}
