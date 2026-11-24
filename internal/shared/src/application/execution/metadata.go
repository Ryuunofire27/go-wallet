package execution

type ExecutionMetadata struct {
	RequestID     string
	CorrelationID string
	IssuerID      *string
}
