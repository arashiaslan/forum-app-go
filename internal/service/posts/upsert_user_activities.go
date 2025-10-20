package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/arashiaslan/forum-app-go/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, PostID int64, UserID int64, request posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostID: PostID,
		UserID: UserID, 
		IsLiked: request.IsLiked, 
		CreatedAt: now.Format(time.RFC3339), 
		UpdatedAt: now.Format(time.RFC3339), 
		CreatedBy: strconv.FormatInt(UserID, 10), 
		UpdatedBy: strconv.FormatInt(UserID, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user activity")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return nil
		}

		err = s.postRepo.CreateUserActivity(ctx, model)
		if err != nil {
			log.Error().Err(err).Msg("failed to create user activity")
			return err
		}
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("failed to update user activity")
		return err
	}
	return nil
}
