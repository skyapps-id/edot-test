package http

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func MiddlewareLoggerWithTrace(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			span := trace.SpanFromContext(c.Request().Context())
			traceID := span.SpanContext().TraceID().String()

			req := c.Request()
			res := c.Response()

			logger.Info("request",
				zap.String("trace_id", traceID),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.Int("status", res.Status),
				zap.String("remote", c.RealIP()),
				zap.Duration("latency", stop.Sub(start)),
			)
			return err
		}
	}
}
