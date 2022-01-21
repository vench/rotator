package http

import (
	"context"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (s *Server) router(ctx context.Context) fasthttp.RequestHandler {
	mux := router.New()

	mux.GET("/check", check)
	mux.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))

	return mux.Handler
}

func check(rCtx *fasthttp.RequestCtx) {
	rCtx.SetStatusCode(http.StatusOK)
	rCtx.SetBodyString(".")
}
