package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	BotApiKey string `json:"bot_api_key"`
}

func GetBotApi() (string, error) {
	jsonFile, ok := os.Open("../../config.json")
	if ok != nil {
		return "", ok
	}

	defer func() {
		closeError := jsonFile.Close()
		if closeError != nil {
			log.Fatal(closeError)
		}
	}()

	byteValue, _ := io.ReadAll(jsonFile)
	var config Config

	if ok = json.Unmarshal(byteValue, &config); ok != nil {
		return "", ok
	}

	return config.BotApiKey, nil
}
