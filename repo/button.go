package repo

import (
	"telegram-pug/model"
)

type IButton interface {
	English() string
	Russian() string
	Type() model.Type
}
