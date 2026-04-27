package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	ID             uuid.UUID       `json:"id"`
	SessionID      uuid.UUID       `json:"session_id"`
	SenderRole     string          `json:"sender_role"`
	MessageContent string          `json:"message_content"`
	Metadata       json.RawMessage `json:"metadata"`
	CreatedAt      time.Time       `json:"created_at"`
}
