package Processor

import (
	"context"
	"errors"
	tgClient "github.com/gwkeo/telegram_favourites_plus/internal/api/telegramApi"
	"github.com/gwkeo/telegram_favourites_plus/internal/db/repository/branch"
	"github.com/gwkeo/telegram_favourites_plus/internal/events"
	"github.com/gwkeo/telegram_favourites_plus/internal/handlers/telegramHandler"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/models/telegram"
	"github.com/gwkeo/telegram_favourites_plus/internal/utils"
)

type Processor struct {
	tg     tgClient.Client
	repo   branch.Repository
	offset int
}

func New(client tgClient.Client, repo branch.Repository) *Processor {
	return &Processor{
		tg:     client,
		repo:   repo,
		offset: 0,
	}
}

func (p *Processor) Start(ctx context.Context) error {

	jobs := make(chan []telegram.Result)

	go func() {

		for {
			select {
			case <-ctx.Done():
				return
			default:
				changes, err := p.fetchChanges()
				if err != nil {
					return
				}

				if len(changes) > 0 {
					jobs <- changes
				}
			}
		}
	}()

	for j := range jobs {
		for _, v := range j {
			eventType := telegramHandler.EventType(v)
			switch eventType {
			case events.MessageType:

				if err := p.forwardMessage(ctx, p.repo, v.Message); err != nil {
					return err
				}
			case events.BotAdded:
				branches, err := p.createInitialBranches(v.MyChatMember)
				if err != nil {
					return err
				}
				for _, br := range branches {
					if err = p.repo.Create(ctx, &br); err != nil {
						return err
					}
				}
			case events.Default:
				continue
			}
		}
	}

	return nil
}

func (p *Processor) fetchChanges() ([]telegram.Result, error) {
	body, err := p.tg.Updates(p.offset)
	if err != nil {
		return nil, errors.New("error while getting updates:\n" + err.Error())
	}
	res, err := utils.ParseUpdates(body)
	if err != nil {
		return nil, errors.New("error while parsing json response:\n" + err.Error())
	}

	if len(res.Result) > 0 {
		p.offset = res.Result[len(res.Result)-1].UpdateID + 1
	}

	return res.Result, nil
}

func (p *Processor) forwardMessage(ctx context.Context, repo branch.Repository, msg *telegram.Message) error {
	msgType := telegramHandler.MsgType(msg)

	br, err := repo.Branch(ctx, msg.Chat.ID, msgType)
	if err != nil {
		return err
	}

	request := telegram.Forward{
		ID:       msg.ID,
		ThreadId: br.ID,
		FromChat: msg.Chat.ID,
		Type:     msgType,
	}

	err = p.tg.ForwardMessage(request)
	return err
}

func (p *Processor) createInitialBranches(msg *telegram.MyChatMember) ([]models.Branch, error) {
	res := make([]models.Branch, len(models.TopicNames))

	for name, val := range models.TopicNames {
		br := &models.Branch{
			ForumID: msg.Chat.ID,
		}
		br.Type = val
		body, err := p.tg.CreateBranch(msg.Chat.ID, name)
		if err != nil {
			return nil, err
		}
		resp, err := utils.ParseCreated(body)
		if err != nil {
			return nil, err
		}
		br.ID = resp.Result.MessageThreadID
		res = append(res, *br)
	}
	return res, nil
}
