package model

type User struct {
	Id           int
	Token        string
	Username     string
	Mobile       string
	Address      string
	Email        string
	Password     string
	AddressCoord Coordinate
}

type UserResponse struct {
	Id           int
	Token        string
	Username     string
	Mobile       string
	Address      string
	Email        string
	AddressCoord Coordinate
}

func UserToUserResponse(user User) UserResponse {
	return UserResponse{
		Id:           user.Id,
		Token:        user.Token,
		Username:     user.Username,
		Mobile:       user.Mobile,
		Address:      user.Address,
		Email:        user.Email,
		AddressCoord: user.AddressCoord,
	}
}
