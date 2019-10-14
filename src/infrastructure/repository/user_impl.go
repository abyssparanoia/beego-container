package repository

import (
	"context"

	"github.com/abyssparanoia/rapid-go/src/domain/model"
	"github.com/abyssparanoia/rapid-go/src/domain/repository"
	"github.com/abyssparanoia/rapid-go/src/infrastructure/entity"
	"github.com/abyssparanoia/rapid-go/src/lib/log"
	"github.com/abyssparanoia/rapid-go/src/lib/mysql"
)

type user struct {
	cli *mysql.Client
}

func (r *user) Get(ctx context.Context, userID string) (*model.User, error) {

	dsts := []*entity.User{}

	db := r.cli.GetDB(ctx).
		Where("id = ?", userID).
		Limit(1).
		Find(&dsts)

	if err := mysql.HandleErrors(db); err != nil {
		log.Errorm(ctx, "db.Find", err)
		return nil, err
	}

	return model.NewUsers(dsts)[0], nil
}

// NewUser ... get user repository
func NewUser(cli *mysql.Client) repository.User {
	return &user{cli}
}
