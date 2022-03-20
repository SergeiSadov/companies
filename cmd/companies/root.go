package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/internal/configs"
	"companies/internal/definitions"
	"companies/pkg/constants"
)

const (
	moduleMain = "main"
)

var rootCmd = &cobra.Command{
	Use:   "companies",
	Short: "companies crud service",
	Run:   RunService,
}

func RunService(_ *cobra.Command, _ []string) {
	di, err := definitions.Build()
	if err != nil {
		log.Fatalf("failed to create di, error: %v", err)
	}

	server := di.Get(definitions.HttpDef).(*fasthttp.Server)
	logger := di.Get(definitions.LoggerDef).(*zap.Logger)
	cfg := di.Get(definitions.CfgDef).(configs.Config)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err = server.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port))
		if err != nil {
			if err != http.ErrServerClosed {
				logger.Error(err.Error(), zap.String(constants.FieldModule, moduleMain))
			}
			return
		}
	}()
	logger.Info(fmt.Sprintf("started listening port %d", cfg.App.Port), zap.String(constants.FieldModule, moduleMain))

	sig := <-sigChan
	logger.Info(fmt.Sprintf("got signal %v, starting shutdown", sig), zap.String(constants.FieldModule, moduleMain))

	if err = server.Shutdown(); err != nil {
		if err != http.ErrServerClosed {
			logger.Error(err.Error(), zap.String(constants.FieldModule, moduleMain), zap.String(constants.FieldAction, "server_shutdown"))
		}
	}

	os.Exit(0)
}
