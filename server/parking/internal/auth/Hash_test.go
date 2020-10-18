package auth

import "testing"

func TestHashService_Generate(t *testing.T) {
	password := "1"
	c := &HashService{}
	hash, err := c.Generate(password)
	if err != nil {
		t.Fail()
	}

	println(hash)
}
