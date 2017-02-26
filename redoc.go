package doc

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func NewRedocHandler(specURL, path string) http.Handler {
	return middleware.Redoc(middleware.RedocOpts{
		BasePath: "/",
		SpecURL:  specURL,
		Path:     path,
	}, http.NotFoundHandler())
}
