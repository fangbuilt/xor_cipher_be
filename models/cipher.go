package models

type CipherRequest struct {
	Text string `json:"text" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

type CipherResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result,omitempty"`
	Error   string `json:"error,omitempty"`
}
