package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,gte=2,lte=60"`
	Description string    `json:"description,omitempty" validate:"omitempty"`
	CreatedOn   time.Time `json:"created_on"`
	Priority    int8      `json:"priority"`
	Status      int8      `json:"status"`
}
