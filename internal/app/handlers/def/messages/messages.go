package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var WoofResponse = message.New("Woof", "Вуф")
var AhchooResponse = message.New("Ahchoo", "Апчхи")
var YipResponse = message.New("Yip", "Блеп")
var ChompResponse = message.New("Chomp", "Фырк")
var SniffResponse = message.New("Sniff sniff", "Нюх нюх")

var DefaultResponse = messages.New(WoofResponse, AhchooResponse, YipResponse, ChompResponse, SniffResponse)
