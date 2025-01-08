-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS message_service;
CREATE TABLE message_service.message (
                          id SERIAL PRIMARY KEY,
                          chat_id integer not null ,
                          sender_id integer not null ,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message_service.message;
-- +goose StatementEnd