package weather

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"telegram-pug/internal/app/handlers/menu/messages"
	"telegram-pug/internal/app/keyboards"
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

func (w *weather) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	user, err := w.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	query := map[string]string{
		"lat":   strconv.FormatFloat(update.Message.Location.Latitude, 'f', -1, 64),
		"lon":   strconv.FormatFloat(update.Message.Location.Longitude, 'f', -1, 64),
		"units": "metric",
		"appId": w.token,
		"lang":  strings.ToLower(user.Language),
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
		msg.Text = messages.WeatherSuccess.CreateResponse(user.Language, weatherInfo.Weather[0].Description,
			weatherInfo.Info.Temp)
	} else {
		msg.Text = messages.WeatherErr.English()
	}

	msg.ReplyMarkup = keyboards.MenuKeyboard.Keyboard(user.Language)
	return &msg, nil
}
