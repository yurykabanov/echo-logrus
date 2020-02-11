package echologrus

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type config struct {
	Skipper middleware.Skipper
}

type Option func(*config)

func WithSkipper(skipper middleware.Skipper) Option {
	return func(c *config) {
		c.Skipper = skipper
	}
}

func Middleware(logger *logrus.Logger, opts ...Option) echo.MiddlewareFunc {
	config := &config{
		Skipper: middleware.DefaultSkipper,
	}

	for _, opt := range opts {
		opt(config)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if config.Skipper(ctx) {
				return next(ctx)
			}

			var err error

			request := ctx.Request()
			response := ctx.Response()

			// Measure request processing time
			start := time.Now()
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}
			stop := time.Now()

			latency := stop.Sub(start)

			// Try to extract request ID from request (propagated) or response (generated)
			requestId := request.Header.Get(echo.HeaderXRequestID)
			if requestId == "" {
				requestId = response.Header().Get(echo.HeaderXRequestID)
			}

			path := request.URL.Path
			if path == "" {
				path = "/"
			}

			// Ignore error if any
			bytesIn, _ := strconv.ParseInt(request.Header.Get(echo.HeaderContentLength), 10, 64)

			logger.WithFields(logrus.Fields{
				"request_id":    requestId,
				"remote_ip":     ctx.RealIP(),
				"host":          request.Host,
				"uri":           request.RequestURI,
				"method":        request.Method,
				"path":          path,
				"referer":       request.Referer(),
				"user_agent":    request.UserAgent(),
				"status":        response.Status,
				"latency_ns":    latency.Nanoseconds(),
				"latency_human": latency.String(),
				"bytes_in":      bytesIn,
				"bytes_out":     response.Size,
			}).Info("HTTP Request")

			return nil
		}
	}
}
