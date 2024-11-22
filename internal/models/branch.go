package models

import "database/sql"

type Branch struct {
	ForumId            sql.NullInt64
	AdminId            sql.NullInt64
	TextsBranchId      sql.NullInt64
	AnimationsBranchId sql.NullInt64
	PhotosBranchId     sql.NullInt64
	DocumentsBranchId  sql.NullInt64
	VideosBranchId     sql.NullInt64
	VoicesBranchId     sql.NullInt64
	VideoNotesBranchId sql.NullInt64
}
