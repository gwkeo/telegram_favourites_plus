package telegram

import (
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/handlers"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
	"io"
	"net/http"
	"strconv"
)

const (
	updatesResponsePath        = "getUpdates"
	forwardMessageResponsePath = "forwardMessage"
	timeout                    = 10
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

func (c *Client) RequestForwardMessagePath(chatId, fromChatId, messageId int) string {
	return c.Request() +
		"/" + forwardMessageResponsePath +
		"?chat_id=" + strconv.Itoa(chatId) +
		"&from_chat_id=" + strconv.Itoa(fromChatId) +
		"&message_id=" + strconv.Itoa(messageId)
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

func (c *Client) ForwardMessage(chatId, fromChatId, messageId int) error {
	requestPath := c.RequestForwardMessagePath(chatId, fromChatId, messageId)
	_, err := http.Post(requestPath, "", nil)
	if err != nil {
		return errors.New("error while forwarding message:\n" + err.Error())
	}
	return nil
}

func (c *Client) MakeRequests(requests []models.Request) error {
	for _, v := range requests {
		err := c.ForwardMessage(v.ForwardToChat, v.FromChat, v.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Run() error {
	clientRunErrorMsg := "error while running client:\n"
	offset := 0
	for {
		messages, err := c.Updates(offset)
		if err != nil {
			return errors.New(clientRunErrorMsg + err.Error())
		}
		if len(messages.Result) > 0 {
			requests := handlers.HandleTelegramResponse(messages.Result)
			ok := c.MakeRequests(requests)
			if ok != nil {
				return errors.New(clientRunErrorMsg + err.Error())
			}
			if messages.Result != nil && len(messages.Result) > 0 {
				offset = messages.Result[0].UpdateId + 1
			}
		}
	}
}
