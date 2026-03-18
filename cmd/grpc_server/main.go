package main

import (
	"context"
	newChatServer "github.com/Coldwws/chat_practice/internal/api/chat"
	"github.com/Coldwws/chat_practice/internal/config"
	chatRepo "github.com/Coldwws/chat_practice/internal/repository/chat"
	"github.com/Coldwws/chat_practice/internal/service/chat"
	desc "github.com/Coldwws/chat_practice/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	_ = godotenv.Load("local.env")

	cfg := config.LoadConfig()
	ctx := context.Background()

	lis, err := net.Listen("tcp", cfg.GRPC.Addr())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	conn, err := pgxpool.Connect(ctx, cfg.PG.DSN())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	repo := chatRepo.NewRepo(conn)

	chatSvc := chat.NewChatService(repo)
	s := newChatServer.NewChatServer(chatSvc)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	desc.RegisterChatV1Server(grpcServer, s)

	log.Printf("Server listening at addr: %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
