package company

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/internal/models/api"
	"companies/pkg/constants"
)

const (
	moduleGetCompanyHandler = "get_company_handler"
)

func (h *Handler) Get(ctx *fasthttp.RequestCtx) {
	id, err := h.extractId(ctx)
	if err != nil {
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	resp, err := h.useCase.Get(ctx, &api.GetCompanyRequest{ID: id})
	if err != nil {
		h.logger.Error(err.Error(), zap.String(constants.FieldModule, moduleGetCompanyHandler))
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.ReplySuccess(ctx, resp)
}
