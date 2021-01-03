package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var woofResponse = message.New("Woof", "Вуф")
var ahchooResponse = message.New("Ahchoo", "Апчхи")
var yipResponse = message.New("Yip", "Блеп")
var chompResponse = message.New("Chomp", "Фырк")
var sniffResponse = message.New("Sniff sniff", "Нюх нюх")

var DefaultResponse = messages.New(woofResponse, ahchooResponse, yipResponse, chompResponse, sniffResponse)

var yesResponse1 = message.New("I guess yes... Yeah, I am completely sure about that",
	"Я думаю, что да... Да, определенно уверен, что да")
var yesResponse2 = message.New("Blep... Oh I mean yes", "Блеп... Ой, прости, то есть да")
var yesResponse3 = message.New("Definitely yes", "Определенно да")

var YesResponse = messages.New(yesResponse1, yesResponse2, yesResponse3)

var noResponse1 = message.New("I think no", "Я думаю нет")
var noResponse2 = message.New("Mlem... Oh sorry I got distracted for a minute... No",
	"Млем... Ой, прости, отвлекся... Нет")
var noResponse3 = message.New("Definitely no", "Определенно нет")

var NoResponse = messages.New(noResponse1, noResponse2, noResponse3)
