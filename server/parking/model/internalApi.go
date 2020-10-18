package model

type InternalRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type InternalResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
