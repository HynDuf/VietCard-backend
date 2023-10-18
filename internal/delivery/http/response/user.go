package response

import "vietcard-backend/internal/domain/entity"

type ErrorResponse struct {
	Message string `json:"error"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateUserResponse struct {
	Success bool        `json:"success"`
	User    entity.User `json:"user"`
}
