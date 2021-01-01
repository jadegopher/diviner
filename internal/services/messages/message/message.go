package message

import (
	"fmt"
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

func (m *message) English(args ...interface{}) string {
	return fmt.Sprintf(m.english, args...)
}

func (m *message) Russian(args ...interface{}) string {
	return fmt.Sprintf(m.russian, args...)
}
