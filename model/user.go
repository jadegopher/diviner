package model

type IUser interface {
	CreateTableIfNotExists() error
	Select(id uint64) (*User, error)
	Insert(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(id uint64) error
}

type User struct {
	Id             uint64 `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Language       string `json:"language"`
	SaveLocation   bool   `json:"save_location"`
	LastTimeOnline int64  `json:"last_time_online"`
}
