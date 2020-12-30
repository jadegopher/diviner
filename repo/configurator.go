package repo

type IConfigurator interface {
	Read() error
	Token() string
}
