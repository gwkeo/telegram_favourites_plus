package telegram

import (
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	"net/http"
)

const (
	updatesResponsePath = "getUpdates"
	timeout             = "30"
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

func (c *Client) Updates() (models.Response, error) {
	requestPath := c.Request() + "/" + updatesResponsePath + "?offset=277446911&timeout=" + timeout
	resp, err := http.Get(requestPath)
	if err != nil {
		return models.Response{}, errors.New("error while getting updates: " + err.Error())
	}
	defer resp.Body.Close()
	var respBody []byte
	_, err = resp.Body.Read(respBody)
	if err != nil {
		return models.Response{}, errors.New("error while reading json response: " + err.Error())
	}

	res, err := utils.Response(respBody)
	if err != nil {
		return models.Response{}, errors.New("error while parsing json response: " + err.Error())
	}
	return res, nil
}
