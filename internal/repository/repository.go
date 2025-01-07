package repository

import (
	"context"
	"github.com/vakhrushevk/message_service/internal/repository/model_repo_level"
)

type MessageRepository interface {
	// GetMessages returns all messages
	GetMessages(ctx context.Context, chatId int64) ([]*model_repo_level.MessageRepositoryLevel, error)
	// GetMessage returns a message by id
	GetMessage(ctx context.Context, id int64) (*model_repo_level.MessageRepositoryLevel, error)
	// CreateMessage creates a new message
	CreateMessage(ctx context.Context, message *model_repo_level.MessageRepositoryLevel) (int64, error)
}
