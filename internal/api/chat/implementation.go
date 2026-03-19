package chat

import (
	"context"
	"github.com/Coldwws/chat_practice/internal/converter"
	desc "github.com/Coldwws/chat_practice/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatID, err := s.chatService.Create(ctx, req.Usernames)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.CreateResponse{Id: chatID}, nil
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	msg := converter.SendMessageProtoToModel(req)
	err := s.chatService.SendMessage(ctx, msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
