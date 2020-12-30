package repo

type IConfigurator interface {
	Read() error
	TelegramToken() string
	WeatherToken() string
	PostgresLogin() string
	PostgresPassword() string
	PostgresPort() string
	PostgresDbName() string
	PostgresHost() string
}
