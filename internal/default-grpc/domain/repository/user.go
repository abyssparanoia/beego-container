package repository

import (
	"context"

	"github.com/abyssparanoia/rapid-go/internal/default-grpc/domain/model"
)

// User ... user interface
type User interface {
	Get(ctx context.Context, userID string) (*model.User, error)
}
