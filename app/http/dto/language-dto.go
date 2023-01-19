package dto

type GetLanguageDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddLanguageDto struct {
	Name        string `json:"name" validate:"required"`
	LangCode    string `json:"lang_code" validate:"required"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveLanguageDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	Name        string `json:"name" validate:"omitempty"`
	LangCode    string `json:"lang_code" validate:"omitempty"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
