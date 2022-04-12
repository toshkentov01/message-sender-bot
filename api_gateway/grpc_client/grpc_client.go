package grpcclient

import (
	"fmt"
	"sync"

	"github.com/toshkentov01/message-sender-bot/api_gateway/config"

	message "github.com/toshkentov01/message-sender-bot/api_gateway/genproto/message_service"
	"google.golang.org/grpc"
)

var cfg = config.Config()
var (
	onceMessageService sync.Once

	instanceMessageService message.MessageServiceClient
)

// MessageService ...
func MessageService() message.MessageServiceClient {
	onceMessageService.Do(func() {
		connMessage, err := grpc.Dial(
			fmt.Sprintf("%s:%d", cfg.MessageServiceHost, cfg.MessageServicePort),
			grpc.WithInsecure())
		if err != nil {
			panic(fmt.Errorf("message service dial host: %s port:%d err: %s",
				cfg.MessageServiceHost, cfg.MessageServicePort, err))
		}

		instanceMessageService = message.NewMessageServiceClient(connMessage)
	})

	return instanceMessageService
}