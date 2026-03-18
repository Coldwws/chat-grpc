package chat

import (
	"context"
	"github.com/Coldwws/chat_practice/internal/model"
	"github.com/Coldwws/chat_practice/internal/repository"
)

type cService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) *cService {
	return &cService{chatRepository: chatRepository}
}

func (c cService) Create(ctx context.Context, usernames []string) (int64, error) {
	id, err := c.chatRepository.Create(ctx, usernames)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c cService) Delete(ctx context.Context, chatID int64) error {
	err := c.chatRepository.Delete(ctx, chatID)
	if err != nil {
		return err
	}
	return nil
}

func (c cService) SendMessage(ctx context.Context, msg *model.Message) error {
	err := c.chatRepository.SendMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
