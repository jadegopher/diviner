package weather

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegram-pug/internal/app/handlers/weather/messages"
	"telegram-pug/internal/services/users"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/tools/http/builder"
	"telegram-pug/usecases"
)

type weather struct {
	userService usecases.IUserService
	token       string
}

func New(dbConn *gorm.DB, token string) (repo.IHandler, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &weather{userService: db, token: token}, nil
}

func (w *weather) Condition(update tgbotapi.Update) bool {
	if update.Message.Location == nil {
		return false
	}
	return true
}

func (w *weather) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	query := map[string]string{
		"lat":   strconv.FormatFloat(update.Message.Location.Latitude, 'f', -1, 64),
		"lon":   strconv.FormatFloat(update.Message.Location.Longitude, 'f', -1, 64),
		"units": "metric",
		"appId": w.token,
	}
	req, err := builder.New(http.MethodGet, "http://api.openweathermap.org/data/2.5/weather",
		builder.WithQuery(query))
	if err != nil {
		return nil, err
	}
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weatherInfo := &model.OpenWeather{}
	if err := json.Unmarshal(body, weatherInfo); err != nil {
		return nil, err
	}
	if len(weatherInfo.Weather) != 0 {
		user, err := w.userService.UserInfo(update)
		if err != nil {
			return nil, err
		}
		msg.Text = messages.WeatherSuccess.CreateResponse(user.Language, weatherInfo.Weather[0].Description,
			weatherInfo.Info.Temp)
	} else {
		msg.Text = messages.WeatherErr.English()
	}
	return &msg, nil
}
