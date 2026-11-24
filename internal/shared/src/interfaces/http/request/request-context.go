package request

import "time"

type Context struct {
	RequestID     string
	CorrelationID string
	IssuerID      *string
	IP            string
	StartedAt     time.Time
}

func (r Context) Duration() time.Duration {
	return time.Since(r.StartedAt)
}

const RequestContextKey = "request_context"
