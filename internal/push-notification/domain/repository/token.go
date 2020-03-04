package repository

import (
	"github.com/abyssparanoia/rapid-go/internal/push-notification/domain/model"

	"context"
)

// Token ... token repository interface
type Token interface {
	GetByPlatformAndDeviceIDAndUserID(ctx context.Context,
		appID, userID, deviceID string,
		platform model.Platform) (*model.Token, error)
	List(ctx context.Context) ([]*model.Token, error)
	ListByUserID(ctx context.Context,
		appID, userID string) ([]*model.Token, error)
	Create(ctx context.Context,
		token *model.Token) (*model.Token, error)
	Update(ctx context.Context,
		token *model.Token) error
	Delete(ctx context.Context,
		tokenID string) error
}
