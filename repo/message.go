package repo

type IMessages interface {
	CreateResponse(language string, args ...interface{}) string
}

type IMessage interface {
	English(args ...interface{}) string
	Russian(args ...interface{}) string
}
