package queue

type ErrorClassification struct {
	Category  string
	Retryable bool
	Status    string
	Detail    string
}

func Classify(httpCode int) ErrorClassification {
	if httpCode >= 200 && httpCode < 300 {
		return ErrorClassification{Category: "success", Retryable: false, Status: "success", Detail: "OK"}
	}
	if httpCode == 401 {
		return ErrorClassification{Category: "unauthorized", Retryable: true, Status: "pending", Detail: "Unauthorized (401)"}
	}
	if httpCode == 429 || (httpCode >= 500 && httpCode < 600) {
		return ErrorClassification{Category: "server_or_rate_limit", Retryable: true, Status: "pending", Detail: "Retryable error"}
	}
	return ErrorClassification{Category: "client_error", Retryable: false, Status: "dlq", Detail: "Client error / DLQ"}
}
