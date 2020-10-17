package model

type LoginRequest struct {
	Login    string
	Password string
}

type LoginResponse struct {
	// add user info
	Token string
	Err   string
}
