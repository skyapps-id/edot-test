package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/auth"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
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

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func TraceIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceID := trace.SpanContextFromContext(c.Request().Context()).TraceID().String()
		c.Response().Header().Set("X-Trace-Id", traceID)
		return next(c)
	}
}

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Get("error-handled") != nil {
			return
		}

		c.Set("error-handled", true)

		status := http.StatusBadRequest
		resp := response.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}

		if ae, ok := err.(*apperror.ApplicationError); ok {
			status = ae.Status
			resp.Message = ae.Message
		}

		_ = c.JSON(status, resp)
	}
}

func ValidateStaticToken(expectedToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Static-Token")
			if token == "" || token != expectedToken {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}
			return next(c)
		}
	}
}

func JWTMiddleware(secretKey []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing or invalid Authorization header",
				})
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := auth.ParseJWT(tokenStr, secretKey)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
			}

			c.Set("user_id", claims.UserID)

			return next(c)
		}
	}
}
