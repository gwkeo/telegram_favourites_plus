package main

import (
	"context"
	"database/sql"
	"github.com/gwkeo/telegram_favourites_plus/internal/api/telegram"
	"github.com/gwkeo/telegram_favourites_plus/internal/handlers"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
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

	tg := telegram.Client{
		ApiKey: apiKey,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	requests := make(chan *models.Request)

	go func() {
		offset := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				messages, err := tg.Updates(offset)
				if err != nil {
					log.Fatal("error while running client:\n" + err.Error())
				}
				if len(messages.Result) > 0 {
					handlers.HandleTelegramResponse(messages.Result, requests)

					offset = messages.Result[0].UpdateId + 1
				}
			}
		}

	}()

	go func() {
		for r := range requests {
			if ok := tg.ForwardMessage(r.ForwardToChat, r.FromChat, r.Id); ok != nil {
				log.Fatal("Error while handling request\n" + ok.Error())
			}
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

}
