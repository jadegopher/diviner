package message

import (
	"fmt"
	"telegram-pug/constants"
	"telegram-pug/repo"
)

type message struct {
	english   string
	russian   string
	frequency uint8
}

func New(english, russian string) repo.IMessage {
	return &message{english: english, russian: russian}
}

func (m *message) Answer(language string, args ...interface{}) string {
	switch language {
	case constants.Eng:
		return m.English(args...)
	case constants.Ru:
		return m.Russian(args...)
	default:
		return m.English(args...)
	}
}

func (m *message) English(args ...interface{}) string {
	return fmt.Sprintf(m.english, args...)
}

func (m *message) Russian(args ...interface{}) string {
	return fmt.Sprintf(m.russian, args...)
}
