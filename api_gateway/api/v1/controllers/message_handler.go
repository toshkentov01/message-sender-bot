package controllers

import (
	"context"
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/toshkentov01/message-sender-bot/api_gateway/api/v1/models"
	"github.com/toshkentov01/message-sender-bot/api_gateway/config"
	client "github.com/toshkentov01/message-sender-bot/api_gateway/grpc_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/toshkentov01/message-sender-bot/api_gateway/genproto/message_service"
)

// SendMessage api to send a message.
// @Description SendMessage API used for sending a message to a channel via telegram bot
// @Tags Message-Sender
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/message/send/ [post]
func SendMessage(c *fiber.Ctx) error {
	var (
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	_, err = client.MessageService().SendMessage(ctx, &pb.Empty{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok || st.Code() == codes.Internal {
			return c.Status(http.StatusInternalServerError).JSON(
				models.Response{
					Error: true,
					Data: models.Error{
						Status:  "Internal Server Error",
						Message: st.Message(),
					},
				},
			)
		} else if st.Code() == codes.InvalidArgument {
			return c.Status(http.StatusBadRequest).JSON(models.Response{
				Error: true,
				Data: models.Error{
					Status: "Bad Request",
					Message: "API TOKEN OR CHANNEL USERNAME IS INVALID",
				},
			})
		}
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Error: false,
		Data: models.Success{
			Success: true,
		},
	})
}
