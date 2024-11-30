package main

import (
	"context"
	"database/sql"
	"github.com/gwkeo/telegram_favourites_plus/internal/api/telegramApi"
	"github.com/gwkeo/telegram_favourites_plus/internal/db/repository/branch"
	"github.com/gwkeo/telegram_favourites_plus/internal/db/repository/branch/sqlite"
	"github.com/gwkeo/telegram_favourites_plus/internal/events/Processor"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")

	db, err := sql.Open("sqlite3", "./forum-branch.db")
	if err != nil {
		log.Fatal(err)
	}

	var repo branch.Repository = sqlite.New(db)
	client := telegramApi.New(apiKey)
	proc := Processor.New(*client, repo)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err = proc.Start(ctx); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down...")
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
