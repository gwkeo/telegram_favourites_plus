package api

import (
	"fmt"
	"github.com/gwkeo/telegram_favourites_plus/pkg/api/json"
	"github.com/gwkeo/telegram_favourites_plus/pkg/config"
	"io"
	"net/http"
	"time"
)

func Connect() error {
	botApi, ok := config.GetBotApi()
	if ok != nil {
		return ok
	}

	var prevUpdateId int
	httpRequestDefault := "https://api.telegram.org/bot" + botApi + "/getUpdates?offset=-1"

	for {
		resp, err := http.Get(httpRequestDefault)
		if err != nil {
			return err
		}

		respBody, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return readErr
		}
		//fmt.Println(string(respBody))
		//
		response, ok := json.ParseGetUpdatesResponse(respBody)
		if ok != nil {
			return ok
		}
		if response.Result[0].UpdateId == prevUpdateId {
			continue
		} else {
			prevUpdateId = response.Result[0].UpdateId
			fmt.Println(response.Result[0].Message.Text)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
