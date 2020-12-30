package model

type Config struct {
	TelegramToken    string `json:"telegram_token"`
	PostgresLogin    string `json:"postgres_login"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDbName   string `json:"postgres_db_name"`
	PostgresPort     string `json:"postgres_port"`
	PostgresHost     string `json:"postgres_host"`
	OpenWeatherToken string `json:"open_weather_token"`
}
