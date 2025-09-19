package memberships

import (
	"context"

	"github.com/arashiaslan/forum-app-go/internal/configs"
	"github.com/arashiaslan/forum-app-go/internal/model/memberships"
)

type membershipRepository interface {
	GetUsers(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user *memberships.UserModel) error
}

type service struct {
	cfg *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg: cfg,
		membershipRepo: membershipRepo,
	}
}