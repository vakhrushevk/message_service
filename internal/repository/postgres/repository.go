package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/vakhrushevk/local-platform/db"
	"github.com/vakhrushevk/message_service/internal/repository"
	mrl "github.com/vakhrushevk/message_service/internal/repository/model_repo_level"
)

type messageRepository struct {
	db db.Client
}

// GetMessages получить все сообщения
func (m *messageRepository) GetMessages(ctx context.Context, chatId int64) ([]*mrl.MessageRepositoryLevel, error) {
	query, args, err :=
		squirrel.Select(mrl.ID,
			mrl.ChatID,
			mrl.SenderID,
			mrl.Content,
			mrl.CreatedAt,
			mrl.UpdatedAt).
			From(mrl.SCHEMA + "." + mrl.MessageTable).
			Where(squirrel.Eq{mrl.ChatID: chatId}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		// TODO: Add logger
		return nil, err
	}

	q := db.Query{
		Name:     "Get Messages",
		QueryRaw: query,
	}
	var messages []*mrl.MessageRepositoryLevel
	err = m.db.DB().ScanAllContext(ctx, &messages, q, args...)
	if err != nil {
		// TODO: Add logger
		return nil, err
	}

	return messages, nil
}

// GetMessage получить сообщение по id
func (m *messageRepository) GetMessage(ctx context.Context, id int64) (*mrl.MessageRepositoryLevel, error) {
	query, args, err :=
		squirrel.Select(mrl.ID,
			mrl.ChatID,
			mrl.SenderID,
			mrl.Content,
			mrl.CreatedAt,
			mrl.UpdatedAt).
			From(mrl.SCHEMA + "." + mrl.MessageTable).
			Where(squirrel.Eq{mrl.ID: id}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		// todo: add logger
		return nil, err
	}

	q := db.Query{
		Name:     "Get Message",
		QueryRaw: query,
	}

	var message mrl.MessageRepositoryLevel
	err = m.db.DB().ScanOneContext(ctx, &message, q, args...)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// CreateMessage создать новое сообщение
func (m *messageRepository) CreateMessage(ctx context.Context, message *mrl.MessageRepositoryLevel) (int64, error) {
	query, args, err :=
		squirrel.Insert(mrl.SCHEMA+"."+mrl.MessageTable).
			Columns(mrl.ChatID, mrl.SenderID, mrl.Content).
			Values(message.ChatID, message.SenderID, message.Content).
			PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		// TODO: Add logger
		return 0, nil
	}

	q := db.Query{
		Name:     "Create Message",
		QueryRaw: query,
	}

	_, err = m.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		// TODO: Add logger
		return 0, err
	}

	// TODO: Переделать что бы возвращал id созданного сообщения
	return 0, nil
}

func NewRepository(db db.Client) repository.MessageRepository {
	return &messageRepository{db: db}
}
