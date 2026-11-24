package middleware

import (
	request "go-wallet/internal/shared/src/interfaces/http/request"
	"time"
	"uuid"

	"github.com/gin-gonic/gin"
)

func SetRequestContext(c *gin.Context) {
	requestID := uuid.NewV7().String()

	correlationID := c.GetHeader("X-Correlation-ID")

	if correlationID == "" {
		correlationID = requestID
	}

	ctx := &request.Context{
		RequestID:     requestID,
		CorrelationID: correlationID,
		IP:            c.ClientIP(),
		StartedAt:     time.Now(),
	}

	c.Set(request.RequestContextKey, ctx)

	c.Next()
}
