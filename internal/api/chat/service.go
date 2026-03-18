package chat

import (
	"github.com/Coldwws/chat_practice/internal/service"
	desc "github.com/Coldwws/chat_practice/pkg/chat_v1"
)

type server struct {
	chatService service.ChatService
	desc.UnimplementedChatV1Server
}

func NewChatServer(chatService service.ChatService) *server {
	return &server{
		chatService: chatService,
	}
}
