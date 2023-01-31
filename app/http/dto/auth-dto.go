package dto

type AuthDto struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"passwd" validate:"required"`
}
