package storage

import (
	"github.com/toshkentov01/message-sender-bot/message_service/storage/repo"
	"github.com/toshkentov01/message-sender-bot/message_service/storage/postgres"
)

// I is an interface for storage
type I interface {
	Message() repo.MessageRepository
}

type storage struct {
	messageRepo repo.MessageRepository
}

// NewStorage ...
func NewStorage() I {
	return &storage{
		messageRepo: postgres.NewMessageRepository(),
	}
}

func (s storage) Message() repo.MessageRepository {
	return s.messageRepo
}
