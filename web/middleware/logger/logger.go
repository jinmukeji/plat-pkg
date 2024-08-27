package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger logrus.FieldLogger) gin.HandlerFunc {

	return func(c *gin.Context) {

		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := stop
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		size := c.Writer.Size()
		if size < 0 {
			size = 0
		}

		entry := logger.WithFields(logrus.Fields{
			"status":     statusCode,
			"latency":    latency,
			"client_ip":  clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"size":       size,
			"user_agent": clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			switch {
			case statusCode >= 200 && statusCode <= 299:
				entry.Info()
			case statusCode >= 300 && statusCode <= 399:
				entry.Info()
			case statusCode >= 400 && statusCode <= 499:
				entry.Warn()
			default:
				entry.Error()
			}
		}
	}
}
