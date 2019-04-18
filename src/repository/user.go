package repository

import (
	"context"

	"github.com/abyssparanoia/rapid-go/src/model"
)

// User ... ユーザーレポジトリのinterface
type User interface {
	Get(ctx context.Context, userID int64) (*model.User, error)
}