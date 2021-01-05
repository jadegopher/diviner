package messages

import (
	"math/rand"
	"telegram-pug/repo"
)

type messages struct {
	messages []repo.IMessage
}

func New(message ...repo.IMessage) repo.IMessages {
	msgs := make([]repo.IMessage, 0, len(message))
	msgs = append(msgs, message...)
	return &messages{messages: msgs}
}

func (m *messages) CreateResponse(language string, args ...interface{}) string {
	n := rand.Intn(len(m.messages))
	return m.messages[n].Answer(language, args...)
}
