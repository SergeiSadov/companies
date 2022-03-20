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
	moduleCreateCompanyHandler = "create_company_handler"
)

func (h *Handler) Create(ctx *fasthttp.RequestCtx) {
	req, err := h.extractCreateRequest(ctx)
	if err != nil {
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	resp, err := h.useCase.Create(ctx, req)
	if err != nil {
		h.logger.Error(err.Error(), zap.String(constants.FieldModule, moduleCreateCompanyHandler))
		h.ReplyCustomError(ctx, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.ReplySuccess(ctx, resp)
}

func (h *Handler) extractCreateRequest(ctx *fasthttp.RequestCtx) (req *api.CreateCompanyRequest, err error) {
	req = new(api.CreateCompanyRequest)
	if err = json.Unmarshal(ctx.PostBody(), req); err != nil {
		return req, http_errors.ErrInvalidJSON
	}
	return req, nil
}
