package model

type ICompliments interface {
	CreateTableIfNotExists() error
	Insert(compliment Compliment) (*Compliment, error)
	SelectRandom(gender string) (*Compliment, error)
}

type Gender int

const (
	Male Gender = iota
	Female
	Neutral
)

type Compliment struct {
	Id      uint64 `json:"id"`
	English string `json:"english"`
	Russian string `json:"russian"`
	Gender  string `json:"gender"`
}

func (d Gender) String() string {
	return [...]string{"male", "female", "neutral"}[d]
}
