package repo

type IConfigurator interface {
	Read() error
	Token() string
	PostgresLogin() string
	PostgresPassword() string
	PostgresPort() string
	PostgresDbName() string
	PostgresHost() string
}
