package weather

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegram-pug/model"
	"telegram-pug/tools/http/builder"
)

func Handle(token string, update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	query := map[string]string{
		"lat":   strconv.FormatFloat(update.Message.Location.Latitude, 'f', -1, 64),
		"lon":   strconv.FormatFloat(update.Message.Location.Longitude, 'f', -1, 64),
		"units": "metric",
		"appId": token,
	}
	req, err := builder.New(http.MethodGet, "http://api.openweathermap.org/data/2.5/weather",
		builder.WithQuery(query))
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	resp, err := req.Do()
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	weatherInfo := &model.OpenWeather{}
	if err := json.Unmarshal(body, weatherInfo); err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	if len(weatherInfo.Weather) != 0 {
		msg.Text = fmt.Sprintf(`Hmmmmm... I see %s and temperature %3.1f degrees. Not bad for a pug who lost his memory`,
			weatherInfo.Weather[0].Description, weatherInfo.Info.Temp)

	} else {
		msg.Text = `awwwwwww, sorry my magic doesn't work(('`
	}
	return msg, nil
}
