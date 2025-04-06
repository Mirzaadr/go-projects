package models

import "time"

type Todo struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}
