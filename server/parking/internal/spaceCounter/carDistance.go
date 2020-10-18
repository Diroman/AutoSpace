package spaceCounter

import (
	"parking/model"
)

const angleRange = 15

type CarsNearest struct {
	Cars map[int][]CarDistance
}

func NewCarsNearest() *CarsNearest {
	return &CarsNearest{
		Cars: map[int][]CarDistance{},
	}
}

type CarDistance struct {
	Car       model.Row
	Distance  float64
	Direction float64
}

func (cn *CarsNearest) TryAddIfCloser(id int, distance, direction float64, addCar model.Row) {
	cars, ok := cn.Cars[id]
	if !ok {
		cn.Cars[id] = []CarDistance{{
			Car:       addCar,
			Distance:  distance,
			Direction: direction,
		}}
		return
	}

	for i, car := range cars {
		if car.Direction-angleRange < direction &&
			car.Direction+angleRange > direction {
			if car.Distance > distance {
				cars[i] = CarDistance{
					Car:       addCar,
					Distance:  distance,
					Direction: direction,
				}
			}
			return
		}
	}

	cn.Cars[id] = append(cars, CarDistance{
		Car:       addCar,
		Distance:  distance,
		Direction: direction,
	})
}

func (cn *CarsNearest) IsCloser(id int, addCar model.Row) bool {
	cars, ok := cn.Cars[addCar.Id]
	if !ok {
		return false
	}

	for _, car := range cars {
		if id == car.Car.Id {
			return true
		}
	}

	return false
}