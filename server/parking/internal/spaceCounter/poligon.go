package spaceCounter

import (
	"log"

	"github.com/JamesLMilner/pip-go"

	"parking/model"
)

type ParkingPolygon struct {
	Vertices []model.Point `json:"vertices"`
}

type ParkingPolygons struct {
	Annotations []ParkingPolygon `json:"annotations"`
}

type ParkingSpace struct {
	Polygons []pip.Polygon
}

func (ps ParkingSpace) AddPolygon(points []int) {
	for _, point := range points {
		log.Println(point)
	}
}

func PointInPolygon() {

}
