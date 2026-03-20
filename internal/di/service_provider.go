package di

import (
	"context"
	apiChat "github.com/Coldwws/chat_practice/internal/api/chat"
	"github.com/Coldwws/chat_practice/internal/client/db"
	"github.com/Coldwws/chat_practice/internal/client/db/pg"
	"github.com/Coldwws/chat_practice/internal/closer"
	"github.com/Coldwws/chat_practice/internal/config"
	"github.com/Coldwws/chat_practice/internal/repository"
	chatRepo "github.com/Coldwws/chat_practice/internal/repository/chat"
	"github.com/Coldwws/chat_practice/internal/service"
	chatServ "github.com/Coldwws/chat_practice/internal/service/chat"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	config *config.Config
	pgPool *pgxpool.Pool

	dbClient db.Client

	chatRepository repository.ChatRepository
	chatService    service.ChatService
	chatApi        *apiChat.Server
}

func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{
		config: cfg,
	}
}

var ctx = context.Background()

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.config == nil {
		log.Fatalf("config is nil")
	}
	return s.config.PG
}

func (s *serviceProvider) PGPool() *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.config.PG.DSN())
		if err != nil {
			log.Fatalf("failed to connect to database %s", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database %s", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}
	return s.pgPool
}

func (s *serviceProvider) DBClient() db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to init pg client %v", err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database %v", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) ChatRepository() repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewRepo(s.DBClient())
	}
	return s.chatRepository
}

func (s *serviceProvider) ChatService() service.ChatService {
	if s.chatService == nil {
		s.chatService = chatServ.NewChatService(s.ChatRepository())
	}
	return s.chatService
}
func (s *serviceProvider) ChatAPI() *apiChat.Server {
	if s.chatApi == nil {
		s.chatApi = apiChat.NewChatServer(s.ChatService())
	}
	return s.chatApi
}
