package echo

import (
	"github.com/d3ta-go/c3/net/http"
	echo "github.com/labstack/echo/v4"
)

type Echo struct {
	http.BaseHTTPAdapter
	ctx *echo.Context
	app *echo.Echo
}
