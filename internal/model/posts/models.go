package posts

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	PostModel struct {
		ID           int64    `db:"id"`
		UserID       int64    `db:"user_id"`
		PostTitle    string   `db:"post_title"`
		PostContent  string   `db:"post_content"`
		PostHashtags string   `db:"post_hashtags"`
		CreatedAt    string   `db:"created_at"`
		UpdatedAt    string   `db:"updated_at"`
		CreatedBy    string   `db:"created_by"`
		UpdatedBy    string   `db:"updated_by"`
	}
)

type (
	GetAllPostsResponse struct {
		Posts      []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID           int64    `json:"id"`
		UserID       int64    `json:"userId"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)
