package model

type InternalRequest struct {
	Content string `json:"content"`
}

type InternalResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
