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
