package services

import (
	"encoding/json"
	"net/http"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"

	"gorm.io/gorm/clause"
)

type LocationService struct {
	core.BaseService
	GetDto dto.GetLocationDto
}

func (ls *LocationService) GetAll() {
	var location []models.Location

	ls.DBLink.Preload(clause.Associations).Find(&location)

	jsonData, _ := json.Marshal(location)

	ls.Wri.Write(jsonData)
}

func (ls *LocationService) GetById() {
	var location models.Location

	ls.DBLink.Preload("Continent").Preload("Language").First(&location, ls.GetDto.Id)

	jsonData, _ := json.Marshal(location)

	ls.Wri.Write(jsonData)
}

func (ls *LocationService) Create(_dto dto.AddLocationDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		location := models.Location{
			LanguageId:  _dto.LanguageId,
			UserId:      _dto.UserId,
			CountryId:   _dto.CountryId,
			Name:        _dto.Name,
			City:        _dto.City,
			Description: _dto.Description,
			IsVisible:   &_dto.IsVisible,
			IsArchive:   &_dto.IsArchive,
			IsDeleted:   &_dto.IsDeleted,
		}

		ls.DBLink.Create(&location)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"location_id": location.Id,
		})

		ls.Wri.WriteHeader(http.StatusCreated)
		ls.Wri.Write(jsonData)
	} else {
		ls.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (ls *LocationService) Save(_dto dto.SaveLocationDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var location models.Location

		ls.DBLink.First(&location, _dto.Id)

		if len(_dto.Name) > 0 && _dto.Name != location.Name {
			location.Name = _dto.Name
		}

		if len(_dto.City) > 0 && _dto.Name != location.City {
			location.City = _dto.City
		}

		if len(_dto.Description) > 0 && _dto.Description != location.Description {
			location.Description = _dto.Description
		}

		if _dto.CountryId > 0 && _dto.CountryId != location.CountryId {
			location.CountryId = _dto.CountryId
		}

		if _dto.LanguageId > 0 && _dto.LanguageId != location.LanguageId {
			location.LanguageId = _dto.LanguageId
		}

		if &_dto.IsVisible != location.IsVisible {
			location.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != location.IsArchive {
			location.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != location.IsDeleted {
			location.IsDeleted = &_dto.IsDeleted
		}

		ls.DBLink.Save(&location)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		ls.Wri.Write(jsonData)
	} else {
		ls.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (ls *LocationService) Delete() {
	var location models.Location

	ls.DBLink.First(&location, ls.GetDto.Id)
	ls.DBLink.Delete(&location)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	ls.Wri.Write(jsonData)
}
