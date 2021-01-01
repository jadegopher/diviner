package weather

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegram-pug/internal/app/handlers/weather/messages"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/tools/http/builder"
)

type weather struct {
	token string
}

func New(token string) repo.IHandler {
	return &weather{token: token}
}

func (w *weather) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	if update.Message.Location == nil {
		return nil, nil
	}
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
		msg.Text = messages.WeatherSuccess.English(weatherInfo.Weather[0].Description, weatherInfo.Info.Temp)
	} else {
		msg.Text = messages.WeatherErr.English()
	}
	return &msg, nil
}
