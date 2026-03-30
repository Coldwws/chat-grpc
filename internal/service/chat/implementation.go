package chat

import (
	"context"
	"github.com/Coldwws/chat_practice/internal/client/db"
	"github.com/Coldwws/chat_practice/internal/model"
	"github.com/Coldwws/chat_practice/internal/repository"
)

type cService struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewChatService(chatRepository repository.ChatRepository, txManager db.TxManager) *cService {
	return &cService{chatRepository: chatRepository, txManager: txManager}
}

func (c *cService) Create(ctx context.Context, usernames []string) (int64, error) {
	var id int64
	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = c.chatRepository.Create(ctx, usernames)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	return id, err
}

func (c *cService) Delete(ctx context.Context, chatID int64) error {
	return c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		return c.chatRepository.Delete(ctx, chatID)
	})
}

func (c *cService) SendMessage(ctx context.Context, msg *model.Message) error {
	return c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		return c.chatRepository.SendMessage(ctx, msg)
	})
}
