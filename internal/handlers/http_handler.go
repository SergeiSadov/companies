package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/pkg/constants"
)

type IHandler interface {
	ReplySuccess(ctx *fasthttp.RequestCtx, result interface{})
	ReplyCustomError(ctx *fasthttp.RequestCtx, code int)
	ReplyWithBody(ctx *fasthttp.RequestCtx, body []byte, statusCode int)
	MetricsWrap(handle ActionFunction) func(ctx *fasthttp.RequestCtx)
	PanicHandler(ctx *fasthttp.RequestCtx, _ interface{})
}

type ActionFunction func(ctx *fasthttp.RequestCtx)

type Handler struct {
	logger *zap.Logger
}

func NewHttpHandler(logger *zap.Logger) *Handler {
	metricsRegister.Do(func() {
		prometheus.MustRegister(totalRequests, responseStatus)
	})

	return &Handler{
		logger: logger,
	}
}

func (h *Handler) MetricsWrap(handle ActionFunction) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		if len(ctx.Request.Body()) == 0 {
			ctx.Request.SetBody([]byte(ctx.URI().String()))
		}
		ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json")

		path := string(ctx.URI().Path())
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		handle(ctx)
		responseStatus.WithLabelValues(strconv.Itoa(ctx.Response.Header.StatusCode())).Inc()
		totalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	}
}

func (h *Handler) ReplyCustomError(ctx *fasthttp.RequestCtx, statusCode int) {
	h.replyError(ctx, statusCode)
}

func (h *Handler) ReplySuccess(ctx *fasthttp.RequestCtx, result interface{}) {
	if result == nil {
		h.ReplyWithBody(ctx, nil, fasthttp.StatusOK)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		h.logger.Error(err.Error(), zap.String(constants.FieldModule, "handler"))
		return
	}

	h.ReplyWithBody(ctx, data, fasthttp.StatusOK)
}

func (h *Handler) ReplyWithBody(ctx *fasthttp.RequestCtx, body []byte, statusCode int) {
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(body)
}

func (h *Handler) PanicHandler(ctx *fasthttp.RequestCtx, req interface{}) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusInternalServerError))
}

func (h *Handler) replyError(ctx *fasthttp.RequestCtx, statusCode int) {
	ctx.SetStatusCode(statusCode)
	ctx.SetBodyString(fasthttp.StatusMessage(statusCode))
}
