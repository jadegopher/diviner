package repo

type IMessages interface {
	CreateResponse(language string, args ...interface{}) string
}

type IMessage interface {
	Answer(language string, args ...interface{}) string
	English(args ...interface{}) string
	Russian(args ...interface{}) string
}
