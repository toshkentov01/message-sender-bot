package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/toshkentov01/message-sender-bot/message_service/service"
	"github.com/toshkentov01/message-sender-bot/message_service/config"
	"github.com/toshkentov01/message-sender-bot/message_service/pkg/logger"

	pb "github.com/toshkentov01/message-sender-bot/message_service/genproto/message_service"

	"google.golang.org/grpc"

)

func main() {
	if info, err := os.Stat(".env"); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(".env")
		}
	}

	var cfg = config.Get()

	log := logger.New(cfg.LogLevel, "message_service")
	defer logger.Cleanup(log)

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error listening tcp port: ", logger.Error(err))
	}

	messageService := service.NewMessageService(log)
	
	server := grpc.NewServer()
	pb.RegisterMessageServiceServer(server, messageService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Fatal("error listening: %v", logger.Error(err))
	}
}
