package messages

import (
	"math/rand"
	"telegram-pug/repo"
)

func Response(language string) string {
	resp := []repo.IMessage{
		WoofResponse,
		AhchooResponse,
		YipResponse,
		ChompResponse,
		SniffResponse,
	}
	n := rand.Intn(len(resp))
	switch language {
	default:
		return resp[n].English()
	}
}
