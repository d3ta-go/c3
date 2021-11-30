package chi

import (
	"github.com/d3ta-go/c3/net/http"
	chi "github.com/go-chi/chi/v5"
)

type Chi struct {
	http.BaseHTTPAdapter
	ctx *chi.Context
	app *chi.Mux
}
