package repository

import (
	"github.com/abyssparanoia/rapid-go/push-notification/domain/model"

	"context"
)

// Token ... token repository interface
type Token interface {
	GetByPlatformAndDeviceID(ctx context.Context,
		appID, userID, deviceID string,
		platform model.Platform) (*model.Token, error)
}
