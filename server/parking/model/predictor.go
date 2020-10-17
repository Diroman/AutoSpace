package model

import (
	pb "parking/internal/api"
)

type Point struct {
	X int
	Y int
}

func (p Point) GetDir(p Point)  {
	
}

type Row struct {
	Id     int
	Area   int64
	PointA Point
	PointB Point
}

func (r Row) Equal(other Row) bool {
	if r.Area != other.Area || r.PointA != other.PointA || r.PointB != other.PointB {
		return false
	}
	return true
}

type Class struct {
	Rows []Row
}

type Prediction struct {
	Classes map[string]Class
}

func PredictResponseToPrediction(resp *pb.Result) Prediction {
	prediction := Prediction{
		Classes: map[string]Class{},
	}

	id := 0
	newClass := Class{}
	for _, list := range resp.Classes {
		for _, row := range list.Data {
			pointA := row.Boxes[0]
			pointB := row.Boxes[1]
			newRow := Row{
				Id: id,
				Area: row.Area,
				PointA: Point{
					X: int(pointA.X),
					Y: int(pointA.Y),
				},
				PointB: Point{
					X: int(pointB.X),
					Y: int(pointB.Y),
				},
			}

			newClass.Rows = append(newClass.Rows, newRow)
			id += 1
		}
	}
	prediction.Classes["cars"] = newClass

	return prediction
}
