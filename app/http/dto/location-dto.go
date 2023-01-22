package dto

type GetLocationDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddLocationDto struct {
	UserId      uint   `json:"user_id" validate:"required,numeric"`
	LanguageId  uint   `json:"language_id" validate:"required,numeric"`
	CountryId   uint   `json:"country_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"required"`
	City        string `json:"city" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveLocationDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	UserId      uint   `json:"user_id" validate:"required,numeric"`
	LanguageId  uint   `json:"language_id" validate:"required,numeric"`
	CountryId   uint   `json:"country_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"omitempty"`
	City        string `json:"city" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
