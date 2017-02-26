package doc

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/utahta/swagger-doc/assets"
)

type SwaggerUIHandler struct{}

func NewSwaggerUIHandler() http.Handler {
	return &SwaggerUIHandler{}
}

func (h *SwaggerUIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := &assetfs.AssetFS{
		Asset:     assets.Asset,
		AssetDir:  assets.AssetDir,
		AssetInfo: assets.AssetInfo,
	}
	http.FileServer(fs).ServeHTTP(w, r)
}
