package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=256,alpha"`
	Username string `json:"username" validate:"required,max=25,min=4,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	Message string
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
