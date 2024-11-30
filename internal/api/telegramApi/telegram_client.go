package telegramApi

import (
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/models/telegram"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	host                       = "api.telegram.org"
	updatesResponsePath        = "getUpdates"
	forwardMessageResponsePath = "forwardMessage"
	createForumTopicPath       = "createForumTopic"
	timeout                    = 10
)

type Client struct {
	apiKey string
}

func (c *Client) basePath() string {
	return "bot" + c.apiKey
}

// Updates делает запрос getUpdates в telegramApi api для получения всех новых сообщений за последний период таймаута
func (c *Client) Updates(offset int) ([]byte, error) {

	queries := url.Values{}
	queries.Add("offset", strconv.Itoa(offset))
	queries.Add("timeout", strconv.Itoa(timeout))

	request := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(c.basePath(), updatesResponsePath),
	}

	request.RawQuery = queries.Encode()

	resp, err := http.Get(request.String())
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

func (c *Client) ForwardMessage(r telegram.Forward) error {
	queries := url.Values{}
	queries.Add("message_thread_id", strconv.Itoa(r.ThreadId))
	queries.Add("chat_id", strconv.Itoa(r.FromChat))
	queries.Add("from_chat_id", strconv.Itoa(r.FromChat))
	queries.Add("message_id", strconv.Itoa(r.ID))

	request := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(c.basePath(), forwardMessageResponsePath),
	}

	request.RawQuery = queries.Encode()

	_, err := http.Post(request.String(), "", nil)
	if err != nil {
		return errors.New("error while forwarding message:\n" + err.Error())
	}

	return nil
}

func (c *Client) CreateBranch(chatID int, name string) ([]byte, error) {
	queries := url.Values{}
	queries.Add("chat_id", strconv.Itoa(chatID))
	queries.Add("name", name)

	request := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(c.basePath(), createForumTopicPath),
	}

	queries.Add("topic_name", name)
	request.RawQuery = queries.Encode()
	v, err := http.Post(request.String(), "", nil)
	if err != nil {
		return nil, errors.New("error while creating initial branches:\n" + err.Error())
	}
	defer v.Body.Close()
	body, err := io.ReadAll(v.Body)
	if err != nil {
		return nil, errors.New("error while parsing response of creating initial branches process:\n" + err.Error())
	}

	return body, nil
}
