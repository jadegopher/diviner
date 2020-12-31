package message

type message struct {
	english string
}

func (m *message) English() string {
	return m.english
}
