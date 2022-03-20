package company

import (
	"github.com/valyala/fasthttp"

	"companies/pkg/constants"
	http_errors "companies/pkg/errors"
)

type IHandler interface {
	Create(ctx *fasthttp.RequestCtx)
	Update(ctx *fasthttp.RequestCtx)
	Get(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
}

func (h *Handler) extractId(ctx *fasthttp.RequestCtx) (id string, err error) {
	var ok bool

	id, ok = ctx.UserValue(constants.ParamCustomerID).(string)
	if !ok {
		return id, http_errors.ErrEmptyParamID
	}

	return id, nil
}
