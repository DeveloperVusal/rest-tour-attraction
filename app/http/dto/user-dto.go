package dto

type GetUserDto struct {
	Id int `json:"id" validate:"required,numeric"`
}

type AddUserDto struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,contains"`
	Email    string `json:"email" validate:"email"`
}

type SaveUserDto struct {
	Id       int    `json:"id" validate:"required,numeric"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,contains"`
	Email    string `json:"email" validate:"email"`
}
