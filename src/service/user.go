package service

import (
	"context"

	"github.com/abyssparanoia/rapid-go/src/domain/model"
)

// User ... inteface of User service
type User interface {
	Get(ctx context.Context, userID int64) (*model.User, error)
}
