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
		url       string
		targetURL string
		handler   func(c echo.Context) error
		resp      string
		status    int
	}{
		"Test get request on hello handler": {
			url:       "/hello",
			targetURL: "/hello",
			handler:   helloHandler,
			resp:      "Hello from the hello module",
			status:    200,
		},
		"Test get request on hello handler with bad path": {
			url:       "/hello",
			targetURL: "/hi",
			handler:   helloHandler,
			resp: `{"message":"Not Found"}
`,
			status: 404,
		},
		"Test get request on goodbye handler": {
			url:       "/goodbye",
			targetURL: "/goodbye",
			handler:   goodbyeHandler,
			resp:      "Goodbye from the hello module",
			status:    200,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// arrange
			e := echo.New()
			e.GET(tc.url, tc.handler)

			// act
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tc.targetURL, nil)
			e.ServeHTTP(w, r)

			// assert
			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.resp, w.Body.String())
		})
	}
}
