package company

import (
	"go.uber.org/zap"

	"companies/internal/handlers"
	"companies/internal/handlers/company/validators"
	"companies/internal/usecases/company"
	"companies/pkg/error_adapters/http_adapter"
)

type Handler struct {
	handlers.IHandler
	useCase      company.IUseCase
	logger       *zap.Logger
	errorAdapter http_adapter.IErrorAdapter
	validator    validators.IValidator
}

type Config struct {
	Internal     handlers.IHandler
	UseCase      company.IUseCase
	Logger       *zap.Logger
	ErrorAdapter http_adapter.IErrorAdapter
	Validator    validators.IValidator
}

func NewHandler(
	cfg Config,
) *Handler {
	return &Handler{
		IHandler:     cfg.Internal,
		useCase:      cfg.UseCase,
		logger:       cfg.Logger,
		errorAdapter: cfg.ErrorAdapter,
		validator:    cfg.Validator,
	}
}
