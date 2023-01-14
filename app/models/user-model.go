package models

import (
	"attrtour/core"
	"encoding/json"
	"strconv"

	"attrtour/app/http/dto"
)

type UserModel struct {
	core.BaseModel
	GetDto dto.GetUserDto
}

func (um *UserModel) Create(_dto dto.AddUserDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		json, _ := json.Marshal(_dto)
		um.Wri.Write(json)
	} else {
		um.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (um *UserModel) Save(_dto dto.SaveUserDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		json, _ := json.Marshal(_dto)
		um.Wri.Write(json)
	} else {
		um.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (um *UserModel) GetId() {
	um.Wri.Write([]byte("Get user id by model => " + strconv.Itoa(um.GetDto.Id) + "\n"))
}
