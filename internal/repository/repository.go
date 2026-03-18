package repository

import (
	"context"
	"github.com/Coldwws/chat_practice/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, msg *model.Message) error
}
