package models

type KnowledgeBase struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	SourceType string   `json:"source_type"`
	Tags       []string `json:"tags,omitempty"`
	VectorID   string   `json:"vector_id,omitempty"`
	CreatedAt  string   `json:"created_at"`
}
