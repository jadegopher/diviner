package messages

import (
	"telegram-pug/internal/services/messages"
	"telegram-pug/internal/services/messages/message"
)

var InitConversation1 = message.New("Henlo %s... I'm the diviner pug."+
	" I forgot everything except for my name... I can't recall what has happened to me but I'm trying my best... I have lost all my skills of "+
	"the greatest diviner.... Oh wow, there is a scroll that keeps some forbidden acsients spells. Hmmmm, maybe I could "+
	"try some of them out? But I need your help. Tap the button when you will be ready. And one more thing. You can always ask me any questions. Just add '?' in the end of your message",
	`Примвет %s... Я мопс-предсказатель. Я забыл все, кроме своего имени... Я не могу вспомнить, что со мной случилось, но я пытаюсь... Я забыл, как предсказывать... Ой, смотри, там свиток древних со всякими магическими приколами. Хмммм, может, я смогу попробовать что-то скастовать? Но мне нужна твоя помощь. Нажми на кнопку, как будешь готов. Ой, и еще. Ты всегда можешь спросить у меня что-то, просто напиши в конце своего сообщения знак '?'`)

var InitConversation = messages.New(InitConversation1)
