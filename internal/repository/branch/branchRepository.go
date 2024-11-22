package branch

import (
	"context"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
)

type Repository interface {
	Create(ctx context.Context, branch *models.Branch) error
	Branch(ctx context.Context, forumId int) (*models.Branch, error)
}
