package sender

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/toshkentov01/message-sender-bot/message_service/config"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func MessageSender(content string, priority string) (err error) {
	var cfg = config.Get()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramApiToken)
	if err != nil {
		log.Println("Error while connection to a bot, error: ", err.Error())
		return err
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	log.Println(cfg.TelegamChannelUsername)

	msg := tgbotapi.NewMessageToChannel(cfg.TelegamChannelUsername, "Message: "+content+"\nPriority: "+priority)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Error while sending a message to a bot, error: ", err.Error())
		return err
	}

	return nil
}
