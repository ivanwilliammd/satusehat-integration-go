package terminology

import "time"

type LoincAnswer struct {
	ID               int64      `json:"id"`
	LoincNumber      *string    `json:"LoincNumber,omitempty"`
	AnswerListID     *string    `json:"AnswerListId,omitempty"`
	AnswerListName   *string    `json:"AnswerListName,omitempty"`
	AnswerStringID   *string    `json:"AnswerStringId,omitempty"`
	SequenceNumber   *int       `json:"SequenceNumber,omitempty"`
	DisplayText      *string    `json:"DisplayText,omitempty"`
	ExtCodeID       *string    `json:"ExtCodeId,omitempty"`
	ExtCodeDisplayName *string  `json:"ExtCodeDisplayName,omitempty"`
	ExtCodeSystem   *string    `json:"ExtCodeSystem,omitempty"`
	CreatedAt       *time.Time  `json:"created_at,omitempty"`
	UpdatedAt       *time.Time  `json:"updated_at,omitempty"`
	DeletedAt       *time.Time  `json:"deleted_at,omitempty"`
}
