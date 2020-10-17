package auth

import "testing"

func TestHashService_Generate(t *testing.T) {
	password := "password"
	c := &HashService{}
	hash, err := c.Generate(password)
	if err != nil {
		t.Fail()
	}

	println(hash)
}
