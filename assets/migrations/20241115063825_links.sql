-- +goose Up
-- +goose StatementBegin
CREATE TABLE branches (
    forumId INTEGER PRIMARY KEY,
    textsBranchId INTEGER,
    animationsBranchId INTEGER,
    photosBranchId INTEGER,
    documentsBranchId INTEGER,
    videosBranchId INTEGER,
    voicesBranchId INTEGER,
    videoNotesBranchId INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE branches;
-- +goose StatementEnd
