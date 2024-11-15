package main

import (
	"database/sql"
	"github.com/gwkeo/telegram_favourites_plus/internal/api/telegram"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, dbErr := sql.Open("sqlite3", "./forum_branches.db")
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println(db)

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
