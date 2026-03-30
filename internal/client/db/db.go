package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	SQLExecer
	Pinger
	Close()
	Transactor
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}
type Query struct {
	Name     string
	QueryRaw string
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}
type Pinger interface {
	Ping(ctx context.Context) error
}

// Интерфейс для работы с транзакциями.
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Менеджер транзакций,который выполняет указанный пользователем обработчик в транзакций.
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Handler func(ctx context.Context) error
