package model

type Coordinates struct {
	UserId    uint64  `json:"user_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Time      uint64  `json:"time"`
}
