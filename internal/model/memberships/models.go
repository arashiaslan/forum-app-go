package memberships

type (
	SignUpRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

type (
	LoginResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

type (
	UserModel struct {
		ID        int64  `db:"id"`
		Email     string `db:"email"`
		Username  string `db:"username"`
		Password  string `db:"password"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)
