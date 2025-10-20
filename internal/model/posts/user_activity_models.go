package posts

type (
	UserActivityRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type (
	UserActivityModel struct {
		ID        int64  `db:"id"`
		PostID    int64  `db:"post_id"`
		UserID    int64  `db:"user_id"`
		IsLiked   bool   `db:"is_liked"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)