package converter

import (
	"github.com/Coldwws/chat_practice/internal/model"
	desc "github.com/Coldwws/chat_practice/pkg/chat_v1"
	"time"
)

func SendMessageProtoToModel(req *desc.SendMessageRequest) *model.Message {
	var createdAt time.Time
	if req.Timestamp != nil {
		createdAt = req.Timestamp.AsTime()
	}
	return &model.Message{
		ChatID:    req.ChatId,
		Text:      req.Text,
		Sender:    req.From,
		CreatedAt: createdAt,
	}
}
