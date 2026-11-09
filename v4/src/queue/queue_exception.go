package queue

import "fmt"

// Error code constants for QueueException.
const (
	ErrCodeConnection = 1
	ErrCodeEnqueue    = 2
	ErrCodeDequeue    = 3
	ErrCodeUpdate     = 4

	E_CONNECTION = ErrCodeConnection
	E_ENQUEUE    = ErrCodeEnqueue
	E_DEQUEUE    = ErrCodeDequeue
	E_UPDATE     = ErrCodeUpdate
)

// QueueException represents an unrecoverable error during queue operations.
type QueueException struct {
	Code    int
	Message string
	Err     error
}

// QueueError is an alias for QueueException.
type QueueError = QueueException

func (e *QueueException) Error() string {
	return e.Message
}

func (e *QueueException) Unwrap() error {
	return e.Err
}

// ConnectionError creates a QueueException for connection errors.
func ConnectionError(msg string, prev ...error) *QueueException {
	var p error
	if len(prev) > 0 {
		p = prev[0]
	}
	return &QueueException{
		Code:    ErrCodeConnection,
		Message: fmt.Sprintf("Queue connection error: %s", msg),
		Err:     p,
	}
}

// EnqueueError creates a QueueException for enqueue failures.
func EnqueueError(msg string, prev ...error) *QueueException {
	var p error
	if len(prev) > 0 {
		p = prev[0]
	}
	return &QueueException{
		Code:    ErrCodeEnqueue,
		Message: fmt.Sprintf("Failed to enqueue: %s", msg),
		Err:     p,
	}
}

// DequeueError creates a QueueException for dequeue failures.
func DequeueError(msg string, prev ...error) *QueueException {
	var p error
	if len(prev) > 0 {
		p = prev[0]
	}
	return &QueueException{
		Code:    ErrCodeDequeue,
		Message: fmt.Sprintf("Failed to dequeue: %s", msg),
		Err:     p,
	}
}

// UpdateError creates a QueueException for job update failures.
func UpdateError(msg string, prev ...error) *QueueException {
	var p error
	if len(prev) > 0 {
		p = prev[0]
	}
	return &QueueException{
		Code:    ErrCodeUpdate,
		Message: fmt.Sprintf("Failed to update job: %s", msg),
		Err:     p,
	}
}
