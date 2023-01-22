package dto

import "mime/multipart"

type GetMainImageDto struct {
	Id uint `schema:"id" validate:"required,numeric"`
}

type UploadMainImageDto struct {
	LocationId       uint                  `schema:"location_id" validate:"omitempty,numeric"`
	UserId           uint                  `schema:"user_id" validate:"required,numeric"`
	LanguageId       uint                  `schema:"language_id" validate:"required,numeric"`
	IsMain           bool                  `schema:"is_main" validate:"omitempty,boolean"`
	File             *multipart.FileHeader `schema:"file" validate:"required"`
	ShortDescription string                `schema:"short_description" validate:"omitempty"`
	IsVisible        bool                  `schema:"is_visible" validate:"omitempty,boolean"`
	IsArchive        bool                  `schema:"is_archive" validate:"omitempty,boolean"`
	IsDeleted        bool                  `schema:"is_deleted" validate:"omitempty,boolean"`
}

type SaveMainImageDto struct {
	Id               uint                  `schema:"id" validate:"required,numeric"`
	LocationId       uint                  `schema:"location_id" validate:"required,numeric"`
	UserId           uint                  `schema:"user_id" validate:"required,numeric"`
	LanguageId       uint                  `schema:"language_id" validate:"required,numeric"`
	IsMain           bool                  `schema:"is_main" validate:"omitempty,boolean"`
	File             *multipart.FileHeader `schema:"file" validate:"omitempty"`
	ShortDescription string                `schema:"short_description" validate:"omitempty"`
	IsVisible        bool                  `schema:"is_visible" validate:"omitempty,boolean"`
	IsArchive        bool                  `schema:"is_archive" validate:"omitempty,boolean"`
	IsDeleted        bool                  `schema:"is_deleted" validate:"omitempty,boolean"`
}
