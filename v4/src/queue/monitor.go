package queue

type QueueStats struct {
	Pending int
	Success int
	DLQ     int
}
