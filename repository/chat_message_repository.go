package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatMessageRepositoryInterface interface {
	Insert(ctx context.Context, message *models.ChatMessage) error
	FindByID(ctx context.Context, id string) (*models.ChatMessage, error)
	FindAll(ctx context.Context) ([]models.ChatMessage, error)
	Update(ctx context.Context, id string, message *models.ChatMessage) error
	Delete(ctx context.Context, id string) error
}

type ChatMessageRepository struct {
	db *pgxpool.Pool
}

func NewChatMessageRepository(db *pgxpool.Pool) *ChatMessageRepository {
	return &ChatMessageRepository{db: db}
}

func (r *ChatMessageRepository) Insert(ctx context.Context, message *models.ChatMessage) error {
	query := `INSERT INTO chat.chat_messages (id, session_id, sender_role, message_content, metadata, created_at) VALUES ($1, $2, $3, $4, $5, NOW())`
	message.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, message.ID, message.SessionID, message.SenderRole, message.MessageContent, message.Metadata)
	return err
}

func (r *ChatMessageRepository) FindByID(ctx context.Context, id string) (*models.ChatMessage, error) {
	query := `SELECT id, session_id, sender_role, message_content, COALESCE(metadata, '{}'::jsonb), created_at FROM chat.chat_messages WHERE id = $1`
	var message models.ChatMessage
	err := r.db.QueryRow(ctx, query, id).Scan(&message.ID, &message.SessionID, &message.SenderRole, &message.MessageContent, &message.Metadata, &message.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *ChatMessageRepository) FindAll(ctx context.Context) ([]models.ChatMessage, error) {
	query := `SELECT id, session_id, sender_role, message_content, COALESCE(metadata, '{}'::jsonb), created_at FROM chat.chat_messages ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := []models.ChatMessage{}
	for rows.Next() {
		var message models.ChatMessage
		if err := rows.Scan(&message.ID, &message.SessionID, &message.SenderRole, &message.MessageContent, &message.Metadata, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, rows.Err()
}

func (r *ChatMessageRepository) Update(ctx context.Context, id string, message *models.ChatMessage) error {
	query := `UPDATE chat.chat_messages SET session_id = $2, sender_role = $3, message_content = $4, metadata = $5 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, message.SessionID, message.SenderRole, message.MessageContent, message.Metadata)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ChatMessageRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM chat.chat_messages WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
