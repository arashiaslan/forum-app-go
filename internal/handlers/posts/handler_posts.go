package posts

import (
	"context"

	"github.com/arashiaslan/forum-app-go/internal/middleware"
	"github.com/arashiaslan/forum-app-go/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePosts(ctx context.Context, userID int64, req posts.CreatePostRequest) error
}
type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
}