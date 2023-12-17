package api

import (
	"net/http"

	"github.com/minguu42/sandbox/teratera/gen/teraterapb/v1/teraterapbconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type handler struct{}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(teraterapbconnect.NewTerateraServiceHandler(handler{}))
	return h2c.NewHandler(mux, &http2.Server{})
}
