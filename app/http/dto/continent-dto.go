package dto

type GetContinentDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddContinentDto struct {
	LanguageId  uint   `json:"language_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"required"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveContinentDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	LanguageId  uint   `json:"language_id" validate:"omitempty,numeric"`
	Name        string `json:"name" validate:"omitempty"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
