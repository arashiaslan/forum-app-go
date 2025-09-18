package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/arashiaslan/forum-app-go/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUsers(ctx, req.Email, req.Username)
	if err!= nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
	}

	return s.membershipRepo.CreateUser(ctx, &model)
}