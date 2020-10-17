package spaceCounter

import (
	"fmt"
	"parking/model"
)

type spaceCounter struct{}

var SpaceCounter = spaceCounter{}

func (sc spaceCounter) GetSpaceCount(cars model.Prediction) int {
	allCars := map[int]model.Row{}
	nearestCars := NewCarsNearest()

	for _, list := range cars.Classes {
		for _, car := range list.Rows {
			allCars[car.Id] = car
			sc.CalculateDistance(car, list.Rows, nearestCars)
		}
	}

	count := sc.CalculateSpaces(nearestCars, allCars)
	return count
}

func (sc spaceCounter) CalculateDistance(currentCar model.Row, cars []model.Row, nearestCar *CarsNearest) {
	for _, watchCar := range cars {
		if watchCar.Equal(currentCar) {
			continue
		}

		direction := currentCar.Middle.GetAngle(watchCar.Middle)
		distance := currentCar.Middle.GetDistance(watchCar.Middle)

		//isCloser := nearestCar.IsCloser(currentCar.Id, watchCar)
		nearestCar.TryAddIfCloser(currentCar.Id, distance, direction, watchCar)
	}
}

func (sc spaceCounter) CalculateSpaces(cn *CarsNearest, allCars map[int]model.Row) int {
	watchedCars := map[int]bool{}
	count := 0

	for id, list := range cn.Cars {
		_ = allCars[id]
		println(id)
		for _, nearCar := range list {
			if _, ok := watchedCars[id]; ok {
				continue
			}

			fmt.Println(nearCar)
			count += 1
		}
		watchedCars[id] = true
	}

	return count
}