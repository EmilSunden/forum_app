package logger

import (
	"app/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// maxPayloadSize defines the maximum number of bytes of the request body to log
const maxPayloadSize = 1024 // 1KB

// maskSensitiveFields accepts a JSON payload as bytes, decodes it,
// and masks sensitive fields. If decoding fails, it returns the original data.
func maskSensitiveFields(data []byte) []byte {
	// Attempt to unmarshal the JSON.
	var payload map[string]interface{}
	if err := json.Unmarshal(data, &payload); err != nil {
		// Not JSON or unable to parse, so just return the original data.
		return data
	}

	// Mask sensitive fields.
	sensitiveKeys := []string{"password", "token", "secret"}

	for _, key := range sensitiveKeys {
		if val, ok := payload[key]; ok {
			if str, ok := val.(string); ok {
				// Mask all but the last 4 characters.
				if len(str) > 4 {
					payload[key] = "****" + str[len(str)-4:]
				} else {
					payload[key] = "****"
				}
			}
		}
	}

	// Marshal the masked payload back to JSON.
	masked, err := json.Marshal(payload)
	if err != nil {
		// Unable to marshal, so just return the original data.
		return data
	}
	return masked
}

// Logger
func Logger(logger zerolog.Logger, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// logger logic goes here
		start := time.Now()

		var bodyStr string

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			limitedReader := io.LimitReader(c.Request.Body, maxPayloadSize)
			data, err := io.ReadAll(limitedReader)
			if err == nil && len(data) > 0 {
				// If the actual body is larger than our limit,
				// add an indicator that it was truncated.
				if c.Request.ContentLength > maxPayloadSize {
					data = append(data, []byte("...")...)
				}
				// Mask sensitive fields if the body is JSON.
				maskedData := maskSensitiveFields(data)
				bodyStr = string(maskedData)
			}
			// Replace the request body so downstream handlers can reaed it
			c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		}

		c.Next()

		elapsed := time.Since(start)
		statusCode := c.Writer.Status()

		// Build a log entry.
		entry := models.RequestLog{
			Model: gorm.Model{
				CreatedAt: time.Now(),
			},
			Method:     c.Request.Method,
			URL:        c.Request.RequestURI,
			StatusCode: statusCode,
			ElapsedMS:  elapsed.Milliseconds(),
			Body:       bodyStr,
		}

		// Log asynchronously to the DB so it doesn't block the request.
		go func(e models.RequestLog) {
			if err := db.Create(&e).Error; err != nil {
				logger.Error().Err(err).Msg("Failed to log request")
			}
		}(entry)

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", statusCode).
			Dur("elapsed", elapsed)

	}
}
