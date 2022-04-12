package repo

import (
	pb "github.com/toshkentov01/message-sender-bot/message_service/genproto/message_service"
)


// MessageRepository ...
type MessageRepository interface {
	SendMessage(*pb.Empty) (*pb.Empty, error)
}