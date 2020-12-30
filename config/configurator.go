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
	f.config.Token = os.Getenv("TOKEN")
	return nil
}

func (f *file) Token() string {
	return f.config.Token
}
