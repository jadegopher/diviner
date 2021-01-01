package model

type ICoordinates interface {
	CreateTableIfNotExists() error
	Select(userId uint64) (Coordinates, error)
	SelectMany(userId, limit, offset uint64) ([]Coordinates, error)
	Insert(coordinates Coordinates) (Coordinates, error)
	Delete(userId uint64) error
}

type Coordinates struct {
	UserId    uint64  `json:"user_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Time      uint64  `json:"time"`
}
