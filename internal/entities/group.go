package entities

import "github.com/google/uuid"

type Group struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"lte=60,required"`
	Description string    `json:"description,omitempty" validate:"omitempty"`
	CreatorId   string    `json:"creator_id"`
}
