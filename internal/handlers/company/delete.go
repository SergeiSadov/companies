package company

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/internal/models/api"
	"companies/pkg/constants"
)

const (
	moduleDeleteCompanyHandler = "delete_company_handler"
)

func (h *Handler) Delete(ctx *fasthttp.RequestCtx) {
	id, err := h.extractId(ctx)
	if err != nil {
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	if err := h.useCase.Delete(ctx, &api.DeleteCompanyRequest{ID: id}); err != nil {
		h.logger.Error(err.Error(), zap.String(constants.FieldModule, moduleDeleteCompanyHandler))
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.ReplySuccess(ctx, nil)
}
