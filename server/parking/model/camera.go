package model

type CameraInfo struct {
	Height       float32
	Latitude     float64
	Longitude    float64
	HorizonAngle int
}

type Camera struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type Cameras struct {
	Cameras []Camera `json:"cameras"`
}
