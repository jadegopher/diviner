package model

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

type Config struct {
	TelegramToken    string `json:"telegram_token"`
	PostgresLogin    string `json:"postgres_login"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDbName   string `json:"postgres_db_name"`
	PostgresPort     string `json:"postgres_port"`
	PostgresHost     string `json:"postgres_host"`
	OpenWeatherToken string `json:"open_weather_token"`
}
