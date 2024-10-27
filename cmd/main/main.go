package main

import (
	"fmt"
	"github.com/gwkeo/telegram_favourites_plus/pkg/api"
)

func main() {
	ok := api.Connect()
	if ok != nil {
		fmt.Println("Error connecting to Telegram", ok)
	}

}
