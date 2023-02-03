package services

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type UserService struct {
	core.BaseService
	GetDto dto.GetUserDto
}

func (us *UserService) Init() {
	groupId := us.createRootGroup()
	us.createRootUser(groupId)
}

func (us *UserService) GetAll() {
	var user []models.User

	us.DBLink.Preload("Group").Find(&user)

	jsonData, _ := json.Marshal(user)

	us.Wri.Write(jsonData)
}

func (us *UserService) GetById() {
	var user models.User

	us.DBLink.Preload("Group").First(&user, us.GetDto.Id)

	jsonData, _ := json.Marshal(user)

	us.Wri.Write(jsonData)
}

func (us *UserService) Create(_dto dto.AddUserDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		hashPasswd, _ := us.HashPassword(_dto.Password)

		user := models.User{
			GroupId:   _dto.GroupId,
			Username:  _dto.Username,
			Password:  hashPasswd,
			IsConfirm: &_dto.IsConfirm,
			IsArchive: &_dto.IsArchive,
			IsBlocked: &_dto.IsBlocked,
			IsDeleted: &_dto.IsDeleted,
		}

		if len(_dto.Email) > 0 {
			user.Email = _dto.Email
		}

		if _dto.Age > 0 {
			user.Age = _dto.Age
		}

		if len(_dto.Gender) > 0 {
			user.Gender = _dto.Gender
		}

		if len(_dto.Name) > 0 {
			user.Name = _dto.Name
		}

		if len(_dto.Surname) > 0 {
			user.Surname = _dto.Surname
		}

		if _dto.ModeratorId > 0 {
			user.ModeratorId = _dto.ModeratorId
		}

		us.DBLink.Create(&user)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"user_id": user.Id,
		})

		us.Wri.WriteHeader(http.StatusCreated)
		us.Wri.Write(jsonData)
	} else {
		us.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (us *UserService) Save(_dto dto.SaveUserDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var user models.User

		us.DBLink.First(&user, _dto.Id)

		if _dto.GroupId > 0 && _dto.GroupId != user.GroupId {
			user.GroupId = _dto.GroupId
		}

		if len(_dto.Username) > 0 && _dto.Username != user.Username {
			user.Username = _dto.Username
		}

		if len(_dto.Password) > 0 {
			if !us.CheckPasswordHash(_dto.Password, user.Password) {
				hashPasswd, _ := us.HashPassword(_dto.Password)
				user.Password = hashPasswd
			}
		}

		if len(_dto.Email) > 0 && _dto.Email != user.Email {
			user.Email = _dto.Email
		}

		if _dto.Age > 0 && _dto.Age != user.Age {
			user.Age = _dto.Age
		}

		if len(_dto.Gender) > 0 && _dto.Gender != user.Gender {
			user.Gender = _dto.Gender
		}

		if len(_dto.Name) > 0 && _dto.Name != user.Name {
			user.Name = _dto.Name
		}

		if len(_dto.Surname) > 0 && _dto.Surname != user.Surname {
			user.Surname = _dto.Surname
		}

		if _dto.ModeratorId > 0 && _dto.ModeratorId != user.ModeratorId {
			user.ModeratorId = _dto.ModeratorId
		}

		if &_dto.IsConfirm != user.IsConfirm {
			user.IsConfirm = &_dto.IsConfirm
		}

		if &_dto.IsArchive != user.IsArchive {
			user.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsBlocked != user.IsBlocked {
			user.IsBlocked = &_dto.IsBlocked
		}

		if &_dto.IsDeleted != user.IsDeleted {
			user.IsDeleted = &_dto.IsDeleted
		}

		us.DBLink.Save(&user)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
			"result": map[string]interface{}{
				"user_id": user.Id,
			},
		})

		us.Wri.Write(jsonData)
	} else {
		us.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (us *UserService) Delete() {
	var user models.User

	us.DBLink.First(&user, us.GetDto.Id)
	us.DBLink.Delete(&user)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	us.Wri.Write(jsonData)
}

func (us *UserService) HashPassword(passwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(bytes), err
}

func (us *UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us *UserService) createRootGroup() uint {
	env := core.Helpers{}
	var group models.Group

	us.DBLink.First(&group, "name = ?", env.Env("APP_ROOT_GROUP"))

	if group.Id <= 0 {
		isVisible := true
		group = models.Group{
			Name:      env.Env("APP_ROOT_GROUP"),
			IsVisible: &isVisible,
		}

		us.DBLink.Save(&group)
	}

	return group.Id
}

func (us *UserService) createRootUser(groupId uint) {
	env := core.Helpers{}

	isConfirm := true
	user := models.User{
		GroupId:   groupId,
		Username:  env.Env("APP_ROOT_LOGIN"),
		Password:  env.Env("APP_ROOT_PASSWD"),
		IsConfirm: &isConfirm,
	}
	us.DBLink.Save(&user)
}
