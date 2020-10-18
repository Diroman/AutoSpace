package spaceCounter

import (
	"log"
	"math"
	"parking/model"
)

const (
	angle    float64 = 40
	height   float64 = 7
	camAngle float64 = 40
	xRes     float64 = 1920
	yRes     float64 = 1080
)

type spaceCounter struct{}

var SpaceCounter = spaceCounter{}

func (sc spaceCounter) GetSpaceCount(cars model.Prediction) (int, map[int]model.Row) {
	allCars := map[int]model.Row{}
	nearestCars := NewCarsNearest()

	for _, list := range cars.Classes {
		for _, car := range list.Rows {
			allCars[car.Id] = car
			sc.CalculateDistance(car, list.Rows, nearestCars)
		}
	}

	count := sc.CalculateSpaces(nearestCars, allCars)
	return count, allCars
}

func (sc spaceCounter) CalculateDistance(currentCar model.Row, cars []model.Row, nearestCar *CarsNearest) {
	for _, watchCar := range cars {
		if watchCar.Equal(currentCar) {
			continue
		}

		direction := currentCar.Middle.GetAngle(watchCar.Middle)
		distance := currentCar.Middle.GetDistance(watchCar.Middle)

		nearestCar.TryAddIfCloser(currentCar.Id, distance, direction, watchCar)
	}
}

func (sc spaceCounter) CalculateSpaces(cn *CarsNearest, allCars map[int]model.Row) int {
	watchedCars := map[int]bool{}
	count := 0

	for id, list := range cn.Cars {
		car := allCars[id]
		for _, nearCar := range list {
			if _, ok := watchedCars[nearCar.Car.Id]; ok {
				continue
			}

			distance := sc.GetDistanceInMeters(car.Middle, nearCar.Car.Middle)
			log.Println(id, " ", nearCar.Car.Id, " ", distance)

			if distance > 3 {
				log.Println(id, " ", nearCar.Car.Id, " ", distance)
				count += 1
			}
		}
		watchedCars[id] = true
	}

	return count
}

func (sc spaceCounter) GetDistanceInMeters(a, b model.Point) float64 {
	x1 := float64(a.X)
	y1 := float64(a.Y)
	x2 := float64(b.X)
	y2 := float64(b.Y)

	mid := max(y1, y2) - min(y1, y2)
	var angleMid = float64(angle - (camAngle / 2) + camAngle*(mid/yRes))

	var angle1 = float64(angle - (camAngle / 2) + camAngle*(y1/yRes))
	var angle2 = float64(angle - (camAngle / 2) + camAngle*(y2/yRes))

	yDist := math.Abs(math.Pow(math.Pow(height/math.Sin(angle1*math.Pi/180), 2)-math.Pow(height, 2), 0.5) -
		math.Pow(math.Pow(height/math.Sin(angle2*math.Pi/180), 2)-math.Pow(height, 2), 0.5))

	xxT := math.Abs(math.Pow(math.Pow(height/math.Sin((angleMid-camAngle/100)*math.Pi/180), 2)-math.Pow(height, 2), 0.5) -
		math.Pow(math.Pow(height/math.Sin((angleMid+camAngle/100)*math.Pi/180), 2)-math.Pow(height, 2), 0.5))

	xDist := (xxT / ((0.02 * camAngle) * yRes)) * math.Abs(x1-x2)

	return math.Pow(math.Pow(xDist, 2)+math.Pow(yDist, 2), 0.5)

}

func (sc spaceCounter) GetCarsCoordinate(allCars map[int]model.Row) []float64 {
	var points []interface{}
	for _, car := range allCars {
		points = append(points, float64(car.Middle.X))
		points = append(points, float64(car.Middle.Y))
	}

	return nil
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
