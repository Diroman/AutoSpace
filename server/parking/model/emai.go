package model

type EmailRequest struct {
	ErrorCode int    `json:"error_code"`
	Comment   string `json:"comment"`
	Email     string `json:"email"`
}
