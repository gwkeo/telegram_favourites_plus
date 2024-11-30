package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
)

type BranchesRepository struct {
	DB *sql.DB
}

func New(DB *sql.DB) *BranchesRepository {
	return &BranchesRepository{DB: DB}
}

func (b BranchesRepository) Create(ctx context.Context, branch *models.Branch) error {

	stmt := "INSERT INTO branches (id, type, forumId) VALUES (?, ?, ?)"
	_, err := b.DB.ExecContext(ctx, stmt,
		branch.ID,
		branch.Type,
		branch.ForumID)
	if err != nil {
		return errors.New("Error creating branches: " + err.Error())
	}
	return nil
}

func (b BranchesRepository) Branch(ctx context.Context, forumId int, branchType models.Type) (*models.Branch, error) {
	rows, err := b.DB.QueryContext(ctx, "SELECT * FROM branches WHERE forumId = ? AND type = ?", forumId, branchType)
	if err != nil {
		return nil, errors.New("Error getting branch: " + err.Error())
	}

	branch := models.Branch{}
	for rows.Next() {
		if err := rows.Scan(
			&branch.ID,
			&branch.Type,
			&branch.ForumID,
		); err != nil {
			return nil, errors.New("Error parsing branch: " + err.Error())
		}
	}

	return &branch, nil
}
