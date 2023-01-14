package models

import (
	"attrtour/core"
	"strconv"

	"attrtour/app/http/dto"
)

type UserModel struct {
	core.BaseModel
	Dto dto.GetUserDto
}

func (um *UserModel) GetId() {
	um.Wri.Write([]byte("Get user id by model => " + strconv.Itoa(um.Dto.Id) + "\n"))
}
