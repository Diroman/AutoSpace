package auth

import "golang.org/x/crypto/bcrypt"

type HashService struct{}

var Hash = HashService{}

//Generate a salted hash for the input string
func (c *HashService) Generate(s string) (string, error) {
	password := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *HashService) ComparePassword(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}