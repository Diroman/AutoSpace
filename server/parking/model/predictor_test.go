package model

import (
	"fmt"
	"testing"
)

func TestPoint_GetAngle(t *testing.T) {
	p := Point{
		X: 1,
		Y: 1,
	}
	pp := Point{
		X: -1,
		Y: -1,
	}

	got := p.GetAngle(pp)

	fmt.Println(got)
}
