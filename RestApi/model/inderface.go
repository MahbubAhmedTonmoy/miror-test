package model

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	Get(ctx context.Context, id uuid.UUID) (*User, error)
}

type UserRepository interface {
	GindById(ctx context.Context, id uuid.UUID) (*User, error)
}
