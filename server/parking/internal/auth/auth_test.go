package auth

import "testing"

func TestCreateNewToken(t *testing.T) {
	got, err := CreateNewToken(1)
	if err != nil {
		t.Fail()
	}

	println(got)
}
