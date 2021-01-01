package message

import (
	"fmt"
	"telegram-pug/repo"
)

type message struct {
	english string
}

func New(english string) repo.IMessage {
	return &message{english: english}
}

func (m *message) English(args ...interface{}) string {
	return fmt.Sprintf(m.english, args...)
}
