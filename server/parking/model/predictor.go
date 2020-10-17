package model

import (
	pb "parking/internal/api"
)

type Point struct {
	X int
	Y int
}

type Row struct {
	Area   int64
	PointA Point
	PointB Point
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

	for class, list := range resp.Classes {
		newClass := Class{}
		for _, row := range list.Data {
			pointA := row.Boxes[0]
			pointB := row.Boxes[1]
			newRow := Row{
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
		}
		prediction.Classes[class] = newClass
	}

	return prediction
}
