package entities

import (
	"bytes"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
	LastName  string    `json:"last_name,omitempty" validate:"omitempty,lte=30,gte=2,alpha"`
	Username  string    `json:"username" validate:"lte=30,gte=2,required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"gte=5,required"`
}

func (u *User) ComparePassword(pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)); err != nil {
		return err
	}

	return nil
}

func (u *User) HashPassword() (string, error) {
	arr, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return bytes.NewBuffer(arr).String(), err

}
