package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type LanguageService struct {
	core.BaseService
	GetDto dto.GetLanguageDto
}

func (ls *LanguageService) GetAll() {
	var language []models.Language

	getFull, _ := strconv.ParseBool(ls.Req.FormValue("full"))

	if getFull {
		ls.DBLink.Find(&language)
	} else {
		ls.DBLink.Find(&language, "is_visible = ? AND is_archive = ?", true, false)
	}

	jsonData, _ := json.Marshal(language)

	ls.Wri.Write(jsonData)
}

func (ls *LanguageService) GetById() {
	var language models.Language

	ls.DBLink.First(&language, ls.GetDto.Id)

	jsonData, _ := json.Marshal(language)

	ls.Wri.Write(jsonData)
}

func (ls *LanguageService) Create(_dto dto.AddLanguageDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		language := models.Language{
			Name:      _dto.Name,
			LangCode:  _dto.LangCode,
			IsVisible: &_dto.IsVisible,
			IsArchive: &_dto.IsArchive,
			IsDeleted: &_dto.IsDeleted,
		}

		if _dto.ModeratorId > 0 {
			language.ModeratorId = _dto.ModeratorId
		}

		ls.DBLink.Create(&language)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"language_id": language.Id,
		})

		ls.Wri.WriteHeader(http.StatusCreated)
		ls.Wri.Write(jsonData)
	} else {
		ls.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (ls *LanguageService) Save(_dto dto.SaveLanguageDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var language models.Language

		ls.DBLink.First(&language, _dto.Id)

		if len(_dto.Name) > 0 && _dto.Name != language.Name {
			language.Name = _dto.Name
		}

		if len(_dto.LangCode) > 0 && _dto.Name != language.LangCode {
			language.LangCode = _dto.LangCode
		}

		if _dto.ModeratorId > 0 && _dto.ModeratorId != language.ModeratorId {
			language.ModeratorId = _dto.ModeratorId
		}

		if &_dto.IsVisible != language.IsVisible {
			language.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != language.IsArchive {
			language.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != language.IsDeleted {
			language.IsDeleted = &_dto.IsDeleted
		}

		ls.DBLink.Save(&language)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		ls.Wri.Write(jsonData)
	} else {
		ls.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (ls *LanguageService) Delete() {
	var language models.Language

	ls.DBLink.First(&language, ls.GetDto.Id)
	ls.DBLink.Delete(&language)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	ls.Wri.Write(jsonData)
}
