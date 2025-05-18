package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/skyapps-id/edot-test/product-service/container"
	"github.com/skyapps-id/edot-test/product-service/pkg/logger"
	"github.com/skyapps-id/edot-test/product-service/pkg/tracer"
	pkgValidator "github.com/skyapps-id/edot-test/product-service/pkg/validator"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func StartHTTP(container *container.Container) {
	if container == nil {
		panic("container is nil")
	}

	tracer := tracer.InitTracer(container.Config.HostOTLP, container.Config.AppName)
	defer tracer(context.Background())
	otel.SetTextMapPropagator(propagation.TraceContext{})

	if err := logger.Init(); err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer logger.Log.Sync()

	if err := logger.Init(); err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer logger.Log.Sync()

	server := echo.New()

	server.Use(otelecho.Middleware(container.Config.AppName))
	server.Use(TraceIDMiddleware)
	server.Use(mw.Recover())
	server.Use(MiddlewareLoggerWithTrace(logger.Log))

	server.HTTPErrorHandler = ErrorHandler()
	server.Validator = &DataValidator{ValidatorData: pkgValidator.SetupValidator()}

	Router(server, container)

	go func() {
		if err := server.Start(fmt.Sprint(":", container.Config.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\nShutting down server http...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server http exited properly")
}
