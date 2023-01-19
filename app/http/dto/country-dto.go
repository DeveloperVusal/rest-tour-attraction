package dto

type GetCountryDto struct {
	Id uint `json:"id" validate:"required,numeric"`
}

type AddCountryDto struct {
	ContinentId uint   `json:"continent_id" validate:"required,numeric"`
	LanguageId  uint   `json:"language_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"required"`
	Capital     string `json:"capital" validate:"omitempty"`
	Code2       string `json:"code2" validate:"omitempty,alpha"`
	Code3       string `json:"code3" validate:"omitempty,alpha"`
	Iso         uint   `json:"iso" validate:"omitempty,numeric"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}

type SaveCountryDto struct {
	Id          uint   `json:"id" validate:"required,numeric"`
	ContinentId uint   `json:"continent_id" validate:"required,numeric"`
	LanguageId  uint   `json:"language_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"omitempty"`
	Capital     string `json:"capital" validate:"omitempty"`
	Code2       string `json:"code2" validate:"omitempty,alpha"`
	Code3       string `json:"code3" validate:"omitempty,alpha"`
	Iso         uint   `json:"iso" validate:"omitempty,numeric"`
	ModeratorId uint   `json:"moderator_id" validate:"omitempty,numeric"`
	IsVisible   bool   `json:"is_visible" validate:"omitempty,boolean"`
	IsArchive   bool   `json:"is_archive" validate:"omitempty,boolean"`
	IsDeleted   bool   `json:"is_deleted" validate:"omitempty,boolean"`
}
