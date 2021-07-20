package io

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type errorType string

const (
	errorNone        errorType = ""
	errorTimeout     errorType = "timeout"
	errorCanceled    errorType = "canceled"
	errorExec        errorType = "execution"
	errorBadData     errorType = "bad_data"
	errorInternal    errorType = "internal"
	errorUnavailable errorType = "unavailable"
	errorNotFound    errorType = "not_found"
)

type Metric struct {
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
}

type Value struct {
	Value    float64
	ValueStr string
}

type Result struct {
	Metric Metric   `json:"metric"`
	Value  []string `json:"value"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Results    []Result `json:"result"`
}

type PromResponse struct {
	Status    status    `json:"status"`
	Data      Data      `json:"data,omitempty"`
	ErrorType errorType `json:"errorType,omitempty"`
	Error     string    `json:"error,omitempty"`
	Warnings  []string  `json:"warnings,omitempty"`
}
