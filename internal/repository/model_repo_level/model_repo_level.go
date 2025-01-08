package model_repo_level

import "database/sql"

//
//CREATE TABLE message_service.message (
//id UUID PRIMARY KEY,
//chat_id integer not null ,
//sender_id integer not null ,
//content TEXT NOT NULL,
//created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

const (
	SCHEMA       = "message_service"
	MessageTable = "message"
	ID           = "id"
	ChatID       = "chat_id"
	SenderID     = "sender_id"
	Content      = "content"
	CreatedAt    = "created_at"
	UpdatedAt    = "updated_at"
)

type MessageRepositoryLevel struct {
	ID        int          `json:"id" db:"id"`
	ChatID    int          `json:"chat_id" db:"chat_id"`
	SenderID  int          `json:"sender_id" db:"sender_id"`
	Content   string       `json:"content" db:"content"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
