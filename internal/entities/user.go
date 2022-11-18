package entities

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
	LastName  string    `json:"last_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
	Username  string    `json:"username" validate:"lte=30,gte=2,required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"gte=5,required"`
}
