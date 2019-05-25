package main

import (
	// System Modules
	"reflect"

	// App Modules
	"github.com/Noxdew/discordgo-starter-bot/bot"
	"github.com/Noxdew/discordgo-starter-bot/handlers"
)

func main() {
	b := bot.Create()
	defer b.Close()

	h := &handlers.Handlers{
		Bot: b,
	}

	handlerReflect := reflect.ValueOf(h)
	for i := 0; i < handlerReflect.NumMethod(); i++ {
		handler := handlerReflect.Method(i)
		b.AddHandler(handler.Interface())
	}

	b.Start()
}
