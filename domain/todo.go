package domain

import "time"

// Todo struct
type Todo struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty`
}
