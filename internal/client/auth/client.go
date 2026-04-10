package auth

import (
	"context"

	access "github.com/Coldwws/chat_practice/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	accessClient access.AccessV1Client
}

func NewClient(authAddr string) (*Client, error) {
	conn, err := grpc.Dial(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		accessClient: access.NewAccessV1Client(conn),
	}, nil
}

func (c *Client) Check(ctx context.Context, endpointAddress string) error {
	_, err := c.accessClient.Check(ctx, &access.CheckRequest{
		EndpointAddress: endpointAddress,
	})
	return err
}