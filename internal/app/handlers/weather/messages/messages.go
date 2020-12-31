package messages

import "telegram-pug/internal/message"

var WeatherSuccess = message.New(`Hmmmmm... I see %s and temperature %3.1f degrees. Not bad for a pug who lost his memory`)

var WeatherErr = message.New(`awwwwwww, sorry my magic doesn't work(('`)
