package database

import (
	"fmt"
	"testing"
)

func TestDatabase_GetParkingSpace(t *testing.T) {
	id := 1
	d := NewDatabase("localhost", 5432, "postgres", "postgres", "password")
	got, err := d.GetParkingSpace(id)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(got)
}