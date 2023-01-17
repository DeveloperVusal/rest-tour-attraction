package dto

type GetGroupDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddGroupDto struct {
	Name        string `json:"name" validate:"required,alphanumunicode"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveGroupDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	Name        string `json:"name" validate:"omitempty,alphanumunicode"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
