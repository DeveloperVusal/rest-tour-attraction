package services

import (
	"encoding/json"
	"net/http"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type GroupService struct {
	core.BaseService
	GetDto dto.GetGroupDto
}

func (gs *GroupService) GetAll() {
	var group []models.Group

	gs.DBLink.Find(&group)

	jsonData, _ := json.Marshal(group)

	gs.Wri.Write(jsonData)
}

func (gs *GroupService) GetById() {
	var group models.Group

	gs.DBLink.First(&group, gs.GetDto.Id)

	jsonData, _ := json.Marshal(group)

	gs.Wri.Write(jsonData)
}

func (gs *GroupService) Create(_dto dto.AddGroupDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		group := models.Group{
			Name:      _dto.Name,
			IsVisible: &_dto.IsVisible,
			IsArchive: &_dto.IsArchive,
			IsDeleted: &_dto.IsDeleted,
		}

		if _dto.ModeratorId > 0 {
			group.ModeratorId = _dto.ModeratorId
		}

		gs.DBLink.Create(&group)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"group_id": group.Id,
		})

		gs.Wri.WriteHeader(http.StatusCreated)
		gs.Wri.Write(jsonData)
	} else {
		gs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (gs *GroupService) Save(_dto dto.SaveGroupDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var group models.Group

		gs.DBLink.First(&group, _dto.Id)

		if len(_dto.Name) > 0 && _dto.Name != group.Name {
			group.Name = _dto.Name
		}

		if _dto.ModeratorId > 0 && _dto.ModeratorId != group.ModeratorId {
			group.ModeratorId = _dto.ModeratorId
		}

		if &_dto.IsVisible != group.IsVisible {
			group.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != group.IsArchive {
			group.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != group.IsDeleted {
			group.IsDeleted = &_dto.IsDeleted
		}

		gs.DBLink.Save(&group)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		gs.Wri.Write(jsonData)
	} else {
		gs.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (gs *GroupService) Delete() {
	var group models.Group

	gs.DBLink.First(&group, gs.GetDto.Id)
	gs.DBLink.Delete(&group)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	gs.Wri.Write(jsonData)
}
