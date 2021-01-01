package messages

import (
	"math/rand"
	"telegram-pug/constants"
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
	switch language {
	case constants.Eng:
		return m.messages[n].English(args...)
	case constants.Ru:
		return m.messages[n].Russian(args...)
	default:
		return m.messages[n].English(args...)
	}
}
