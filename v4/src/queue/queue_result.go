package queue

// QueueResult represents the immutable result of a queue job processed by Worker.
type QueueResult struct {
	ID       int
	Status   string
	Error    string
	Response map[string]any
	Attempts int
}

// NewSuccessResult creates a QueueResult representing a successful job execution.
func NewSuccessResult(id int, response map[string]any, attempts int) *QueueResult {
	if attempts <= 0 {
		attempts = 1
	}
	return &QueueResult{
		ID:       id,
		Status:   StatusSuccess,
		Error:    "",
		Response: response,
		Attempts: attempts,
	}
}

// NewFailedResult creates a QueueResult representing a failed (retriable) job execution.
func NewFailedResult(id int, errMsg string, response map[string]any, attempts int) *QueueResult {
	if attempts <= 0 {
		attempts = 1
	}
	return &QueueResult{
		ID:       id,
		Status:   StatusFailed,
		Error:    errMsg,
		Response: response,
		Attempts: attempts,
	}
}

// NewDlqResult creates a QueueResult representing a job moved to the Dead-Letter Queue.
func NewDlqResult(id int, reason string, response map[string]any, attempts int) *QueueResult {
	return &QueueResult{
		ID:       id,
		Status:   StatusDlq,
		Error:    reason,
		Response: response,
		Attempts: attempts,
	}
}

func (r *QueueResult) IsSuccess() bool {
	return r.Status == StatusSuccess
}

func (r *QueueResult) IsFailed() bool {
	return r.Status == StatusFailed
}

func (r *QueueResult) IsDlq() bool {
	return r.Status == StatusDlq
}

func (r *QueueResult) IsRetriable() bool {
	return r.IsFailed() && r.Attempts < DefaultMaxAttempts
}
