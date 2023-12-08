package api

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}
