package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CountryService struct {
	core.BaseService
	GetDto dto.GetCountryDto
}

func (cs *CountryService) GetAll() {
	var country []models.Country
	var res *gorm.DB

	getLangId := cs.Req.FormValue("language_id")
	getFull, _ := strconv.ParseBool(cs.Req.FormValue("full"))

	if getFull {
		res = cs.DBLink.Preload(clause.Associations).Find(&country, "language_id = ?", getLangId)
	} else {
		res = cs.DBLink.Preload(clause.Associations).Find(&country, "language_id = ? AND is_visible = ? AND is_archive = ?", getLangId, true, false)
	}

	jsonData, _ := json.Marshal(map[string]interface{}{
		"count": res.RowsAffected,
		"list":  country,
	})

	cs.Wri.Write(jsonData)
}

func (cs *CountryService) GetById() {
	var country models.Country

	cs.DBLink.Preload("Continent").Preload("Language").First(&country, cs.GetDto.Id)

	jsonData, _ := json.Marshal(country)

	cs.Wri.Write(jsonData)
}

func (cs *CountryService) Create(_dto dto.AddCountryDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		country := models.Country{
			ContinentId: _dto.ContinentId,
			LanguageId:  _dto.LanguageId,
			Name:        _dto.Name,
			Capital:     _dto.Capital,
			Code2:       _dto.Code2,
			Code3:       _dto.Code3,
			Iso:         _dto.Iso,
			IsVisible:   &_dto.IsVisible,
			IsArchive:   &_dto.IsArchive,
			IsDeleted:   &_dto.IsDeleted,
		}

		if _dto.ModeratorId > 0 {
			country.ModeratorId = _dto.ModeratorId
		}

		cs.DBLink.Create(&country)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"country_id": country.Id,
		})

		cs.Wri.WriteHeader(http.StatusCreated)
		cs.Wri.Write(jsonData)
	} else {
		cs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (cs *CountryService) Save(_dto dto.SaveCountryDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var country models.Country

		cs.DBLink.First(&country, _dto.Id)

		if len(_dto.Name) > 0 && _dto.Name != country.Name {
			country.Name = _dto.Name
		}

		if len(_dto.Capital) > 0 && _dto.Capital != country.Capital {
			country.Capital = _dto.Capital
		}

		if len(_dto.Code2) > 0 && _dto.Code2 != country.Code2 {
			country.Code2 = _dto.Code2
		}

		if _dto.ContinentId > 0 && _dto.ContinentId != country.ContinentId {
			country.ContinentId = _dto.ContinentId
		}

		if _dto.LanguageId > 0 && _dto.LanguageId != country.LanguageId {
			country.LanguageId = _dto.LanguageId
		}

		if _dto.Iso > 0 && _dto.Iso != country.Iso {
			country.Iso = _dto.Iso
		}

		if _dto.ModeratorId > 0 && _dto.ModeratorId != country.ModeratorId {
			country.ModeratorId = _dto.ModeratorId
		}

		if &_dto.IsVisible != country.IsVisible {
			country.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != country.IsArchive {
			country.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != country.IsDeleted {
			country.IsDeleted = &_dto.IsDeleted
		}

		cs.DBLink.Save(&country)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		cs.Wri.Write(jsonData)
	} else {
		cs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (cs *CountryService) Delete() {
	var country models.Country

	cs.DBLink.First(&country, cs.GetDto.Id)
	cs.DBLink.Delete(&country)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	cs.Wri.Write(jsonData)
}
