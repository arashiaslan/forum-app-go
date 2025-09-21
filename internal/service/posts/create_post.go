package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/arashiaslan/forum-app-go/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePosts(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHashtags, ",")
	
	now := time.Now()
	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHastags,
		CreatedAt:    now.String(),
		UpdatedAt:    now.String(),
		CreatedBy:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreatePosts(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
		return err
	}
	return nil
}
