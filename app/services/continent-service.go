package services

import (
	"encoding/json"
	"net/http"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type ContinentService struct {
	core.BaseService
	GetDto dto.GetContinentDto
}

func (cs *ContinentService) GetAll() {
	var continent []models.Continent

	cs.DBLink.Preload("Language").Find(&continent)

	jsonData, _ := json.Marshal(continent)

	cs.Wri.Write(jsonData)
}

func (cs *ContinentService) GetById() {
	var continent models.Continent

	cs.DBLink.Preload("Language").First(&continent, cs.GetDto.Id)

	jsonData, _ := json.Marshal(continent)

	cs.Wri.Write(jsonData)
}

func (cs *ContinentService) Create(_dto dto.AddContinentDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		continent := models.Continent{
			Name:       _dto.Name,
			LanguageId: _dto.LanguageId,
			IsVisible:  &_dto.IsVisible,
			IsArchive:  &_dto.IsArchive,
			IsDeleted:  &_dto.IsDeleted,
		}

		if _dto.ModeratorId > 0 {
			continent.ModeratorId = _dto.ModeratorId
		}

		cs.DBLink.Create(&continent)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"continent_id": continent.Id,
		})

		cs.Wri.WriteHeader(http.StatusCreated)
		cs.Wri.Write(jsonData)
	} else {
		cs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (cs *ContinentService) Save(_dto dto.SaveContinentDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var continent models.Continent

		cs.DBLink.First(&continent, _dto.Id)

		if len(_dto.Name) > 0 && _dto.Name != continent.Name {
			continent.Name = _dto.Name
		}

		if _dto.LanguageId > 0 && _dto.LanguageId != continent.LanguageId {
			continent.LanguageId = _dto.LanguageId
		}

		if _dto.ModeratorId > 0 && _dto.ModeratorId != continent.ModeratorId {
			continent.ModeratorId = _dto.ModeratorId
		}

		if &_dto.IsVisible != continent.IsVisible {
			continent.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != continent.IsArchive {
			continent.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != continent.IsDeleted {
			continent.IsDeleted = &_dto.IsDeleted
		}

		cs.DBLink.Save(&continent)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		cs.Wri.Write(jsonData)
	} else {
		cs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (cs *ContinentService) Delete() {
	var continent models.Continent

	cs.DBLink.First(&continent, cs.GetDto.Id)
	cs.DBLink.Delete(&continent)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	cs.Wri.Write(jsonData)
}
