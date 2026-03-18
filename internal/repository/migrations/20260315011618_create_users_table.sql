-- +goose Up
CREATE TABLE chats (
                       id BIGSERIAL PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE chat_users (
                            chat_id BIGINT NOT NULL,
                            username TEXT NOT NULL,
                            PRIMARY KEY (chat_id, username),
                            FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);

CREATE TABLE messages (
                          id BIGSERIAL PRIMARY KEY,
                          chat_id BIGINT NOT NULL,
                          sender TEXT NOT NULL,
                          text TEXT NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                          FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);

CREATE INDEX idx_messages_chat_id ON messages(chat_id);

-- +goose Down
DROP INDEX idx_messages_chat_id;
DROP TABLE messages;
DROP TABLE chat_users;
DROP TABLE chats;