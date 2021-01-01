package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var WoofResponse = message.New("Woof")
var AhchooResponse = message.New("Ahchoo")
var YipResponse = message.New("Yip")
var ChompResponse = message.New("Chomp")
var SniffResponse = message.New("Sniff sniff")

var DefaultResponse = messages.New(WoofResponse, AhchooResponse, YipResponse, ChompResponse, SniffResponse)
