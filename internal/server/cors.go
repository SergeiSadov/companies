package server

import (
	"reflect"

	"github.com/valyala/fasthttp"
)

func Cors(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, "*")
		ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json; charset=utf-8")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowMethods, "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowHeaders, "Origin, Content-Type, Authorization, Request-ID")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlExposeHeaders, "Authorization")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlMaxAge, "3600")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowCredentials, "true")

		if reflect.DeepEqual(ctx.Request.Header.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.Response.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}
