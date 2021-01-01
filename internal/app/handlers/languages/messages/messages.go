package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var InitConversation1 = message.New("Henlo %s... I'm the diviner pug." +
	" I forgot everything except my name... I can't recall what happened with me but I'm trying... I lost skills of " +
	"the greatest diviner.... Oh wow, there is a scroll of the acsients with some magic stuff. Hmmmm, may be I could " +
	"try to do some street magic? But I need your help.Tap the button when you will be ready")

var InitConversation = messages.New(InitConversation1)
