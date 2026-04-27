package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatSessionRepositoryInterface interface {
	Insert(ctx context.Context, session *models.ChatSession) error
	FindByID(ctx context.Context, id string) (*models.ChatSession, error)
	FindAll(ctx context.Context) ([]models.ChatSession, error)
	Update(ctx context.Context, id string, session *models.ChatSession) error
	Delete(ctx context.Context, id string) error
}

type ChatSessionRepository struct {
	db *pgxpool.Pool
}

func NewChatSessionRepository(db *pgxpool.Pool) *ChatSessionRepository {
	return &ChatSessionRepository{db: db}
}

func (r *ChatSessionRepository) Insert(ctx context.Context, session *models.ChatSession) error {
	query := `INSERT INTO chat.chat_sessions (id, user_id, started_at, last_activity) VALUES ($1, $2, NOW(), NOW())`
	session.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, session.ID, session.UserID)
	return err
}

func (r *ChatSessionRepository) FindByID(ctx context.Context, id string) (*models.ChatSession, error) {
	query := `SELECT id, user_id, started_at, last_activity FROM chat.chat_sessions WHERE id = $1`
	var session models.ChatSession
	err := r.db.QueryRow(ctx, query, id).Scan(&session.ID, &session.UserID, &session.StartedAt, &session.LastActivity)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *ChatSessionRepository) FindAll(ctx context.Context) ([]models.ChatSession, error) {
	query := `SELECT id, user_id, started_at, last_activity FROM chat.chat_sessions ORDER BY last_activity DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sessions := []models.ChatSession{}
	for rows.Next() {
		var session models.ChatSession
		if err := rows.Scan(&session.ID, &session.UserID, &session.StartedAt, &session.LastActivity); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, rows.Err()
}

func (r *ChatSessionRepository) Update(ctx context.Context, id string, session *models.ChatSession) error {
	query := `UPDATE chat.chat_sessions SET user_id = $2, last_activity = NOW() WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, session.UserID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ChatSessionRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM chat.chat_sessions WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
