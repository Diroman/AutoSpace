package database

import "errors"

var (
	UserNotFound = errors.New("User not found!\n")
	IncorrectPassword = errors.New("Incorrect password!\n")
)


