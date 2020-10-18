package model

type FrameRequest struct {
	Id int `json:"id"`
}

type FrameResponse struct {
	Image string `json:"image"`
}
