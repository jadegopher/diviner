package file

import (
	"github.com/joho/godotenv"
	"os"
	"telegram-pug/model"
	"telegram-pug/repo"
)

type file struct {
	path   string
	config model.Config
}

func New(path string) repo.IConfigurator {
	return &file{
		path: path,
	}
}

func (f *file) Read() error {
	if err := godotenv.Load(f.path); err != nil {
		return err
	}
	f.config.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	f.config.PostgresLogin = os.Getenv("POSTGRES_LOGIN")
	f.config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	f.config.PostgresDbName = os.Getenv("POSTGRES_DB_NAME")
	f.config.PostgresHost = os.Getenv("POSTGRES_HOST")
	f.config.PostgresPort = os.Getenv("POSTGRES_PORT")
	f.config.OpenWeatherToken = os.Getenv("OPEN_WEATHER_TOKEN")

	return nil
}

func (f *file) TelegramToken() string {
	return f.config.TelegramToken
}

func (f *file) WeatherToken() string {
	return f.config.OpenWeatherToken
}

func (f *file) PostgresLogin() string {
	return f.config.PostgresLogin
}

func (f *file) PostgresPassword() string {
	return f.config.PostgresPassword
}

func (f *file) PostgresPort() string {
	return f.config.PostgresPort
}

func (f *file) PostgresDbName() string {
	return f.config.PostgresDbName
}

func (f *file) PostgresHost() string {
	return f.config.PostgresHost
}
