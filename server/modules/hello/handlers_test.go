package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_handlerGetRequests(t *testing.T) {
	testCases := map[string]struct {
		url     string
		handler func(c echo.Context) error
		resp    string
	}{
		"Test get request on hello handler": {
			url:     "/hello",
			handler: helloHandler,
			resp:    "Hello from the hello module",
		},
		"Test get request on goodbye handler": {
			url:     "/goodbye",
			handler: goodbyeHandler,
			resp:    "Goodbye from the hello module",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// arrange
			e := echo.New()
			e.GET(tc.url, tc.handler)

			// act
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			// assert
			assert.Equal(t, 200, rec.Code)
			assert.Equal(t, tc.resp, rec.Body.String())
		})
	}
}
