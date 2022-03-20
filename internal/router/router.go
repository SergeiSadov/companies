package router

import (
	"fmt"
	"time"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"

	"companies/internal/configs"
	"companies/internal/handlers"
	"companies/internal/handlers/company"
)

type Config struct {
	InternalHandler handlers.IHandler
	CompanyHandler  company.IHandler
	App             configs.App
	Logger          *zap.Logger
}

func New(cfg Config) *router.Router {
	r := router.New()

	setMethods(r, cfg)

	return r
}

func setMethods(r *router.Router, cfg Config) {
	r.POST("/company", ResponseBodyLoggerMiddleware(cfg.InternalHandler.MetricsWrap(cfg.CompanyHandler.Create), cfg.Logger))
	r.PUT("/company", ResponseBodyLoggerMiddleware(cfg.InternalHandler.MetricsWrap(cfg.CompanyHandler.Update), cfg.Logger))
	r.DELETE("/company/{companyId}", ResponseBodyLoggerMiddleware(cfg.InternalHandler.MetricsWrap(cfg.CompanyHandler.Delete), cfg.Logger))
	r.GET("/company/{companyId}", ResponseBodyLoggerMiddleware(cfg.InternalHandler.MetricsWrap(cfg.CompanyHandler.Get), cfg.Logger))

	r.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))

	r.PanicHandler = cfg.InternalHandler.PanicHandler
}

func ResponseBodyLoggerMiddleware(handle handlers.ActionFunction, lgr *zap.Logger) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		t := time.Now()
		handle(ctx)
		diff := time.Since(t).Milliseconds()
		format := "url: %v, code: %v, time (ms): %v"
		params := []interface{}{
			ctx.Request.URI().String(), ctx.Response.StatusCode(), diff,
		}
		lgr.Debug(fmt.Sprintf(format, params...))
	}
}
