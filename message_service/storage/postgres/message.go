package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/toshkentov01/message-sender-bot/message_service/platform/postgres"
	"github.com/toshkentov01/message-sender-bot/message_service/storage/repo"

	pb "github.com/toshkentov01/message-sender-bot/message_service/genproto/message_service"
	sender "github.com/toshkentov01/message-sender-bot/message_service/pkg/sender"
)

type messageRepo struct {
	db *sqlx.DB
}

// NewMessageRepository ...
func NewMessageRepository() repo.MessageRepository {
	return &messageRepo{
		db: postgres.DB(),
	}
}

// SendMessage ...
func (mr *messageRepo) SendMessage(*pb.Empty) (*pb.Empty, error) {
	messages := []*pb.Message{}

	rows, err := mr.db.Query(GetMessagesByPrioritySQL)
	if err != nil {
		log.Println("Error while getting messages by priority, error: ", err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		message := pb.Message{}

		err := rows.Scan(
			&message.Id,
			&message.Content,
			&message.Priority,
		)
		if err != nil {
			log.Println("Error while scanning, error: ", err.Error())
			return nil, err
		}

		messages = append(messages, &message)
	}

	for _, data := range messages {
		err := sender.MessageSender(data.Content, data.Priority)
		if err != nil {
			log.Println("Error while sending a message, error: ", err.Error())
			return nil, err
		}
	}

	return &pb.Empty{}, nil
}
