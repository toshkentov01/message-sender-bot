package controllers

import (
	"context"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	pb "github.com/toshkentov01/message-sender-bot/api_gateway/genproto/message_service"
	client "github.com/toshkentov01/message-sender-bot/api_gateway/grpc_client"
)

// SendMessagePeriodically api to send a message periodically.
// @Description SendMessagePeriodically API used for sending a message to a channel via telegram bot periodically
// @Description THIS API WILL NEVER FINISHE ITS WORK
// @Tags Message-Sender
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/message/send-periodically/ [post]
func SendMessagePeriodically(c *fiber.Ctx) error{
	for {
		var (
			err error
		)

		_, err = client.MessageService().SendMessage(context.Background(), &pb.Empty{})
		if err != nil {
			log.Println("Error while sending a message in API, error: ", err.Error())
		}

		<-time.After(5 * time.Second)
	}
}
