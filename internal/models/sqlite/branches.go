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

func (b BranchesRepository) Create(ctx context.Context, branch *models.Branch) error {

	stmt := "INSERT INTO branches (forumId, adminId, textsBranchId, animationsBranchId, photosBranchId, documentsBranchId, videosBranchId, voicesBranchId, videoNotesBranchId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := b.DB.Exec(stmt,
		branch.ForumId,
		branch.AdminId,
		branch.TextsBranchId,
		branch.AnimationsBranchId,
		branch.PhotosBranchId,
		branch.DocumentsBranchId,
		branch.VideosBranchId,
		branch.VoicesBranchId,
		branch.VideoNotesBranchId)
	if err != nil {
		return errors.New("Error creating branches: " + err.Error())
	}
	return nil
}

func (b BranchesRepository) Branch(ctx context.Context, forumId int) (*models.Branch, error) {
	rows, err := b.DB.Query("SELECT * FROM branches WHERE forumId = ?", forumId)
	if err != nil {
		return nil, errors.New("Error getting branch: " + err.Error())
	}

	branch := models.Branch{}
	for rows.Next() {
		if err := rows.Scan(
			&branch.ForumId,
			&branch.AdminId,
			&branch.TextsBranchId,
			&branch.AnimationsBranchId,
			&branch.PhotosBranchId,
			&branch.DocumentsBranchId,
			&branch.VideosBranchId,
			&branch.VoicesBranchId,
			&branch.VideoNotesBranchId,
		); err != nil {
			return nil, errors.New("Error parsing branch: " + err.Error())
		}
	}

	return &branch, nil
}
