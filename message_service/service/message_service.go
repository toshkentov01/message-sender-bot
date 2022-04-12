package service

import (
	"context"

	l "github.com/toshkentov01/message-sender-bot/message_service/pkg/logger"
	"github.com/toshkentov01/message-sender-bot/message_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/toshkentov01/message-sender-bot/message_service/genproto/message_service"
)

// MessageService ...
type MessageService struct {
	storage storage.I
	logger  l.Logger
}

// NewMessageService ...
func NewMessageService(log l.Logger) *MessageService {
	return &MessageService{
		storage: storage.NewStorage(),
		logger:  log,
	}
}

// SendMessage ... 
func (ms *MessageService) SendMessage(ctx context.Context, request *pb.Empty) (*pb.Empty, error) {
	result, err := ms.storage.Message().SendMessage(request)
	if err != nil {
		ms.logger.Error("Error while sending a message, error: "+err.Error())
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	return result, nil
}