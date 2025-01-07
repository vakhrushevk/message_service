package main

import (
	"context"
	"fmt"
	"github.com/vakhrushevk/local-platform/db/pg"
	"github.com/vakhrushevk/message_service/internal/repository/model_repo_level"
	"github.com/vakhrushevk/message_service/internal/repository/postgres"
	"log"
)

const DSN = "host=localhost port=5555 dbname=chat user=chat-user password=qwerty sslmode=disable"

func main() {
	client, err := pg.New(context.Background(), DSN)
	if err != nil {
		log.Fatalf("failed to create db client: %v", err)
	}
	err = client.DB().Ping(context.TODO())
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}

	rep := postgres.NewRepository(client)
	mrl := model_repo_level.MessageRepositoryLevel{
		ChatID:   1,
		SenderID: 1,
		Content:  "Hello, World!",
	}

	fmt.Println("Creating message...")
	k, err := rep.CreateMessage(context.TODO(), &mrl)
	if err != nil {
		log.Fatalf("failed to create message: %v", err)
	}
	fmt.Println("Created message with id:", k)

	fmt.Println("Getting message...")
	msg, err := rep.GetMessage(context.TODO(), 1)
	if err != nil {
		fmt.Println("error getting message:", err)
	}
	fmt.Println("Message:", msg)

	fmt.Println("Getting messages...")
	msgs, err := rep.GetMessages(context.TODO(), 1)
	if err != nil {
		fmt.Println("error getting messages:", err)
	}
	for _, v := range msgs {
		fmt.Println(v)
	}

}
