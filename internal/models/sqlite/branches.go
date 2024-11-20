package sqlite

import "database/sql"

type BranchesModel struct {
	DB *sql.DB
}

func (b *BranchesModel) Create(forumId, adminId int) error {
	stmt := "INSERT INTO branches (forumId, adminId) VALUES (?, ?)"
	_, err := b.DB.Exec(stmt, forumId, adminId)
	return err
}
