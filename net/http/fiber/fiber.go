package fiber

import (
	"github.com/d3ta-go/c3/net/http"
	fiber "github.com/gofiber/fiber/v2"
)

type Fiber struct {
	http.BaseHTTPAdapter
	ctx *fiber.Ctx
	app *fiber.App
}
