package company

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/internal/models/api"
	"companies/pkg/constants"
	http_errors "companies/pkg/errors"
)

const (
	moduleUpdateCompanyHandler = "update_company_handler"
)

func (h *Handler) Update(ctx *fasthttp.RequestCtx) {
	req, err := h.extractUpdateRequest(ctx)
	if err != nil {
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	if err = h.validator.ValidateUpdateRequest(req); err != nil {
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	resp, err := h.useCase.Update(ctx, req)
	if err != nil {
		h.logger.Error(err.Error(), zap.String(constants.FieldModule, moduleUpdateCompanyHandler))
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.ReplySuccess(ctx, resp)
}

func (h *Handler) extractUpdateRequest(ctx *fasthttp.RequestCtx) (req *api.UpdateCompanyRequest, err error) {
	req = new(api.UpdateCompanyRequest)
	if err = json.Unmarshal(ctx.PostBody(), req); err != nil {
		return req, http_errors.ErrInvalidJSON
	}
	return req, nil
}
