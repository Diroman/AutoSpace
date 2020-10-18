package model

type Camera struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type Cameras struct {
	Cameras []Camera `json:"cameras"`
}

