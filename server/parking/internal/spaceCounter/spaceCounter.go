package spaceCounter

import (
	"parking/model"
)

type spaceCounter struct {}

var SpaceCounter = spaceCounter{}

func (sc spaceCounter) GetSpaceCount(cars model.Prediction) int {
	nearestCars := map[int]model.Row{}
	for _, list := range cars.Classes {
		for _, car := range list.Rows {
			sc.CalculateDistance(car, list.Rows, nearestCars)
		}
	}

	println(cars)
	return 0
}

func (sc spaceCounter) CalculateDistance(currentCar model.Row, cars []model.Row, nearestCar map[int]model.Row) {
	nearLeft := struct {
		Distance float32

	}{}
	nearRigth := struct {
		Distance float32
	}{}

	for _, watchCar := range cars {
		if watchCar.Equal(currentCar) {
			continue
		}

		direction := ""
	}
}