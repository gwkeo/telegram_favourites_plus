package main

import (
	"github.com/gwkeo/telegram_favourites_plus/internal/api/telegram"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	"log"
)

func main() {
	apiKey, ok := utils.Api()
	if ok != nil {
		log.Fatal(ok.Error())
	}

	tg := telegram.Client{}
	tg.New(apiKey, "https://api.telegram.org/bot")
	err := tg.Run()

	if err != nil {
		log.Fatal(err)
	}
}
