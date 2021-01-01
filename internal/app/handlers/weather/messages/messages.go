package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var WeatherSuccess1 = message.New(`Hmmmmm... I see %s and temperature %3.1f degrees and a lot of cupcakes... Ops sorry... Just lost my mind for a second`)
var WeatherSuccess2 = message.New(`Wow it works!!! %s overhead and temperature %3.1f degrees`)
var WeatherSuccess3 = message.New(`Ha-ha not bad for me. My third eye says that you have %s overhead. Wait wait... Temperature  %3.1f degrees Celsius`)

var WeatherSuccess = messages.New(WeatherSuccess1, WeatherSuccess2, WeatherSuccess3)

var WeatherErr = message.New(`awwwwwww, sorry my magic doesn't work(('`)
