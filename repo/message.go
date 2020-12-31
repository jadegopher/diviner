package repo

type IMessage interface {
	English(args ...interface{}) string
}
