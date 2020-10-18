package spaceCounter

import (
	"parking/model"
	"testing"
)

func Test_spaceCounter_GetDistanceInMeters(t *testing.T) {
	a := model.Point{
		X: 560,
		Y: 832,
	}
	b := model.Point{
		X: 460,
		Y: 328,
	}

	sc := spaceCounter{}
	got := sc.GetDistanceInMeters(a, b)

	println(got)
}