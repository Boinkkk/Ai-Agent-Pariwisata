package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type KnowledgeBaseRepositoryInterface interface {
	Insert(ctx context.Context, knowledge *models.KnowledgeBase) error
	FindByID(ctx context.Context, id string) (*models.KnowledgeBase, error)
	FindAll(ctx context.Context) ([]models.KnowledgeBase, error)
	Update(ctx context.Context, id string, knowledge *models.KnowledgeBase) error
	Delete(ctx context.Context, id string) error
}

type KnowledgeBaseRepository struct {
	db *pgxpool.Pool
}

func NewKnowledgeBaseRepository(db *pgxpool.Pool) *KnowledgeBaseRepository {
	return &KnowledgeBaseRepository{db: db}
}

func (r *KnowledgeBaseRepository) Insert(ctx context.Context, knowledge *models.KnowledgeBase) error {
	query := `INSERT INTO rag.knowledge_base (id, title, content, source_type, tags, vector_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	knowledge.ID = uuid.NewString()
	_, err := r.db.Exec(ctx, query, knowledge.ID, knowledge.Title, knowledge.Content, knowledge.SourceType, knowledge.Tags, nullIfEmpty(knowledge.VectorID))
	return err
}

func (r *KnowledgeBaseRepository) FindByID(ctx context.Context, id string) (*models.KnowledgeBase, error) {
	query := `SELECT id::text, title, content, source_type, COALESCE(tags, '{}'), COALESCE(vector_id::text, ''), created_at::text FROM rag.knowledge_base WHERE id = $1`
	var knowledge models.KnowledgeBase
	err := r.db.QueryRow(ctx, query, id).Scan(&knowledge.ID, &knowledge.Title, &knowledge.Content, &knowledge.SourceType, &knowledge.Tags, &knowledge.VectorID, &knowledge.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &knowledge, nil
}

func (r *KnowledgeBaseRepository) FindAll(ctx context.Context) ([]models.KnowledgeBase, error) {
	query := `SELECT id::text, title, content, source_type, COALESCE(tags, '{}'), COALESCE(vector_id::text, ''), created_at::text FROM rag.knowledge_base ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.KnowledgeBase{}
	for rows.Next() {
		var knowledge models.KnowledgeBase
		if err := rows.Scan(&knowledge.ID, &knowledge.Title, &knowledge.Content, &knowledge.SourceType, &knowledge.Tags, &knowledge.VectorID, &knowledge.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, knowledge)
	}
	return items, rows.Err()
}

func (r *KnowledgeBaseRepository) Update(ctx context.Context, id string, knowledge *models.KnowledgeBase) error {
	query := `UPDATE rag.knowledge_base SET title = $2, content = $3, source_type = $4, tags = $5, vector_id = $6 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, knowledge.Title, knowledge.Content, knowledge.SourceType, knowledge.Tags, nullIfEmpty(knowledge.VectorID))
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *KnowledgeBaseRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM rag.knowledge_base WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
