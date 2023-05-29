package models

import "time"

type Base struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
}
