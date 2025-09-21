package posts

import (
	"context"

	"github.com/arashiaslan/forum-app-go/internal/configs"
	"github.com/arashiaslan/forum-app-go/internal/model/posts"
)

type postRepository interface {
	CreatePosts(ctx context.Context, model posts.PostModel) error 
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
