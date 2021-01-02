package button

import (
	"telegram-pug/model"
	"telegram-pug/repo"
)

type button struct {
	english string
	russian string
	btnType model.Type
}

func New(english, russian string, btnType model.Type) repo.IButton {
	return &button{
		english: english,
		russian: russian,
		btnType: btnType,
	}
}

func (b *button) English() string {
	return b.english
}

func (b *button) Russian() string {
	return b.russian
}

func (b *button) Type() model.Type {
	return b.btnType
}
