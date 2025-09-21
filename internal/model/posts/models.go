package posts

type (
	CreatePostRequest struct {
		PostTitle    string `json:"postTitle"`
		PostContent  string `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	PostModel struct {
		ID           int64  `db:"id"`
		UserID       int64  `db:"user_id"`
		PostTitle    string `db:"post_title"`
		PostContent  string `db:"post_content"`
		PostHashtags string `db:"post_hashtags"`
		CreatedAt    string `db:"created_at"`
		UpdatedAt    string `db:"updated_at"`
		CreatedBy    string `db:"created_by"`
		UpdatedBy    string `db:"updated_by"`
	}
)