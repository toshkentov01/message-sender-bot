package postgres

import (
	"fmt"
	"os"
	"testing"

	// "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/toshkentov01/message-sender-bot/message_service/config"
	"github.com/toshkentov01/message-sender-bot/message_service/pkg/logger"
	"github.com/toshkentov01/message-sender-bot/message_service/storage/repo"

	pb "github.com/toshkentov01/message-sender-bot/message_service/genproto/message_service"
)

var (
	// connDb  *sqlx.DB
	loggerr logger.Logger
	cfg     *config.Config
	mrRepo  repo.MessageRepository
)

func init() {
	const path = "/home/sardor/go/src/github.com/sardortoshkentov/message-sender-bot/message_service/.env"
	if info, err := os.Stat(path); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(path)
			if err != nil {
				fmt.Println("Err:", err)
			}
		}
	} else {
		fmt.Println("Not exists")
	}

	cfg = config.Get()
	loggerr = logger.New(cfg.LogLevel, "user_service")
	mrRepo = NewMessageRepository()
}

func TestSendMessage(t *testing.T) {

	_, err := mrRepo.SendMessage(&pb.Empty{})

	if err != nil {
		t.Error("Error while sending a message, error: ", err.Error())
	}

}
