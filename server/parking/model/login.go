package model

type Coordinate struct {
	Latitude  float32
	Longitude float32
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User UserResponse
	Err  string
}
