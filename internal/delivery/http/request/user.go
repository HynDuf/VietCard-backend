package request

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type UpdateUserRequest struct {
	Name             *string `json:"name" bson:"name,omitempty"`
	HashedPassword   *string `json:"hashed_password" bson:"hashed_password,omitempty"`
	MaxNewCardsLearn *int    `json:"max_new_cards_learn" bson:"max_new_cards_learn,omitempty"`
	MaxCardsReview   *int    `json:"max_cards_review" bson:"max_cards_review,omitempty"`
}

type AddXPRequest struct {
	XP int `json:"xp" binding:"required"`
}
