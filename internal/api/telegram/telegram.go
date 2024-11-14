package telegram

import (
	"errors"
	"fmt"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	"io"
	"net/http"
	"strconv"
)

const (
	updatesResponsePath = "getUpdates"
	timeout             = 2
)

type Client struct {
	apiKey  string
	baseUrl string
}

func (c *Client) New(apiKey string, baseUrl string) {
	c.apiKey = apiKey
	c.baseUrl = baseUrl
}

func (c *Client) Request() string {
	return c.baseUrl + c.apiKey
}

func (c *Client) RequestUpdatesPath(offset int) string {
	return c.Request() + "/" + updatesResponsePath + "?offset=" + strconv.Itoa(offset) + "&timeout=" + strconv.Itoa(timeout)
}

// LastMessages делает запрос getUpdates в telegram api для получения всех новых сообщений за последний период таймаута
func (c *Client) LastMessages(offset int) ([]byte, error) {
	requestPath := c.RequestUpdatesPath(offset)
	resp, err := http.Get(requestPath)
	if err != nil {
		return nil, errors.New("error while getting updates:\n" + err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error while reading json response:\n" + err.Error())
	}
	return body, nil
}

func (c *Client) Updates(offset int) (*models.Response, error) {

	body, err := c.LastMessages(offset)
	if err != nil {
		return nil, errors.New("error while getting updates:\n" + err.Error())
	}
	res, err := utils.Response(body)
	if err != nil {
		return nil, errors.New("error while parsing json response:\n" + err.Error())
	}
	return res, nil
}

func (c *Client) Run() error {
	offset := 0
	for {
		messages, err := c.Updates(offset)
		if err != nil {
			return errors.New("error while running client:\n" + err.Error())
		}
		fmt.Println(messages.Ok)
		offset = messages.Result[0].UpdateId + 1
		fmt.Println(messages.Result[0].Message.Text, offset)
	}
}
