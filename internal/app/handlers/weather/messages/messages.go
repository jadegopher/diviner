package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var WeatherSuccess1 = message.New(`Hmmmmm... I see %s and temperature %3.1f degrees and a lot of cupcakes... Ops sorry... Just lost my mind for a second`,
	`Хмммм... Я вижу, что у тебя там %s и температура %3.1f, а также много кексиков... Ой прости... Немного отвлекся`)
var WeatherSuccess2 = message.New(`Wow it works!!! %s overhead and temperature %3.1f degrees`,
	`Урааа, оно работает!!! %s над головой и температура %3.1f градусов`)
var WeatherSuccess3 = message.New(`Ha-ha not bad for me. My third eye says that you have %s overhead. Wait wait... Temperature  %3.1f degrees Celsius`,
	`Ха-ха, совсем неплохо. Мой третий глаз подсказывает мне, что у тебя %s над головой. Погоди погоди... Температура %3.1f градусов по Цельсию`)

var WeatherSuccess = messages.New(WeatherSuccess1, WeatherSuccess2, WeatherSuccess3)

var WeatherErr = message.New(`awwwwwww, sorry my magic doesn't work(('`,
	`блеп, что то свиток то не работает(`)
