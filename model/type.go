package model

type Type int

const (
	Default Type = iota
	Location
)

func (d Type) String() string {
	return [...]string{"Default", "Location"}[d]
}
