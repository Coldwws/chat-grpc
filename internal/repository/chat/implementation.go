package chat

import (
	"context"
	"fmt"
	"github.com/Coldwws/chat_practice/internal/client/db"
	"github.com/Coldwws/chat_practice/internal/model"
	"github.com/Coldwws/chat_practice/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type repo struct {
	db db.Client
}

func NewRepo(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, usernames []string) (int64, error) {

	qb := sq.Insert("chats").Columns("created_at").
		PlaceholderFormat(sq.Dollar).
		Values(sq.Expr("NOW()")).Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{Name: "chat_repository.Create", QueryRaw: query}
	var chatID int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	for _, username := range usernames {
		qbUser := sq.Insert("chat_users").Columns("chat_id", "username").
			Values(chatID, username).
			PlaceholderFormat(sq.Dollar)

		queryUser, argsUser, err := qbUser.ToSql()
		if err != nil {
			return 0, err
		}
		q2 := db.Query{Name: "chat_repository.Create", QueryRaw: queryUser}

		_, err = r.db.DB().ExecContext(ctx, q2, argsUser...)
		if err != nil {
			return 0, err
		}
	}
	return chatID, nil
}

func (r *repo) Delete(ctx context.Context, chatID int64) error {
	qbDelete := sq.Delete("chats").
		Where(sq.Eq{"id": chatID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := qbDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{Name: "chat_repository.Delete", QueryRaw: query}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	log.Printf("Chat with ID: %d deleted", chatID)

	return nil
}

func (r *repo) SendMessage(ctx context.Context, msg *model.Message) error {
	if msg == nil {
		return fmt.Errorf("nil message")
	}

	qb := sq.Insert("messages").
		Columns("chat_id", "sender", "text", "created_at").
		PlaceholderFormat(sq.Dollar).
		Values(msg.ChatID, msg.Sender, msg.Text, msg.CreatedAt)

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	log.Printf("Message with ID: %d sent", msg.ChatID)
	return nil
}
