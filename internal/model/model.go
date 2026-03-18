package model

import "time"

type Chat struct {
	ID        int64
	CreatedAt time.Time
}

type ChatUser struct {
	ChatID   int64
	Username string
}

type Message struct {
	ID        int64
	ChatID    int64
	Sender    string
	Text      string
	CreatedAt time.Time
}
