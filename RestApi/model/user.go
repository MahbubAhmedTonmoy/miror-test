package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"image_url" json:"imageUrl"`
	Website  string    `db:"website" json:"website"`
}
