package dto

type GetUserDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddUserDto struct {
	GroupId     uint   `json:"group_id" validate:"required,numeric"`
	Username    string `json:"username" validate:"required,alphanum"`
	Password    string `json:"password" validate:"required,contains"`
	Email       string `json:"email" validate:"omitempty,email"`
	Name        string `json:"name" validate:"omitempty,alphaunicode"`
	Surname     string `json:"surname" validate:"omitempty,alphaunicode"`
	Age         uint8  `json:"age" validate:"omitempty,numeric"`
	Gender      string `json:"gender" validate:"omitempty,alpha"`
	IsConfirm   bool   `json:"is_confirm" validate:"omitempty,boolean"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsBlocked   bool   `json:"is_blocked" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveUserDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	GroupId     uint   `json:"group_id" validate:"required,numeric"`
	Username    string `json:"username" validate:"required,alphanum"`
	Password    string `json:"password" validate:"required,contains"`
	Email       string `json:"email" validate:"omitempty,email"`
	Name        string `json:"name" validate:"omitempty,alphaunicode"`
	Surname     string `json:"surname" validate:"omitempty,alphaunicode"`
	Age         uint8  `json:"age" validate:"omitempty,numeric"`
	Gender      string `json:"gender" validate:"omitempty,alpha"`
	IsConfirm   bool   `json:"is_confirm" validate:"omitempty,boolean"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsBlocked   bool   `json:"is_blocked" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
