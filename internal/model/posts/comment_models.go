package posts

type (
	CreateCommentRequest struct {
		CommentContent string `json:"commentContent"`
	}
)

type (
	CommentModel struct {
		ID             int64  `db:"id"`
		PostID         int64  `db:"post_id"`
		UserID         int64  `db:"user_id"`
		CommentContent string `db:"comment_content"`
		CreatedAt      string `db:"created_at"`
		UpdatedAt      string `db:"updated_at"`
		CreatedBy      string `db:"created_by"`
		UpdatedBy      string `db:"updated_by"`
	}
)
