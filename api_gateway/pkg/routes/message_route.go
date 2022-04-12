package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/toshkentov01/message-sender-bot/api_gateway/api/v1/controllers"
)

// MessageRoutes func for describe group of public routes.
func MessageRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Post("/message/send/", controllers.SendMessage)
	route.Post("/message/send-periodically/", controllers.SendMessagePeriodically)


}